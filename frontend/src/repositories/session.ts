import { ApiSessionRepository } from "./api/session";
import { WailsSessionRepository } from "./wails/session";

export function createSessionRepository(): SessionRepository {
  if (import.meta.env.VITE_RUNTIME === "wails") {
    return new WailsSessionRepository()
  }

  return new ApiSessionRepository()
}

export interface SessionRepository {
  active(): Promise<any>
}