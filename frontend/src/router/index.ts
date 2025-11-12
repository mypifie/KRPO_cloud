import { createRouter, createWebHistory } from 'vue-router'
import RegistrationPage from "@/views/RegistrationPage.vue";
import LoginPage from "@/views/LoginPage.vue";
import LastFilesPage from '@/views/LastFilesPage.vue'

const routes = [
  { path: '/login', component: RegistrationPage },
  { path: '/registration', component: LoginPage },
  { path: '/last-files', component: LastFilesPage },
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: routes,
})


export default router
