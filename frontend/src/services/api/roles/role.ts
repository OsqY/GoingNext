import { RolesResponse } from "@/types/roles/rolesGetDTO";
import { api } from "..";
import Cookies from "js-cookie";

export const roleService = {
  async getRoles() {
    try {
      const token = Cookies.get("token");
      const response = await api.get<RolesResponse>("/api/roles/all", {
        headers: { Authorization: `Bearer ${token}` },
      });

      return response;
    } catch (error) {
      console.error(error);
    }
  },
};
