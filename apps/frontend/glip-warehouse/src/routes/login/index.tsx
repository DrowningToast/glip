"use client";

import type React from "react";
import { createFileRoute, redirect, useNavigate } from "@tanstack/react-router";
import { useMutation } from "@tanstack/react-query";
import { toast } from "react-hot-toast";

import { useState } from "react";
import { Link } from "@tanstack/react-router";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { EyeIcon, EyeOffIcon } from "lucide-react";

import { login, getUser } from "@/queries/auth/query";
import { LoginResponse } from "@/queries/auth/type";
import { APIError } from "@/libs/axiosClient";

export const Route = createFileRoute("/login/")({
  component: RouteComponent,
  loader: async () => {
    try {
      const user = await getUser();
      if (user) {
        return redirect({ to: "/" });
      }
      return null;
    } catch (error) {
      console.error(error);
      return null;
    }
  }
});

function RouteComponent() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [showPassword, setShowPassword] = useState(false);
  const navigate = useNavigate();
  const loginMutation = useMutation({
    mutationFn: login,
  })

  const handleEmailChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value;
    // Filter out special characters except @ and .
    const filteredValue = value.replace(/[^a-zA-Z0-9@.]/g, '');
    setEmail(filteredValue);
  };

  const handlePasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value;
    // Filter out special characters except basic alphanumeric
    const filteredValue = value.replace(/[^a-zA-Z0-9]/g, '');
    setPassword(filteredValue);
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    const id = toast.loading("กำลังเข้าสู่ระบบ...")

    loginMutation.mutateAsync({
      email,
      password,
    }).then((res: LoginResponse) => {
      toast.success(res.message, { id })
      navigate({ to: "/" })
    }).catch((err: APIError) => {
      toast.error(err.message, { id })
    })
  };

  return (
    <div className="flex min-h-screen items-center justify-center bg-gray-50 p-4">
      <Card className="w-full max-w-md">
        <CardHeader className="space-y-1">
          <CardTitle className="text-2xl font-bold">เข้าสู่ระบบ</CardTitle>
          <CardDescription>
            กรุณากรอกอีเมลและรหัสผ่านเพื่อเข้าสู่ระบบ
          </CardDescription>
        </CardHeader>
        <form onSubmit={handleSubmit}>
          <CardContent className="space-y-4">
            <div className="space-y-2">
              <Label htmlFor="email">อีเมล</Label>
              <Input
                id="email"
                type="email"
                placeholder="your.email@example.com"
                value={email}
                onChange={handleEmailChange}
                required
              />
            </div>
            <div className="space-y-2">
              <Label htmlFor="password">รหัสผ่าน</Label>
              <div className="relative">
                <Input
                  id="password"
                  type={showPassword ? "text" : "password"}
                  value={password}
                  onChange={handlePasswordChange}
                  required
                />
                <Button
                  type="button"
                  variant="ghost"
                  size="icon"
                  className="absolute right-0 top-0 h-full px-3 py-2 text-gray-400 hover:text-gray-600"
                  onClick={() => setShowPassword(!showPassword)}
                >
                  {showPassword ? (
                    <EyeOffIcon className="h-4 w-4" />
                  ) : (
                    <EyeIcon className="h-4 w-4" />
                  )}
                </Button>
              </div>
            </div>
          </CardContent>
          <CardFooter className="flex flex-col space-y-4 mt-4">
            <Button className="w-full" type="submit" disabled={loginMutation.isPending}>
              เข้าสู่ระบบ
            </Button>
            <p className="text-center text-sm">
              ยังไม่มีบัญชีผู้ใช้?{" "}
              <Link
                to="/register"
                className="font-semibold text-primary hover:underline"
              >
                สมัครสมาชิก
              </Link>
            </p>
          </CardFooter>
        </form>
      </Card>
    </div>
  );
}
