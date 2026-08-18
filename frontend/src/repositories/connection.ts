import type { Connection } from "@/models";
import { WailsConnectionRepository } from "./wails/connection";
import { ApiConnectionRepository } from "./api/connection";

export function createConnectionRepository(): ConnectionRepository {
  if (import.meta.env.VITE_RUNTIME === "wails") {
    return new WailsConnectionRepository()
  }

  return new ApiConnectionRepository()
}

export interface ConnectionRepository {
  getAll(): Promise<Connection[]>
}