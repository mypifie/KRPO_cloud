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
            <Input v-model="registrationData.username" id="name" placeholder="Введите ваши ФИО" />
          </div>
          <div class="flex flex-col space-y-1.5">
            <Label for="email">E-mail:</Label>
            <Input v-model="registrationData.email" id="email" placeholder="Введите ваш e-mail" />
          </div>
          <div class="flex flex-col space-y-1.5">
            <Label for="password">Пароль:</Label>
            <Input
              v-model="registrationData.password"
              id="password"
              type="password"
              placeholder="Введите пароль"
            />
          </div>
          <div class="flex flex-col space-y-1.5">
            <span
              >Уже есть аккаунт?
              <RouterLink class="text-blue-400" to="/login">Войти</RouterLink></span
            >
          </div>
        </div>
      </form>
    </CardContent>
    <CardFooter class="flex justify-end px-6 pb-6">
      <Button @click="handleRegister">Регистрация</Button>
    </CardFooter>
  </Card>
</template>

<script setup lang="ts">
import { Button } from '@/shared/components/ui/button'
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/shared/components/ui/card'
import { Input } from '@/shared/components/ui/input'
import { Label } from '@/shared/components/ui/label'
import { useUserStore } from '@/stores/user.ts'
import { reactive } from 'vue'
import type { RegistrationDataI } from '@/shared/types/UserI.ts'

const userStore = useUserStore()

const registrationData = reactive<RegistrationDataI>({
  username: '',
  email: '',
  password: '',
})

const handleRegister = async () => {
  await userStore.register(registrationData)
}
</script>
