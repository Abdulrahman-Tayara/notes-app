import { Constructor, commandsContainer } from "di/containers"

export default function useCommand<T>(cons: Constructor<T>): T {
    const instance = commandsContainer.resolve<T>(cons)

    return instance
}