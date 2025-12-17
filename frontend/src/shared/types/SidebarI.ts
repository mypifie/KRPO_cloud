import type { FunctionalComponent } from 'vue'

export interface SidebarItemI {
  title: string
  url: string
  icon: FunctionalComponent
}

export interface SideBarI {
  items: SidebarItemI[]
  username: string | undefined
  isAuth: boolean
}
