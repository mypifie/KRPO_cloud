import json
import fastapi
from fastapi import FastAPI, Request, Depends
from fastapi.middleware.cors import CORSMiddleware
from contextlib import asynccontextmanager

# Импортируем роутеры
import Router_Auth
from database_main import create_tables


@asynccontextmanager
async def lifespan(app: FastAPI):
    # При запуске приложения
    print("Создание таблиц...")
    await create_tables()
    print("Таблицы созданы!")
    yield
    # При остановке приложения
    print("Приложение остановлено")


app = FastAPI(redirect_slashes=False, lifespan=lifespan)

# Включаем роутеры
app.include_router(Router_Auth.router)

# CORS middleware
app.add_middleware(
    CORSMiddleware,
    allow_origins=[
        "*",
    ],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
    expose_headers=["*"],
)


@app.get("/")
async def root():
    return {"message": "Main page"}


@app.get("/protected")
async def protected_route(current_user=Depends(Router_Auth.get_current_user_db)):
    return {"message": "This is a protected route", "user": current_user}
