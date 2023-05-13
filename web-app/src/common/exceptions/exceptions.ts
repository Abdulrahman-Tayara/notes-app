export class HTTPException extends Error {
    constructor(readonly code: number, message?: string) {
        super(message ?? `HTTPException with code ${code}`);
    }
}