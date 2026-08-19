import type { Result } from '..';
import { GetAll, Select, Create, Update } from '../../../wailsjs/go/bindings/DatabaseBinding'
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

  async select(database: string): Promise<Result<void>> {
    try {
      await Select(database)
      return {
        ok: true, data: undefined
      }
    } catch (error) {
      return {
        ok: false, error: error instanceof Error ? error.message : "Erro ao conectar"
      }
    }
  }

  async create(database: any): Promise<Result<void>> {
    try {
      await Create(database)
      return {
        ok: true, data: undefined
      }
    } catch (error) {
      return {
        ok: false, error: error instanceof Error ? error.message : "Erro ao conectar"
      }
    }
  }

  async update(database: any): Promise<Result<void>> {
    try {
      await Update(database)
      return {
        ok: true, data: undefined
      }
    } catch (error) {
      return {
        ok: false, error: error instanceof Error ? error.message : "Erro ao conectar"
      }
    }
  }
}