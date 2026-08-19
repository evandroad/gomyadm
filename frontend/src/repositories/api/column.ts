import { API_URL } from "@/api";
import type { Result } from "..";
import type { ColumnRepository } from "../column";

export class ApiColumnRepository implements ColumnRepository {
  async getAll(table: string): Promise<Result<any>> {
    const response = await fetch(`${API_URL}/api/table/column/${table}`)

    if (!response.ok) {
      const error = await response.json().catch(() => null)
      return {
        ok: false, error: error?.error ?? "Erro ao buscar colunas"
      }
    }

    return {
      ok: true, data: await response.json()
    }
  }

  async create(item: any): Promise<Result<void>> {
    const response = await fetch(`${API_URL}/api/column/item`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(item)
    })

    if (!response.ok) {
      const error = await response.json().catch(() => null)
      return {
        ok: false, error: error?.error ?? "Erro ao criar coluna"
      }
    }

    return {
      ok: true, data: undefined
    }
  }

  async update(item: any): Promise<Result<void>> {
    const response = await fetch(`${API_URL}/api/table/column`, {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(item)
    })

    if (!response.ok) {
      const error = await response.json().catch(() => null)
      return {
        ok: false, error: error?.error ?? "Erro ao editar coluna"
      }
    }

    return {
      ok: true, data: undefined
    }
  }

  async delete(table: string, column: string): Promise<Result<void>> {
    const response = await fetch(`${API_URL}/api/table/column/${table}/${column}`, {
      method: "DELETE",
      headers: { "Content-Type": "application/json" }
    })

    if (!response.ok) {
      const error = await response.json().catch(() => null)
      return {
        ok: false, error: error?.error ?? "Erro ao apagar coluna"
      }
    }

    return {
      ok: true, data: undefined
    }
  }
}