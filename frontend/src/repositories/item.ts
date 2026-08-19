import type { Result } from ".";
import { ApiItemRepository } from "./api/item";
import { WailsItemRepository } from "./wails/item";

export interface ItemRepository {
  getAll(table: string): Promise<Result<any>>
  create(con: any): Promise<Result<void>>
  update(con: any): Promise<Result<void>>
  delete(id: string): Promise<Result<void>>
}

export function createItemRepository(): ItemRepository {
  if (import.meta.env.VITE_RUNTIME === "wails") {
    return new WailsItemRepository()
  }

  return new ApiItemRepository()
}