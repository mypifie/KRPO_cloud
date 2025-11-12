<template>
  <SidebarProvider v-if="isAuth">
    <AppSidebar />
    <SidebarTrigger />
    <main class="w-screen h-screen flex flex-row items-center justify-center">
      <RouterView />
    </main>
  </SidebarProvider>
  <main v-else class="w-screen h-screen flex flex-row items-center justify-center">
    <RouterView />
  </main>
</template>

<script setup lang="ts">
import AppSidebar from '@/shared/components/AppSidebar.vue'
import { SidebarProvider, SidebarTrigger } from '@/shared/components/ui/sidebar'
import { useUserStore } from '@/stores/user.ts'

const {isAuth} = useUserStore();
</script>
