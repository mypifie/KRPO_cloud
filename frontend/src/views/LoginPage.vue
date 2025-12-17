<template>
  <Card class="w-[350px]">
    <CardHeader>
      <CardTitle>Вход</CardTitle>
      <CardDescription>Введите ваши данные для входа</CardDescription>
    </CardHeader>
    <CardContent>
      <form>
        <div class="grid items-center w-full gap-4">
          <div class="flex flex-col space-y-1.5">
            <Label for="email">E-mail:</Label>
            <Input v-model="authData.email" id="email" placeholder="Введите ваш e-mail" />
          </div>
          <div class="flex flex-col space-y-1.5">
            <Label for="current-password">Пароль:</Label>
            <Input
              v-model="authData.password"
              id="current-password"
              type="password"
              placeholder="Введите пароль"
            />
          </div>
          <div class="flex flex-col space-y-1.5">
            <span
              >Нет аккаунта?
              <RouterLink class="text-blue-400" to="/registration"
                >Зарегистрироваться</RouterLink
              ></span
            >
          </div>
        </div>
      </form>
    </CardContent>
    <CardFooter class="flex justify-end px-6 pb-6">
      <Button @click="handleLogIn">Вход</Button>
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
import type { AuthDataI } from '@/shared/types/UserI.ts'
import { reactive } from 'vue'

const userStore = useUserStore()

const authData = reactive<AuthDataI>({
  email: '',
  password: '',
})

const handleLogIn = async () => {
  await userStore.auth(authData);
}
</script>
