export class CommandRequest {

}

export interface Command<TRequest extends CommandRequest, TResponse> {
    handle(request: TRequest): TResponse
}

export interface AsyncCommand<TRequest extends CommandRequest, TResponse> {
    handle(request: TRequest): Promise<TResponse>
}