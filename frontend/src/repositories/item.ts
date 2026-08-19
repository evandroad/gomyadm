import type { Result } from ".";
import { ApiItemRepository } from "./api/item";
import { WailsItemRepository } from "./wails/item";

export interface ItemRepository {
  getAll(table: string): Promise<Result<any>>
  create(item: any): Promise<Result<void>>
  update(item: any): Promise<Result<void>>
  delete(item: any): Promise<Result<void>>
}

export function createItemRepository(): ItemRepository {
  if (import.meta.env.VITE_RUNTIME === "wails") {
    return new WailsItemRepository()
  }

  return new ApiItemRepository()
}