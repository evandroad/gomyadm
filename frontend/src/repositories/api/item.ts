import { API_URL } from "@/api";
import type { Result } from "..";
import type { ItemRepository } from "../item";

export class ApiItemRepository implements ItemRepository {
  async getAll(table: string): Promise<Result<any>> {
    const response = await fetch(`${API_URL}/api/table/item/${table}`)

    if (!response.ok) {
      const error = await response.json().catch(() => null)
      return {
        ok: false, error: error?.error ?? "Erro ao buscar versão"
      }
    }

    return {
      ok: true, data: await response.json()
    }
  }

  async create(table: any): Promise<Result<void>> {
    const response = await fetch(`${API_URL}/api/table`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(table),
    })

    if (!response.ok) {
      const error = await response.json().catch(() => null)
      return {
        ok: false, error: error?.error ?? "Erro ao buscar versão"
      }
    }

    return {
      ok: true, data: undefined
    }
  }

  async update(oldName: any): Promise<Result<void>> {
    const response = await fetch(`${API_URL}/api/table/${oldName}`, {
      method: "PUT",
      headers: { "Content-Type": "application/json" }
    })

    if (!response.ok) {
      const error = await response.json().catch(() => null)
      return {
        ok: false, error: error?.error ?? "Erro ao buscar versão"
      }
    }

    return {
      ok: true, data: undefined
    }
  }

  async delete(table: string): Promise<Result<void>> {
    const response = await fetch(`${API_URL}/api/table/${table}`, {
      method: "DELETE",
      headers: { "Content-Type": "application/json" }
    })

    if (!response.ok) {
      const error = await response.json().catch(() => null)
      return {
        ok: false, error: error?.error ?? "Erro ao buscar versão"
      }
    }

    return {
      ok: true, data: undefined
    }
  }
}