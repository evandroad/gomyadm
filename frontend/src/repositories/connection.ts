import type { Connection } from "@/models";
import { WailsConnectionRepository } from "./wails/connection";
import { ApiConnectionRepository } from "./api/connection";
import type { Result } from ".";

export function createConnectionRepository(): ConnectionRepository {
  if (import.meta.env.VITE_RUNTIME === "wails") {
    return new WailsConnectionRepository()
  }

  return new ApiConnectionRepository()
}

export interface ConnectionRepository {
  getAll(): Promise<Result<Connection[]>>
  create(con: any): Promise<Result<void>>
  update(con: any): Promise<Result<void>>
  delete(id: string): Promise<Result<void>>
}