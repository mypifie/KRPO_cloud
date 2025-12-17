<template>
  <Sidebar>
    <SidebarHeader>
      <Avatar v-if="isAuth">
        <AvatarImage src="#" />
        <AvatarFallback>{{initials}}</AvatarFallback>
      </Avatar>
    </SidebarHeader>
    <SidebarContent>
      <SidebarGroup>
        <SidebarGroupLabel>Облачное хранилище</SidebarGroupLabel>
        <SidebarGroupContent>
          <SidebarMenu>
            <Button>
              <Upload />
              <span>Загрузить</span>
            </Button>
            <SidebarMenuItem v-for="item in items" :key="item.title">
              <SidebarMenuButton asChild>
                <RouterLink :to="item.url">
                  <component :is="item.icon" />
                  <span>{{ item.title }}</span>
                </RouterLink>
              </SidebarMenuButton>
            </SidebarMenuItem>
          </SidebarMenu>
        </SidebarGroupContent>
      </SidebarGroup>
    </SidebarContent>
    <SidebarFooter>
      <SidebarMenuButton
        v-if="isAuth"
        @click="handleLogout"
      >
        <log-out-icon/>
        Выйти
      </SidebarMenuButton>
      <SidebarMenuButton v-else asChild>
        <RouterLink to="/login">
          <LogInIcon/>
          <span>Войти</span>
        </RouterLink>
      </SidebarMenuButton>
    </SidebarFooter>
  </Sidebar>
</template>

<script setup lang="ts">
import { Upload, LogInIcon, LogOutIcon } from 'lucide-vue-next'
import {
  Sidebar,
  SidebarHeader,
  SidebarContent,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarFooter,
} from '@/shared/components/ui/sidebar'
import { Button } from '@/shared/components/ui/button'
import { Avatar, AvatarFallback, AvatarImage } from '@/shared/components/ui/avatar'
import type { SideBarI } from '@/shared/types/SidebarI.ts'
import { computed } from 'vue'

const props = defineProps<SideBarI>();
const emit = defineEmits<{
  (e: 'logout'): void
}>();

const handleLogout = () => {
  emit('logout');
};

const initials = computed(() => props?.username.split(' ').slice(0,2).map(word => word[0].toUpperCase()).join(''))
</script>
