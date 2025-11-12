<template>
  <Sidebar>
    <SidebarHeader>
      <Avatar>
        <AvatarImage src="#"/>
        <AvatarFallback>КР</AvatarFallback>
      </Avatar>
    </SidebarHeader>
    <SidebarContent>
      <SidebarGroup>
        <SidebarGroupLabel>Облачное хранилище</SidebarGroupLabel>
        <SidebarGroupContent>
          <SidebarMenu>
            <Button>
              <component :is="Upload"/>
              <span>Загрузить</span>
            </Button>
            <SidebarMenuItem v-for="item in items" :key="item.title">
              <SidebarMenuButton asChild>
                <a :href="item.url">
                  <component :is="item.icon" />
                  <span>{{item.title}}</span>
                </a>
              </SidebarMenuButton>
            </SidebarMenuItem>
          </SidebarMenu>
        </SidebarGroupContent>
      </SidebarGroup>
    </SidebarContent>
  </Sidebar>
</template>

<script setup lang="ts">
import { Settings, Files, Clock, Archive, Star, Upload } from "lucide-vue-next"
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
} from "@/shared/components/ui/sidebar"
import { Button } from '@/shared/components/ui/button'
import { Avatar, AvatarFallback, AvatarImage } from '@/shared/components/ui/avatar'

defineProps<>({})

const items = [
  {
    title: "Последние",
    url: "/last-files",
    icon: Clock,
  },
  {
    title: "Файлы",
    url: "#",
    icon: Files,
  },
  {
    title: "Архив",
    url: "#",
    icon: Archive,
  },
  {
    title: "Помеченные",
    url: "#",
    icon: Star,
  },
  {
    title: "Настройки",
    url: "#",
    icon: Settings,
  },
];
</script>
