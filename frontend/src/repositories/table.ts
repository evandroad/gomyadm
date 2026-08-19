import type { Result } from ".";
import { ApiTableRepository } from "./api/table";
import { WailsTableRepository } from "./wails/table";

export interface TableRepository {
  getAll(): Promise<Result<string[]>>
}

export function createTableRepository(): TableRepository {
  if (import.meta.env.VITE_RUNTIME === "wails") {
    return new WailsTableRepository()
  }

  return new ApiTableRepository()
}
