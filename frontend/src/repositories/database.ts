import { WailsDatabaseRepository } from "./wails/database";
import { ApiDatabaseRepository } from "./api/database";

export function createDatabaseRepository(): DatabaseRepository {
  if (import.meta.env.VITE_RUNTIME === "wails") {
    return new WailsDatabaseRepository()
  }

  return new ApiDatabaseRepository()
}

export interface DatabaseRepository {
  getAll(): Promise<any>
}