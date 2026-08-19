import { GetAll, Create } from '../../../wailsjs/go/bindings/TableBinding'
import type { Result } from "..";
import type { TableRepository } from "../table";

export class WailsTableRepository implements TableRepository {
  async getAll(): Promise<Result<string[]>> {
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

  async create(table: any): Promise<Result<void>> {
    try {
      await Create(table)
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