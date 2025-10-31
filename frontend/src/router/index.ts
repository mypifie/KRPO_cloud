import { createRouter, createWebHistory } from 'vue-router'
import RegistrationPage from "@/views/RegistrationPage.vue";
import LoginPage from "@/views/LoginPage.vue";

const routes = [
  { path: '/', component: RegistrationPage },
  { path: '/reg', component: LoginPage },

]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: routes,
})


export default router
