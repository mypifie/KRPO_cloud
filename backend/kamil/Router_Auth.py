from datetime import datetime, timedelta
from typing import Optional
from fastapi import APIRouter, Depends, HTTPException, status, Request
from fastapi.security import OAuth2PasswordBearer
from jose import JWTError, jwt
from passlib.context import CryptContext
from sqlalchemy.ext.asyncio import AsyncSession


from database_main import get_async_session
from models import *
from config import settings
from database_ORM import *

router = APIRouter(prefix="", tags=["auth"])

oauth2_scheme = OAuth2PasswordBearer(tokenUrl="auth/login")
pwd_context = CryptContext(schemes=["bcrypt"], deprecated="auto")


def verify_password(plain_password: str, hashed_password: str) -> bool:
    return pwd_context.verify(plain_password, hashed_password)


def get_password_hash(password: str) -> str:
    return pwd_context.hash(password)


def create_access_token(
    data: dict, expires_minutes: int = settings.ACCESS_TOKEN_EXPIRE_MINUTES
) -> str:
    to_encode = data.copy()
    expire = datetime.utcnow() + timedelta(minutes=expires_minutes)
    to_encode.update({"exp": expire})
    return jwt.encode(to_encode, settings.SECRET_KEY, algorithm=settings.ALGORITHM)


async def get_current_user_db(
    token: str = Depends(oauth2_scheme),
    db: AsyncSession = Depends(get_async_session),
):
    credentials_exception = HTTPException(
        status_code=status.HTTP_401_UNAUTHORIZED,
        detail="Неверные учетные данные",
        headers={"WWW-Authenticate": "Bearer"},
    )

    try:
        payload = jwt.decode(
            token, settings.SECRET_KEY, algorithms=[settings.ALGORITHM]
        )
        print("PAYLOAD", payload)
        email: str = payload.get("sub")
        if email is None:
            raise credentials_exception
    except JWTError:
        raise credentials_exception

    user = await get_user_by_email(db, email)
    if user is None:
        raise credentials_exception

    return user


# Роутеры
@router.post("/register")
async def register(request: Request, db: AsyncSession = Depends(get_async_session)):
    try:

        # Получаем JSON из тела запроса
        json_data = await request.json()

        # Валидация обязательных полей
        required_fields = ["email", "password", "username"]
        for field in required_fields:
            if field not in json_data:
                error_msg = f"Отсутствует обязательное поле: {field}"
                raise HTTPException(
                    status_code=status.HTTP_400_BAD_REQUEST,
                    detail=error_msg,
                )

        # Проверка уникальности email
        existing_user_email = await get_user_by_email(db, json_data["email"])

        if existing_user_email is not None:
            error_msg = "Email уже зарегистрирован"
            raise HTTPException(status_code=400, detail=error_msg)

        # Добавление пользователя
        hashed_password = get_password_hash(json_data["password"])

        auth_user = AuthUsers(email=json_data["email"], password_hash=hashed_password)
        db.add(auth_user)
        await db.commit()
        await db.refresh(auth_user)

        # Добавление профиля
        auth_user = await get_user_by_email(db, json_data["email"])

        user_profile = UserProfiles(id=auth_user.id, username=json_data["username"])
        db.add(user_profile)
        await db.commit()
        await db.refresh(user_profile)

        # Создание токена
        user_profile = await get_user_profile_by_id(db, auth_user.id)
        access_token = create_access_token(data={"sub": auth_user.email})

        return {
            "access_token": access_token,
            "token_type": "bearer",
            "user": {
                "email": auth_user.email,
                "username": user_profile.username,
                "isAdmin": user_profile.is_admin,
                "isBlocked": user_profile.is_blocked,
            },
        }

    except HTTPException as he:
        print("ОШИБКА РЕГИСТР")
        await db.rollback()
        raise he
    except Exception as e:
        import traceback

        print("ОШИБКА РЕГИСТР")
        traceback.print_exc()
        await db.rollback()
        raise HTTPException(
            status_code=status.HTTP_400_BAD_REQUEST,
            detail=f"Ошибка регистрации: {str(e)}",
        )


@router.post("/login")
async def login(request: Request, db: AsyncSession = Depends(get_async_session)):
    try:
        # Получаем JSON из тела запроса
        json_data = await request.json()

        # Валидация обязательных полей
        required_fields = ["email", "password"]
        for field in required_fields:
            if field not in json_data:
                raise HTTPException(
                    status_code=status.HTTP_400_BAD_REQUEST,
                    detail=f"Отсутствует обязательное поле: {field}",
                )

        email = json_data["email"]
        password = json_data["password"]

        # Поиск пользователя
        auth_user = await get_user_by_email(db, email)
        if not auth_user:
            raise HTTPException(status_code=401, detail="Неверный email/логин")

        # Проверка пароля
        if not verify_password(password, auth_user.password_hash):
            raise HTTPException(status_code=401, detail="Неверный пароль")

        # Создание токена
        access_token = create_access_token(data={"sub": auth_user.email})

        user_profile = await get_user_profile_by_id(db, auth_user.id)
        return {
            "access_token": access_token,
            "token_type": "bearer",
            "user": {
                "email": auth_user.email,
                "username": user_profile.username,
                "isAdmin": user_profile.is_admin,
                "isBlocked": user_profile.is_blocked,
            },
        }

    except HTTPException as he:
        raise he
    except Exception as e:
        raise HTTPException(
            status_code=status.HTTP_400_BAD_REQUEST,
            detail=f"Ошибка входа: {str(e)}",
        )
