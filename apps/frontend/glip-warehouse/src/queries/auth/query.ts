import axiosClient from "@/libs/axiosClient";

import { LoginData, LoginResponse, RegisterData, RegisterResponse, User } from "./type";

export const register = async (data: RegisterData): Promise<RegisterResponse> => {
    const response = await axiosClient.post("/auth/register", data);
    return response.data;
};

export const login = async (data: LoginData): Promise<LoginResponse> => {
    const response = await axiosClient.post("/auth/login", data);
    return response.data;
};

export const getUser = async (): Promise<User> => {
    const response = await axiosClient.get("/auth/me");
    return response.data;
};

export const logout = async (): Promise<void> => {
    const response = await axiosClient.post("/auth/logout");
    return response.data;
};
