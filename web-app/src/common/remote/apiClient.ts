import axios, { AxiosError } from "axios";
import { AxiosResponse } from "axios";
import { HTTPException } from "common/exceptions/exceptions";

interface Response<T> {
  status: number;
  ok: boolean;

  throwOnFail(errorMessageKey?: string): Promise<void>;
  serialize(): Promise<T>;
}

export interface IApiClient {
  post<T>(url: string, body?: any): Promise<Response<T>>;
}

export class ApiClient implements IApiClient {
  constructor(private readonly baseUrl?: string) {}
  private constructResponse<T>(libResponse: AxiosResponse<T>): Response<T> {
    const isOk = libResponse.status < 300;
    return {
      status: libResponse.status,
      ok: isOk,
      throwOnFail: async (errorMessageKey?: string) => {
        if (!isOk) {
          const jsonResponse = libResponse.data as any;
          const errorMessage = jsonResponse[errorMessageKey ?? "error"];
          throw new HTTPException(libResponse.status, errorMessage);
        }
      },
      serialize: async (): Promise<T> => {
        const jsonResponse = libResponse.data;
        return jsonResponse;
      },
    };
  }

  private constructErrorResponse<T>(error: AxiosError): Response<T> {
    return {
      status: error.status ?? 0,
      ok: false,
      throwOnFail: async (errorMessageKey?: string) => {
        const jsonResponse = error.response && (error.response as any);
        const errorMessage = jsonResponse?.[errorMessageKey ?? "error"];

        throw new HTTPException(error.status ?? 0, errorMessage);
      },
      serialize: async (): Promise<T> => {
        const jsonResponse = error.response as T;
        return jsonResponse;
      },
    };
  }

  private constructUrl(url: string): string {
    return `${this.baseUrl ?? ""}${
      this.baseUrl?.endsWith("/") || url.startsWith("/") ? "" : "/"
    }${url}`;
  }

  async post<T>(url: string, body?: any): Promise<Response<T>> {
    try {
      const axiosResponse = await axios<T>({
        method: "post",
        url: this.constructUrl(url),
        data: body,
      });
      return this.constructResponse<T>(axiosResponse);
    } catch (error) {
      if (error instanceof AxiosError) {
        return this.constructErrorResponse<T>(error);
      } else {
        throw error;
      }
    }
  }
}
