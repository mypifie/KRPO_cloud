import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export const useUserStore = defineStore('user', () => {
  const isAuth = ref<boolean>(true);
  const name = ref<string>('');

  return { isAuth, name, }
})
