import type { AppRepository } from "./app"
import { ApiAppRepository } from "./api/app"
import { WailsAppRepository } from "./wails/app"

function createAppRepository(): AppRepository {
  if (import.meta.env.VITE_RUNTIME === "wails") {
    return new WailsAppRepository()
  }

  return new ApiAppRepository()
}

export const appRepository = createAppRepository()

export const repositories = {
  app: createAppRepository(),
  // connection: createConnectionRepository(),
  // session: createSessionRepository(),
  // database: createDatabaseRepository(),
  // table: createTableRepository(),
  // column: createColumnRepository(),
  // item: createItemRepository(),
  // query: createQueryRepository(),
}