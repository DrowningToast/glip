import axios from "axios";
import urlConfig from "@/config/endpoint.json";

export type APIError = {
    status: number;
    message: string;
}

const region = localStorage.getItem("region") || "USA1";
const endpoint = urlConfig.find((item) => item.region === region)?.endpoint;

const axiosClient = axios.create({
  baseURL: `${endpoint}`,
  withCredentials: true,
});

axiosClient.interceptors.response.use(
  (response) => response,
  (error) => {
    if (axios.isAxiosError(error)) {
      return Promise.reject({
        status: error.response?.status || 500,
        message: error.response?.data?.message || "เกิดข้อผิดพลาดที่ไม่ทราบสาเหตุ",
      });
    }
    return Promise.reject(error);
  }
);

export default axiosClient;