import axios from "axios";
import {baseUrl} from "@/shared/api/consts.ts";

const api = axios.create({
  baseURL: baseUrl,
  headers: {
    "Content-Type": "application/json",
  }
})

api.interceptors.request.use(
  (config) => {
    const sessionData = localStorage?.getItem('session');
    if (sessionData) {
      const accessToken = JSON.parse(sessionData)?.accessToken
      config.headers.Authorization = `Bearer  ${accessToken}`
    }
    return config;
  }
)

api.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => {
    if (error.response?.status === 401) {
      localStorage.removeItem('session');
      window.location.href = '/login';
    }

    return Promise.reject(error);
  }
);

export default api;
