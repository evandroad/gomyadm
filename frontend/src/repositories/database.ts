import { WailsConnectionRepository } from "./wails/connection";
import { ApiConnectionRepository } from "./api/connection";

export function createDatabaseRepository(): DatabaseRepository {
  if (import.meta.env.VITE_RUNTIME === "wails") {
    return new WailsConnectionRepository()
  }

  return new ApiConnectionRepository()
}

export interface DatabaseRepository {
  getAll(): Promise<any>
}