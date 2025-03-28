"use client";

import type React from "react";

import { AlertCircle, Loader2 } from "lucide-react";
import { useState } from "react";
import { useNavigate } from "react-router";
import { toast } from "sonner";

import { Alert, AlertDescription } from "../../components/ui/alert";
import { Button } from "../../components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "../../components/ui/card";
import { Input } from "../../components/ui/input";
import { Label } from "../../components/ui/label";

export default function Page() {
  const [secretKey, setSecretKey] = useState("");
  const [error, setError] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const navigate = useNavigate();

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!secretKey.trim()) {
      setError("Authorization Secret Key is required");
      return;
    }

    setIsLoading(true);
    setError(null);

    const secret = import.meta.env.VITE_REGISTRY_SECRET_KEY;
    console.log(secret, secretKey);
    if (secretKey !== secret) {
      setError("Invalid secret key");
      setIsLoading(false);
      return;
    }
    localStorage.setItem("is_authenticated", "1");
    toast.success("Login successful");
    navigate("/");

    setIsLoading(false);
  };

  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-50 w-full">
      <Card className="w-full max-w-md">
        <CardHeader>
          <CardTitle className="text-2xl">Admin Login</CardTitle>
          <CardDescription>
            Enter your Authorization Secret Key to access the admin panel
          </CardDescription>
        </CardHeader>
        <form onSubmit={handleLogin}>
          <CardContent>
            {error && (
              <Alert variant="destructive" className="mb-4">
                <AlertCircle className="h-4 w-4" />
                <AlertDescription>{error}</AlertDescription>
              </Alert>
            )}
            <div className="space-y-4">
              <div className="space-y-2">
                <Label htmlFor="secretKey">Authorization Secret Key</Label>
                <Input
                  id="secretKey"
                  type="password"
                  placeholder="Enter your secret key"
                  value={secretKey}
                  onChange={(e) => setSecretKey(e.target.value)}
                  disabled={isLoading}
                  autoComplete="off"
                />
              </div>
            </div>
          </CardContent>
          <CardFooter className="mt-4">
            <Button type="submit" className="w-full" disabled={isLoading}>
              {isLoading ? (
                <>
                  <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                  Authenticating...
                </>
              ) : (
                "Login"
              )}
            </Button>
          </CardFooter>
        </form>
      </Card>
    </div>
  );
}
