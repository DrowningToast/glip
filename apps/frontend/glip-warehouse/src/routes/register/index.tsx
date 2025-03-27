import type React from "react";
import { createFileRoute, redirect, useNavigate } from "@tanstack/react-router";
import { useMutation } from "@tanstack/react-query";
import { toast } from "react-hot-toast";

import { useState } from "react"
import { Link } from "@tanstack/react-router"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { EyeIcon, EyeOffIcon } from "lucide-react"

import { getUser, register } from "@/queries/auth/query";
import { RegisterResponse } from "@/queries/auth/type";
import { APIError } from "@/libs/axiosClient";

export const Route = createFileRoute('/register/')({
  component: RouteComponent,
  loader: async () => {
    try {
      const user = await getUser();
      if (user) {
        return redirect({ to: "/login" });
      }
      return null;
    } catch (error) {
      console.error(error);
      return null;
    }
  }
})

function RouteComponent() {
  const navigate = useNavigate();
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")
  const [confirmPassword, setConfirmPassword] = useState("")
  const [showPassword, setShowPassword] = useState(false)
  const [passwordError, setPasswordError] = useState("")

  const registerMutation = useMutation({
    mutationFn: register,
  })

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()

    // Validate passwords match
    if (password !== confirmPassword) {
      setPasswordError("รหัสผ่านไม่ตรงกัน")
      return
    }

    setPasswordError("")

    const id = toast.loading("กำลังสมัครสมาชิก...")

    registerMutation.mutateAsync({
      email,
      password,
    }).then((res: RegisterResponse) => {
      toast.success(res.message, { id })
      navigate({ to: "/login" })
    }).catch((err: APIError) => {
      toast.error(err.message, { id })
    })
  }

  return (
    <div className="flex min-h-screen items-center justify-center bg-gray-50 p-4">
      <Card className="w-full max-w-md">
        <CardHeader className="space-y-1">
          <CardTitle className="text-2xl font-bold">สมัครสมาชิก</CardTitle>
          <CardDescription>กรุณากรอกข้อมูลเพื่อสมัครสมาชิกใหม่</CardDescription>
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
                onChange={(e) => setEmail(e.target.value)}
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
                  onChange={(e) => setPassword(e.target.value)}
                  required
                />
                <Button
                  type="button"
                  variant="ghost"
                  size="icon"
                  className="absolute right-0 top-0 h-full px-3 py-2 text-gray-400 hover:text-gray-600"
                  onClick={() => setShowPassword(!showPassword)}
                >
                  {showPassword ? <EyeOffIcon className="h-4 w-4" /> : <EyeIcon className="h-4 w-4" />}
                </Button>
              </div>
            </div>
            <div className="space-y-2">
              <Label htmlFor="confirm-password">ยืนยันรหัสผ่าน</Label>
              <Input
                id="confirm-password"
                type={showPassword ? "text" : "password"}
                value={confirmPassword}
                onChange={(e) => setConfirmPassword(e.target.value)}
                required
              />
              {passwordError && <p className="text-sm text-red-500">{passwordError}</p>}
            </div>
          </CardContent>
          <CardFooter className="flex flex-col space-y-4 mt-4">
            <Button className="w-full" type="submit" disabled={registerMutation.isPending}>
              สมัครสมาชิก
            </Button>
            <p className="text-center text-sm">
              มีบัญชีผู้ใช้แล้ว?{" "}
              <Link to="/login" className="font-semibold text-primary hover:underline">
                เข้าสู่ระบบ
              </Link>
            </p>
          </CardFooter>
        </form>
      </Card>
    </div>
  )
}
