<template>
  <Card class="w-[350px]">
    <CardHeader>
      <CardTitle>Регистрация</CardTitle>
      <CardDescription>Введите данные для регистрации</CardDescription>
    </CardHeader>
    <CardContent>
      <form>
        <div class="grid items-center w-full gap-4">
          <div class="flex flex-col space-y-1.5">
            <Label for="name">ФИО:</Label>
            <Input id="name" placeholder="Введите ваши ФИО" />
          </div>
          <div class="flex flex-col space-y-1.5">
            <Label for="name">E-mail:</Label>
            <Input id="name" placeholder="Введите ваш e-mail" />
          </div>
          <div class="flex flex-col space-y-1.5">
            <Label for="name">Пароль:</Label>
            <Input id="name" type="password" placeholder="Введите пароль" />
          </div>
          <div class="flex flex-col space-y-1.5">
            <span>Уже есть аккаунт? <RouterLink class="text-blue-400" to="/reg">Войти</RouterLink></span>
          </div>
        </div>
      </form>
    </CardContent>
    <CardFooter class="flex justify-between px-6 pb-6">
      <Button variant="outline">
        Отмена
      </Button>
      <Button>Регистрация</Button>
    </CardFooter>
  </Card>
</template>

<script setup lang='ts'>
import { Button } from "@/shared/components/ui/button"
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/shared/components/ui/card"
import { Input } from "@/shared/components/ui/input"
import {Label} from "@/shared/components/ui/label";
</script>
