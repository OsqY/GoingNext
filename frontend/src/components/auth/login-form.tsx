import { cn } from "@/lib/utils";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "./ui/card";
import { Label } from "./ui/label";
import { Input } from "./ui/input";
import { Button } from "./ui/button";
import { useState } from "react";
import { LoginDTO } from "@/types/auth/loginDTO";
import { authService } from "@/services/api/auth";

export function LoginForm({ className, ...props }: React.ComponentPropsWithoutRef<"div">) {
  const [loginCredentials, setLoginCredentials] = useState<LoginDTO>({ email: "", password: "" })

  return (
    <div className={cn("flex flex-col gap-9", className)}{...props}>
      <Card>
        <CardHeader className="text-center">
          <CardTitle className="text-xl">Welcome!</CardTitle>
          <CardDescription>
            Login with your account
          </CardDescription>
        </CardHeader>
        <CardContent>
          <form>
            <div className="grid gap-6">
              <div className="grid gap-2">
                <Label htmlFor="email">Email</Label>
                <Input id="email" type="email"
                  onChange={(e) => setLoginCredentials({ ...loginCredentials, email: e.target.value })}
                  placeholder="email@example.com"
                  required />
              </div>
              <div className="grid gap-2">
                <div className="flex items-center">
                  <Label htmlFor="password">Password</Label>

                </div>
                <Input id="password" type="password" onChange={(e) => setLoginCredentials({ ...loginCredentials, password: e.target.value })} />
              </div>
              <Button type="submit" className="w-ful" onClick={() => {
                authService.login(loginCredentials)
              }}>Login</Button>
            </div>
          </form>
        </CardContent>
      </Card>
    </div>
  )
}
