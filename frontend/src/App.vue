<template>
  <SidebarProvider>
    <AppSidebar
      :username="userStore.userData?.username"
      :items="items"
      :is-auth="userStore.isAuth"
      @logout="userStore.logOut"
    />
    <SidebarTrigger />
    <main class="w-screen h-screen flex flex-row items-center justify-center">
      <RouterView />
    </main>
  </SidebarProvider>
</template>

<script setup lang="ts">
import AppSidebar from '@/shared/components/AppSidebar.vue'
import { SidebarProvider, SidebarTrigger } from '@/shared/components/ui/sidebar'
import { useUserStore } from '@/stores/user.ts'
import { Archive, Clock, Files, Settings, Star } from 'lucide-vue-next'

const userStore = useUserStore()

const items = [
  {
    title: 'Последние',
    url: '/last-files',
    icon: Clock,
  },
  {
    title: 'Файлы',
    url: '/',
    icon: Files,
  },
  {
    title: 'Архив',
    url: '/archived-files',
    icon: Archive,
  },
  {
    title: 'Помеченные',
    url: '/marked-files',
    icon: Star,
  },
  {
    title: 'Настройки',
    url: '/settings',
    icon: Settings,
  },
]

</script>
