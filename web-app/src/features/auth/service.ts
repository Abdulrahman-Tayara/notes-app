import User from "common/models/user";
import { IApiClient } from "common/remote/apiClient";
import { SignUpRequest } from "./dto";
import { ApiResponse, NotesApiUrls } from "common/remote/contracts";

export interface IAuthService {
  signUp(request: SignUpRequest): Promise<User>;
}

export class AuthService implements IAuthService {
  constructor(
    private readonly apiClient: IApiClient,
    private readonly urlsContracts: NotesApiUrls
  ) {}

  async signUp(request: SignUpRequest): Promise<User> {
    const response = await this.apiClient.post<ApiResponse<User>>(this.urlsContracts.signUpUrl, request);

    await response.throwOnFail()

    const serializedResponse = await response.serialize()

    return serializedResponse.data
  }
}
