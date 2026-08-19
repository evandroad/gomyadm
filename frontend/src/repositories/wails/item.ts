import { GetAll , Create, Update, Delete } from '../../../wailsjs/go/bindings/ItemBinding'
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

  async create(item: any): Promise<Result<void>> {
    try {
      await Create(item)
      return {
        ok: true, data: undefined
      }
    } catch (error) {
      return {
        ok: false, error: error instanceof Error ? error.message : "Erro ao conectar"
      }
    }
  }

  async update(item: any): Promise<Result<void>> {
    try {
      await Update(item)
      return {
        ok: true, data: undefined
      }
    } catch (error) {
      return {
        ok: false, error: error instanceof Error ? error.message : "Erro ao conectar"
      }
    }
  }

  async delete(item: any): Promise<Result<void>> {
    try {
      await Delete(item)
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