import { GetAll , Create, Update, Delete } from '../../../wailsjs/go/bindings/ColumnBinding'
import type { Result } from "..";
import type { ColumnRepository } from '../column';

export class WailsColumnRepository implements ColumnRepository {
  async getAll(table: string): Promise<Result<any>> {
    try {
      return {
        ok: true, data: await GetAll(table)
      }
    } catch (error) {
      return {
        ok: false, error: error instanceof Error ? error.message : "Erro ao conectar"
      }
    }
  }

  async create(column: any): Promise<Result<void>> {
    try {
      await Create(column)
      return {
        ok: true, data: undefined
      }
    } catch (error) {
      return {
        ok: false, error: error instanceof Error ? error.message : "Erro ao conectar"
      }
    }
  }

  async update(column: any): Promise<Result<void>> {
    try {
      await Update(column)
      return {
        ok: true, data: undefined
      }
    } catch (error) {
      return {
        ok: false, error: error instanceof Error ? error.message : "Erro ao conectar"
      }
    }
  }

  async delete(table: string, column: string): Promise<Result<void>> {
    try {
      await Delete(table, column)
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