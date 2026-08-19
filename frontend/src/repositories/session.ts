import { ApiSessionRepository } from "./api/session";
import { WailsSessionRepository } from "./wails/session";
import type { Result } from ".";

export function createSessionRepository(): SessionRepository {
  if (import.meta.env.VITE_RUNTIME === "wails") {
    return new WailsSessionRepository()
  }

  return new ApiSessionRepository()
}

export interface SessionRepository {
  active(): Promise<Result<any>>
  connect(con: any): Promise<Result<any>>
  disconnect(): Promise<Result<void>>
}