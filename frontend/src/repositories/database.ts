import { WailsDatabaseRepository } from "./wails/database";
import { ApiDatabaseRepository } from "./api/database";
import type { Result } from ".";

export interface DatabaseRepository {
  getAll(): Promise<Result<any>>
  select(database: string): Promise<Result<void>>
  create(database: any): Promise<Result<void>>
  update(database: any): Promise<Result<void>>
  delete(database: string): Promise<Result<void>>
}

export function createDatabaseRepository(): DatabaseRepository {
  if (import.meta.env.VITE_RUNTIME === "wails") {
    return new WailsDatabaseRepository()
  }

  return new ApiDatabaseRepository()
}