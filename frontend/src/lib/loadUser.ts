import { authService } from "@/services/api/auth/auth";
import { AppRouterInstance } from "next/dist/shared/lib/app-router-context.shared-runtime";

export const userCheck = {
  async loadUser(router: AppRouterInstance) {
    try {
      const currentUser = await authService.getCurrentUser();
      if (currentUser) {
        return currentUser;
      } else {
        router.push("/login");
      }
    } catch (error) {
      console.error("Failed to load user:", error);
      router.push("/login");
    } finally {
      return false;
    }
  },
};
