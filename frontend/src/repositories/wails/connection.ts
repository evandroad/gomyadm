import type { Connection } from '@/models';
import { GetAll, Create, Update, Delete } from '../../../wailsjs/go/bindings/ConnectionBinding'
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
      await Create(con)
      return {
        ok: true, data: undefined
      }
    } catch (error) {
      return {
        ok: false, error: error instanceof Error ? error.message : "Erro ao conectar"
      }
    }
  }

  async update(id: string, con: any): Promise<Result<void>> {
    try {
      await Update(id, con)
      return {
        ok: true, data: undefined
      }
    } catch (error) {
      return {
        ok: false, error: error instanceof Error ? error.message : "Erro ao conectar"
      }
    }
  }

  async delete(id: string): Promise<Result<void>> {
    try {
      await Delete(id)
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