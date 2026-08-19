import type { Connection } from '@/models';
import { GetAll, Create } from '../../../wailsjs/go/bindings/ConnectionBinding'
import type { ConnectionRepository } from "../connection";
import type { Result } from '..';

export class WailsConnectionRepository implements ConnectionRepository {
  async getAll(): Promise<Result<Connection[]>> {
    try {
      return {
        ok: true, data: await GetAll()
      }
    } catch (error) {
      return {
        ok: false, error: error instanceof Error ? error.message : "Erro ao conectar"
      }
    }
  }

  async create(con: any): Promise<Result<void>> {
    try {
      return {
        ok: true, data: await Create(con)
      }
    } catch (error) {
      return {
        ok: false, error: error instanceof Error ? error.message : "Erro ao conectar"
      }
    }
  }
}