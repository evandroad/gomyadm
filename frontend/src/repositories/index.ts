import { createConnectionRepository } from "./connection";
import { createAppRepository } from "./app";
import { createDatabaseRepository } from "./database";

export const repositories = {
  app: createAppRepository(),
  connection: createConnectionRepository(),
  // session: createSessionRepository(),
  database: createDatabaseRepository(),
  // table: createTableRepository(),
  // column: createColumnRepository(),
  // item: createItemRepository(),
  // query: createQueryRepository(),
}