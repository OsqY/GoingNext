import { UpdateUserDTO } from "@/types/auth/user";
import Cookies from "js-cookie";
import { api } from "..";

export const userService = {
  async updateUser(updateUser: UpdateUserDTO) {
    const token = Cookies.get("token");

    const response = await api.put(
      `/api/users/update/${updateUser.ID}`,
      updateUser,
      { headers: { Authorization: `Bearer ${token}` } },
    );

    if (response.status === 200) {
      return true;
    }
    return false;
  },
};
