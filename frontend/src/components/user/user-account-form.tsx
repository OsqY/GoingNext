"use client";

import { useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { useEffect, useState } from "react";
import { UpdateUserDTO, UserInfo } from "@/types/auth/user";
import { authService } from "@/services/api/auth/auth";
import { userService } from "@/services/api/users/user";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "../ui/form";
import { Input } from "../ui/input";
import { Button } from "../ui/button";
import { roleService } from "@/services/api/roles/role";
import { RolesResponse } from "@/types/roles/rolesGetDTO";
import { Select } from "../ui/select";
import {
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@radix-ui/react-select";

import Image from "next/image";
import { filesService } from "@/services/api/files/files";

const MAX_FILE_SIZE = 5000000;
const ACCEPTED_IMAGE_TYPES = [
  "image/jpeg",
  "image/jpg",
  "image/png",
  "image/webp",
];

const formSchema = z.object({
  email: z
    .string()
    .email("Invalid email address")
    .min(1, "Email is required")
    .max(256, "Email must be less than 256 characters"),
  username: z.string().min(3, "Minimum username length is 3 characters"),
  password: z.string().min(8, "Minimum password length is 8 characters"),
  roleId: z.string().min(1, "Role is required"),
  image: z
    .any()
    .refine((files) => files?.length == 1, "Image is required.")
    .refine(
      (files) => files?.[0]?.size <= MAX_FILE_SIZE,
      `Max file size is 5MB.`,
    )
    .refine(
      (files) => ACCEPTED_IMAGE_TYPES.includes(files?.[0]?.type),
      ".jpg, .jpeg, .png and .webp files are accepted.",
    ),
});

export function UserAccountForm() {
  const [userInfo, setUserInfo] = useState<UserInfo>();
  const [roles, setRoles] = useState<RolesResponse>([]);
  const [imagePreview, setImagePreview] = useState<string | null>(null);

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      email: userInfo?.Email,
      username: userInfo?.Username,
      password: userInfo?.Password,
      roleId: userInfo?.RoleID.toString(),
      image: undefined,
    },
  });

  useEffect(() => {
    const fetchData = async () => {
      try {
        const data = await authService.getCurrentUser();
        setUserInfo(data!.data);
        form.reset({
          email: data!.data.Email,
          username: data!.data.Username,
          password: data!.data.Password,
          roleId: data!.data.RoleID.toString(),
        });
        const rolesData = await roleService.getRoles();
        if (rolesData?.data) {
          setRoles(rolesData.data);
        }
      } catch (error) {
        console.error(error);
      }
    };
    fetchData();
  }, [form]);

  async function onSubmit(values: z.infer<typeof formSchema>) {
    try {
      let imageUrl = userInfo?.ImageUrl ?? "";
      if (values.image && values.image[0]) {
        const uploadedUrl = await filesService.sendFileToS3(values.image[0]);
        if (uploadedUrl) {
          imageUrl = uploadedUrl;
        }
      }

      if (userInfo) {
        const updateUserDTO: UpdateUserDTO = {
          Email: values.email,
          Username: values.username,
          Password: values.password,
          RoleID: +values.roleId,
          ID: userInfo.ID,
          ImageUrl: imageUrl,
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
          name="image"
          render={({ field: { onChange, value, ...rest } }) => (
            <FormItem>
              <FormLabel>Profile picture</FormLabel>
              <FormControl>
                <Input
                  type="file"
                  accept={ACCEPTED_IMAGE_TYPES.join(".")}
                  onChange={(e) => {
                    const file = e.target.files?.[0];

                    if (file) {
                      onChange(e.target.files);
                      const reader = new FileReader();
                      reader.onloadend = () => {
                        setImagePreview(reader.result as string);
                      };

                      reader.readAsDataURL(file);
                    }
                  }}
                  {...rest}
                />
              </FormControl>
              {imagePreview && (
                <div className="mt-2">
                  <Image
                    src={imagePreview}
                    alt="Preview"
                    width={100}
                    height={100}
                    className="rounded-full"
                  />
                </div>
              )}
              <FormDescription>Your picture</FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />

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
        <FormField
          control={form.control}
          name="password"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Password</FormLabel>
              <FormControl>
                <Input type="password" placeholder="*********" {...field} />
              </FormControl>
              <FormDescription>
                Enter a new password to change it
              </FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="roleId"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Role</FormLabel>
              <Select onValueChange={field.onChange} defaultValue={field.value}>
                <FormControl>
                  <SelectTrigger>
                    <SelectValue placeholder="Select a role" />
                  </SelectTrigger>
                </FormControl>
                <SelectContent>
                  {roles!.map((role) => (
                    <SelectItem key={role.ID} value={role.ID.toString()}>
                      {role.Name}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
              <FormDescription>Your role </FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />
        <Button type="submit">Update</Button>
      </form>
    </Form>
  );
}
