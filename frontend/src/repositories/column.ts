import type { Result } from ".";
import { ApiColumnRepository } from "./api/column";
import { WailsColumnRepository } from "./wails/column";

export interface ColumnRepository {
  getAll(table: string): Promise<Result<any>>
  create(item: any): Promise<Result<void>>
  update(item: any): Promise<Result<void>>
  delete(table: string, column: string): Promise<Result<void>>
}

export function createColumnRepository(): ColumnRepository {
  if (import.meta.env.VITE_RUNTIME === "wails") {
    return new WailsColumnRepository()
  }

  return new ApiColumnRepository()
}