"use client";

import { authService } from "@/services/api/auth/auth";
import { Sidebar, SidebarContent, SidebarHeader } from "../ui/sidebar";
import { NavUser } from "./nav-user";
import { useEffect, useState } from "react";
import { User } from "@/types/auth/currentUser";
import { NavMain } from "./nav-main";
import { Network } from "lucide-react";

const routes = {
  navMain: [
    {
      title: "Administration",
      url: "#",
      icon: Network,
      isActive: true,
      items: [
        {
          title: "Users",
          url: "#",
        },
        {
          title: "Email",
          url: "#",
        },
        {
          title: "Notifications",
          url: "#",
        },
      ],
    },
  ],
  administration: [],
};

export function AppSidebar({ ...props }: React.ComponentProps<typeof Sidebar>) {
  const [user, setUser] = useState<User>();

  useEffect(() => {
    const fetchData = async () => {
      const user = await authService.getCurrentUser();
      setUser(user!.data);
    };
    fetchData();
  }, []);

  return (
    <Sidebar collapsible="icon" {...props}>
      <SidebarHeader>{user && <NavUser user={user} />}</SidebarHeader>
      <SidebarContent>
        <NavMain items={routes.navMain} />
      </SidebarContent>
    </Sidebar>
  );
}
