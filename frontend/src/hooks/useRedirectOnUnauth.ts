import { authService } from "@/services/api/auth/auth";
import { useRouter } from "next/navigation";
import { useEffect } from "react";

export function useRedirectOnUnauth(redirectTo: string = "/main/dashboard") {
  const router = useRouter();

  useEffect(() => {
    async function checkAuth() {
      const user = await authService.getCurrentUser();
      if (user) {
        router.push(redirectTo);
      }
    }
    checkAuth();
  }, [router, redirectTo]);
}
