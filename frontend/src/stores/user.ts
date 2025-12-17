import { defineStore } from 'pinia'
import { ref} from 'vue'
import userAPI from '@/shared/api/userAPI.ts'
import type { AuthDataI, RegistrationDataI, UserI } from '@/shared/types/UserI.ts'
import router from '@/router'

export const useUserStore = defineStore('user', () => {
  const isAuth = ref<boolean>(false);
  const userData = ref<UserI | null>(null);

  const register = async (info: RegistrationDataI) => {
    try {
      const response = await userAPI.register(info)
      if (response.status === 200) {
        const {access_token, user} = response.data
        userData.value = user;
        localStorage.setItem('session', JSON.stringify({ access_token }));
        isAuth.value = true;
        await router.push('/');
      }
    } catch (e) {
      console.error(e);
    }
  }

  const auth = async (info: AuthDataI) => {
    console.log(info)
    try {
      const response = await userAPI.auth(info)
      if (response.status === 200) {
        const {access_token, user} = response.data
        userData.value = user;
        localStorage.setItem('session', `{ access_token: ${access_token} }`);
        isAuth.value = true;
        await router.push('/');
      }
    } catch (e) {
      console.error(e);
    }
  }

  const logOut = () => {
    localStorage.removeItem('session')
    isAuth.value = false;
    userData.value = null;
  }

  return {
    isAuth,
    userData,
    logOut,
    auth,
    register
  }
}, {
  persist: true
})
