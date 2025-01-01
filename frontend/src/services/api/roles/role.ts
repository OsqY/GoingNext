import { RolesResponse } from "@/types/roles/rolesGetDTO";
import { api } from "..";

export const roleService = {
  async getRoles() {
    try {
      const token = localStorage.getItem("token");
      const response = await api.get<RolesResponse>("/api/roles/all", {
        headers: { Authorization: `Bearer ${token}` },
      });

      return response;
    } catch (error) {
      console.error(error);
    }
  },
};
