import { createConnectionRepository } from "./connection";
import { createAppRepository } from "./app";

export const repositories = {
  app: createAppRepository(),
  connection: createConnectionRepository(),
  // session: createSessionRepository(),
  // database: createDatabaseRepository(),
  // table: createTableRepository(),
  // column: createColumnRepository(),
  // item: createItemRepository(),
  // query: createQueryRepository(),
}