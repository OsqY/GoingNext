
import { UpdateUserDTO } from "@/types/auth/user"
import { api } from ".."

export const userService = {
  async updateUser(updateUser: UpdateUserDTO) {

    const response = await api.put(`/api/users/update/${updateUser.ID}`)

    if (response.status === 200) {
      return true
    }
    return false
  }
}
