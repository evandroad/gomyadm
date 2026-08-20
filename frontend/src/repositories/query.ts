import type { Result } from ".";
import { ApiQueryRepository } from "./api/query";
import { WailsQueryRepository } from "./wails/query";

export interface QueryRepository {
  query(query: string): Promise<Result<any>>
}

export function createQueryRepository(): QueryRepository {
  if (import.meta.env.VITE_RUNTIME === "wails") {
    return new WailsQueryRepository()
  }

  return new ApiQueryRepository()
}