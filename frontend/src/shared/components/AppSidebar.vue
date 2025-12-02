<template>
  <Sidebar>
    <SidebarHeader>
      <Avatar>
        <AvatarImage src="#" />
        <AvatarFallback>КР</AvatarFallback>
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
  </Sidebar>
</template>

<script setup lang="ts">
import { Upload } from 'lucide-vue-next'
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
} from '@/shared/components/ui/sidebar'
import { Button } from '@/shared/components/ui/button'
import { Avatar, AvatarFallback, AvatarImage } from '@/shared/components/ui/avatar'
import type { SideBarI } from '@/shared/types/SidebarI.ts'

defineProps<SideBarI>()
</script>
