import { AsyncCommand } from "common/commands/command";
import { SignUpRequest } from "../dto";
import User from "common/models/user";
import { IAuthService } from "../service";

export class SignUpCommand implements AsyncCommand<SignUpRequest, User> {
  constructor(private readonly service: IAuthService) {}

  handle(request: SignUpRequest): Promise<User> {
    return this.service.signUp(request);
  }
}
