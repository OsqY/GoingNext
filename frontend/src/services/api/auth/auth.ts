
import { LoginDTO } from "@/types/auth/loginDTO";
import { api } from "../";
import { UserInfo } from "@/types/auth/user";


export const authService = {
  async login(loginInfo: LoginDTO) {
    const response = await api.post('/login', loginInfo)
    localStorage.setItem('token', response.data)
    return response
  },
  async getCurrentUser() {
    const token = localStorage.getItem("token")
    const response = await api.get<UserInfo>("/api/users/current", { headers: { "Authorization": `Bearer ${token}` } })
    return response.data
  }
}
