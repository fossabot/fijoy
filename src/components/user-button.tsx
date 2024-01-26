"use client";

import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { cn } from "@/lib/utils";
import { Button } from "@/components/ui/button";
import { env } from "@/env";
// import { useUser } from "@/hooks/use-user";
// import { useRouter } from "@tanstack/react-router";

function UserButton() {
  // const user = useUser();
  // const router = useRouter();
  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button
          variant="outline"
          className={cn(
            "h-8 justify-start gap-2 rounded-[0.5rem] bg-background text-sm font-normal  shadow-none",
          )}
        >
          <Avatar className="h-5 w-5">
            {/* <AvatarImage */}
            {/*   src={`https://avatar.vercel.sh/${session.user.email}.png`} */}
            {/*   alt={session.user.email} */}
            {/*   className="grayscale" */}
            {/* /> */}
            <AvatarFallback>SC</AvatarFallback>
          </Avatar>
          <div>{/* {session.user.email} / {workspace} */}</div>
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent>
        <DropdownMenuLabel>My Account</DropdownMenuLabel>
        <DropdownMenuSeparator />
        <DropdownMenuItem
          onSelect={async () => {
            // return router.navigate("/profile");
          }}
        >
          Profile
        </DropdownMenuItem>
        <DropdownMenuSeparator />
        <DropdownMenuItem
          onSelect={() => {
            window.location.replace(
              env.VITE_BACKEND_URL + "/auth/google/logout",
            );
          }}
        >
          Logout
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  );
}

export default UserButton;
