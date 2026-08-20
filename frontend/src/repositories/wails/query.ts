import { Query } from '../../../wailsjs/go/bindings/QueryBinding'
import type { Result } from "..";
import type { QueryRepository } from "../query";

export class WailsQueryRepository implements QueryRepository {
  async query(query: string): Promise<Result<any>> {
    try {
      return {
        ok: true, data: await Query(query)
      }
    } catch (error) {
      return {
        ok: false, error: error instanceof Error ? error.message : "Erro ao conectar"
      }
    }
  }
}