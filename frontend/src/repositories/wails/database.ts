import type { Result } from '..';
import { GetAll } from '../../../wailsjs/go/bindings/DatabaseBinding'
import type { DatabaseRepository } from '../database';

export class WailsDatabaseRepository implements DatabaseRepository {
  async getAll(): Promise<Result<any>> {
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
}