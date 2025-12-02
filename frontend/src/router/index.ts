import { createRouter, createWebHistory } from 'vue-router'
import RegistrationPage from "@/views/RegistrationPage.vue";
import LoginPage from "@/views/LoginPage.vue";
import LastFilesPage from '@/views/LastFilesPage.vue'
import MainPage from '@/views/MainPage.vue'
import ArchivePage from '@/views/ArchivePage.vue'
import MarkedPage from '@/views/MarkedPage.vue'
import SettingsPage from '@/views/SettingsPage.vue'

const routes = [
  { path: '/registration', component: RegistrationPage },
  { path: '/login', component: LoginPage },
  { path: '/last-files', component: LastFilesPage },
  { path: '/', component: MainPage },
  { path: '/archived-files', component: ArchivePage },
  { path: '/marked-files', component: MarkedPage },
  { path: '/settings', component: SettingsPage },
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: routes,
})


export default router
