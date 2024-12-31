"use client"

import { useForm } from "react-hook-form";
import { z } from "zod"
import { zodResolver } from "@hookform/resolvers/zod"
import { useEffect, useState } from "react";
import { UpdateUserDTO, UserInfo } from "@/types/auth/user";
import { authService } from "@/services/api/auth/auth";
import { userService } from "@/services/api/user/user";
import { Form, FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage } from "../ui/form";
import { Input } from "../ui/input";
import { Button } from "../ui/button";

const formSchema = z.object({
  email: z
    .string()
    .email("Invalid email address")
    .min(1, "Email is required")
    .max(256, "Email must be less than 256 characters"),
  username: z.string().min(3, "Minimum username length is 3 characters"),
  password: z.string().min(8, "Minimum password length is 8 characters"),
});

export function UserAccountForm() {
  const [userInfo, setUserInfo] = useState<UserInfo>();

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      email: userInfo?.Email,
      username: userInfo?.Username,
      password: userInfo?.Password,
    },
  });

  useEffect(() => {
    const fetchData = async () => {
      try {
        const data = await authService.getCurrentUser();
        setUserInfo(data);
        form.reset({
          email: data.Email,
          username: data.Username,
          password: data.Password,
        });
      } catch (error) {
        console.error(error);
      }
    };
    fetchData();
  }, [form]);

  async function onSubmit(values: z.infer<typeof formSchema>) {
    try {
      if (userInfo) {
        const updateUserDTO: UpdateUserDTO = {
          Email: values.email,
          Username: values.username,
          Password: values.password,
          ID: userInfo.ID,
          ImageUrl: userInfo.ImageUrl,
        };
        await userService.updateUser(updateUserDTO);
      }
    } catch (error) {
      console.error(error);
    }
  }

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
        <FormField
          control={form.control}
          name="email"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Email</FormLabel>
              <FormControl>
                <Input placeholder="email@example.com" {...field} />
              </FormControl>
              <FormDescription>Your email</FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="username"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Username</FormLabel>
              <FormControl>
                <Input placeholder="username" {...field} />
              </FormControl>
              <FormDescription>Your username</FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />
        <Button type="submit">Update</Button>
      </form>
    </Form>
  );
}

