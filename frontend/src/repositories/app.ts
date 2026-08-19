import type { Result } from ".";
import { ApiAppRepository } from "./api/app";
import { WailsAppRepository } from "./wails/app";

export interface AppRepository {
  version(): Promise<Result<string>>
}

export function createAppRepository(): AppRepository {
  if (import.meta.env.VITE_RUNTIME === "wails") {
    return new WailsAppRepository()
  }

  return new ApiAppRepository()
}