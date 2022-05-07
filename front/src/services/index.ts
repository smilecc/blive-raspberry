import { Request, Response } from "@/utils";
import { AxiosPromise } from "axios";

export namespace ConfigService {
  export interface IConfig {
    id: number;
    name: string;
    value: string;
  }

  export function getConfig(key: string): Response<IConfig> {
    return Request.get(`/api/config/${key}`);
  }

  export function setConfig(key: string, value: string): Response<IConfig> {
    return Request.put(`/api/config/${key}`, value);
  }
}

export namespace NeteaseService {
  export function sendCaptcha(host: string, phone: string): AxiosPromise {
    return Request.get(`${host}/captcha/sent`, {
      params: {
        phone,
      },
    });
  }

  export function loginByCaptcha(
    host: string,
    phone: string,
    captcha: string
  ): AxiosPromise<any> {
    return Request.get(`${host}/login/cellphone`, {
      params: {
        phone,
        captcha,
      },
    });
  }
}
