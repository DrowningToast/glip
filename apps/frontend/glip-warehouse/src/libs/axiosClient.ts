import axios from "axios";

export type APIError = {
    status: number;
    message: string;
}

const axiosClient = axios.create({
  baseURL: `${import.meta.env.VITE_API_URL}`,
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