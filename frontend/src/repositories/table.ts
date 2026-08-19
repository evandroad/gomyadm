import type { Result } from ".";
import { ApiTableRepository } from "./api/table";
import { WailsTableRepository } from "./wails/table";

export interface TableRepository {
  getAll(): Promise<Result<string[]>>
  create(table: any): Promise<Result<void>>
  update(oldName: string, newName: string): Promise<Result<void>>
}

export function createTableRepository(): TableRepository {
  if (import.meta.env.VITE_RUNTIME === "wails") {
    return new WailsTableRepository()
  }

  return new ApiTableRepository()
}
