import { GetAll /*, Create, Update, Delete*/ } from '../../../wailsjs/go/bindings/ItemBinding'
import type { Result } from "..";
import type { ItemRepository } from '../item';

export class WailsItemRepository implements ItemRepository {
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

  async create(/*table: any*/): Promise<Result<void>> {
    try {
      // await Create(table)
      return {
        ok: true, data: undefined
      }
    } catch (error) {
      return {
        ok: false, error: error instanceof Error ? error.message : "Erro ao conectar"
      }
    }
  }

  async update(/*oldName: any*/): Promise<Result<void>> {
    try {
      // await Update(oldName, newName)
      return {
        ok: true, data: undefined
      }
    } catch (error) {
      return {
        ok: false, error: error instanceof Error ? error.message : "Erro ao conectar"
      }
    }
  }

  async delete(/*table: string*/): Promise<Result<void>> {
    try {
      // await Delete(table)
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