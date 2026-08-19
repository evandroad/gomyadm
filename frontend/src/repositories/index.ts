import { createConnectionRepository } from "./connection";
import { createAppRepository } from "./app";
import { createDatabaseRepository } from "./database";
import { createSessionRepository } from "./session";
import { createTableRepository } from "./table";
import { createItemRepository } from "./item";

export const repositories = {
  app: createAppRepository(),
  connection: createConnectionRepository(),
  session: createSessionRepository(),
  database: createDatabaseRepository(),
  table: createTableRepository(),
  // column: createColumnRepository(),
  item: createItemRepository(),
  // query: createQueryRepository(),
}

export type Result<T> = | { ok: true; data: T } | { ok: false; error: string }