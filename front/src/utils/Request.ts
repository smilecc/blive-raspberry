import type { AxiosError, AxiosPromise } from "axios";
import axios from "axios";
import { useMessage } from "naive-ui";

export interface IResponse<T> {
  code: number;
  message: string;
  data: T;
}

export type Response<T = any> = AxiosPromise<IResponse<T>>;

export type ResponseError<T = any> = AxiosError<IResponse<T>>;

export const ToastErrorHandler = (error: ResponseError) => {
  if (error.isAxiosError) {
    console.error("请求错误", error.response);
    const message = useMessage();
    message.error(error.response?.data.message || "服务器走丢啦，请稍后再试");
  }
};

export const ModalErrorHandler = (error: ResponseError) => {
  if (error.isAxiosError) {
    console.error("请求错误", error.response?.data);
  }
};

axios.interceptors.request.use(async (config) => {
  if (import.meta.env.DEV) {
    config.baseURL = "http://127.0.0.1:18000";
  }

  return config;
});

axios.interceptors.response.use(
  (response) => {
    // console.log('收到请求响应', response);
    return response;
  },
  (error: AxiosError) => {
    return Promise.reject(error);
  }
);

export const Request = axios;
