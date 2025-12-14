from sqlalchemy import (
    Table,
    Column,
    Integer,
    String,
    MetaData,
    DateTime,
    Date,
    Time,
    ForeignKey,
    text,
)
from sqlalchemy.orm import Mapped, mapped_column, relationship, DeclarativeBase
from datetime import datetime, date, time


class Base(DeclarativeBase):
    pass


metadata_obj = MetaData()


class AuthUsers(Base):
    __tablename__ = "auth_users"

    id: Mapped[int] = mapped_column(primary_key=True)
    email: Mapped[str]
    password_hash: Mapped[str]
    is_email_verified: Mapped[bool] = mapped_column(default=False)


class UserProfiles(Base):
    __tablename__ = "user_profiles"

    id: Mapped[int] = mapped_column(
        ForeignKey("auth_users.id"), primary_key=True, index=True
    )
    username: Mapped[str]
    is_admin: Mapped[bool] = mapped_column(default=False)
    is_blocked: Mapped[bool] = mapped_column(default=False)


# class Lessons(Base):
#     __tablename__ = "lessons"

#     id: Mapped[int] = mapped_column(primary_key=True)
#     user_id: Mapped[int] = mapped_column(ForeignKey("users.id"), index=True)

#     teacher_name: Mapped[str]
#     subject_name: Mapped[str]
#     room: Mapped[str]

#     startDate: Mapped[date] = mapped_column(Date, index=True)
#     endDate: Mapped[date] = mapped_column(Date)

#     startTime: Mapped[time] = mapped_column(Time)
#     endTime: Mapped[time] = mapped_column(Time)

#     repeat: Mapped[str]
#     dayOfWeek: Mapped[str]

#     user: Mapped["Users"] = relationship(back_populates="lessons")
#     comments: Mapped[list["Comments"]] = relationship(back_populates="lesson")
#     attendance: Mapped[list["Attendance"]] = relationship(back_populates="lesson")


# class Comments(Base):
#     __tablename__ = "comments"

#     id: Mapped[int] = mapped_column(primary_key=True)
#     lesson_id: Mapped[int] = mapped_column(ForeignKey("lessons.id"))
#     lesson_date: Mapped[date] = mapped_column(Date)
#     text: Mapped[str]

#     lesson: Mapped["Lessons"] = relationship(back_populates="comments")


# class Attendance(Base):
#     __tablename__ = "attendance"

#     id: Mapped[int] = mapped_column(primary_key=True)
#     lesson_id: Mapped[int] = mapped_column(ForeignKey("lessons.id"))
#     lesson_date: Mapped[date] = mapped_column(Date)

#     lesson: Mapped["Lessons"] = relationship(back_populates="attendance")
