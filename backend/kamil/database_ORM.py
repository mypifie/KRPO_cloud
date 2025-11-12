from typing import Optional
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy.exc import SQLAlchemyError
from sqlalchemy.future import select
from sqlalchemy import update, insert, delete
from models import *
import inspect


def get_caller_name():
    return inspect.currentframe().f_back.f_back.f_code.co_name


# Асинхронные функции для работы с БД
async def get_user_by_email(db: AsyncSession, email: str):
    try:
        result = await db.execute(select(AuthUsers).where(AuthUsers.email == email))
        return result.scalar_one_or_none()

    except SQLAlchemyError as e:
        print(f"ОШИБКА В {get_caller_name()}")
        await db.rollback()
        raise e
