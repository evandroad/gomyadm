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

  async create(item: any): Promise<Result<void>> {
    const response = await fetch(`${API_URL}/api/table/item`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(item)
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

  async update(item: any): Promise<Result<void>> {
    const response = await fetch(`${API_URL}/api/table/item`, {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(item)
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

  async delete(item: any): Promise<Result<void>> {
    const response = await fetch(`${API_URL}/api/table/item`, {
      method: "DELETE",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(item)
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