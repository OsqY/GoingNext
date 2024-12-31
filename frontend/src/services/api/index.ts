import { config } from "@/config/env"
import axios from "axios"

export const api = axios.create({
  baseURL: config.API_URL, timeout: 10000, headers: {
    'Content-Type': 'application/json'
  }
})
