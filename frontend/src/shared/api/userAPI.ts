import type { AuthDataI, RegistrationDataI } from '@/shared/types/UserI.ts'
import axios from 'axios'
import { baseUrl } from '@/shared/api/consts.ts'

const userAPI = {
  register: (data: RegistrationDataI) => axios.post(`${baseUrl}/register`, data, { headers: { 'Content-Type': 'application/json' } }),
  auth: (data: AuthDataI) => axios.post(`${baseUrl}/login`, data, { headers: { 'Content-Type': 'application/json' } }),
};

export default userAPI;
