import axios from "axios";
import {baseUrl} from "@/shared/api/consts.ts";

const api = axios.create({
  baseURL: baseUrl,
  headers: {
    "Content-Type": "application/json",
  }
})

api.interceptors.request.use(
  (response) => response,
  (error) => {}
)

api.interceptors.response.use()

export default api;
