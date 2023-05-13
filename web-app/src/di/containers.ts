import { ContainerFactory, Constructor as libConstructor } from "get-it-di";
import { ApiClient, IApiClient } from "common/remote/apiClient";
import { AuthService, IAuthService } from "features/auth/service";
import { NotesApiUrls, urls } from "common/remote/contracts";
import { SignUpCommand } from "features/auth/signup/commands";

export declare type Constructor<T> = libConstructor<T>;

const commandsContainer = ContainerFactory.create();
const servicesContainer = ContainerFactory.create();
const commonContainer = ContainerFactory.create();

const setupCommon = () => {
  const serverUrl = process.env.REACT_APP_API_URL ?? "";

  commonContainer.register<IApiClient>("API_CLIENT", () => new ApiClient(serverUrl));
  commonContainer.registerSingleton<NotesApiUrls>(
    "API_URLS",
    urls,
  );
};

const setupServices = () => {
  servicesContainer.registerLazySingleton<IAuthService>(
    "AUTH_SERVICE",
    () =>
      new AuthService(
        commonContainer.resolve<IApiClient>("API_CLIENT"),
        commonContainer.resolve<NotesApiUrls>("API_URLS")
      )
  );
};

const setupCommands = () => {
  commandsContainer.register<SignUpCommand>(
    SignUpCommand,
    () =>
      new SignUpCommand(servicesContainer.resolve<IAuthService>("AUTH_SERVICE"))
  );
};

const setup = () => {
  setupCommon();
  setupServices();
  setupCommands();
};

export { setup, commandsContainer, servicesContainer };
