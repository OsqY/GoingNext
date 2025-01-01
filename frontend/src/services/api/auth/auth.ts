import { LoginDTO } from "@/types/auth/loginDTO";
import { api } from "../";
import { UserInfo } from "@/types/auth/user";
import Cookies from "js-cookie";

export const authService = {
  async login(loginInfo: LoginDTO) {
    try {
      const response = await api.post("/login", loginInfo);

      if (response.status === 200 && response.data) {
        Cookies.set("token", response.data, { expires: 1 });

        return true;
      }
    } catch (error) {
      return false;
    }
  },
  async getCurrentUser() {
    try {
      const token = Cookies.get("token");
      if (!token) {
        return null;
      }
      const response = await api.get<UserInfo>("/api/users/current", {
        headers: { Authorization: `Bearer ${token}` },
      });
      return response;
    } catch (error) {
      console.error("Get current user error:", error);
      return null;
    }
  },
  logout() {
    Cookies.remove("token");
  },
};
