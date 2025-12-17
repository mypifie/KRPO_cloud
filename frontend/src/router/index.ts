import { createRouter, createWebHistory } from 'vue-router'
import RegistrationPage from "@/views/RegistrationPage.vue";
import LoginPage from "@/views/LoginPage.vue";
import LastFilesPage from '@/views/LastFilesPage.vue'
import MainPage from '@/views/MainPage.vue'
import ArchivePage from '@/views/ArchivePage.vue'
import MarkedPage from '@/views/MarkedPage.vue'
import SettingsPage from '@/views/SettingsPage.vue'
import { useUserStore } from '@/stores/user.ts'

const routes = [
  { path: '/registration', name: 'Registration', component: RegistrationPage },
  { path: '/login', name: 'Login', component: LoginPage },
  { path: '/last-files', component: LastFilesPage },
  { path: '/', name: 'Main', component: MainPage },
  { path: '/archived-files', component: ArchivePage },
  { path: '/marked-files', component: MarkedPage },
  { path: '/settings', component: SettingsPage },
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: routes,
})

router.beforeEach((to, from) => {
  const userStore = useUserStore()
  if (
    userStore.isAuth &&
    ( to.name === 'Login' ||
    to.name === 'Registration' )
  ) {
    return { name: 'Main' }
  }
})

export default router
