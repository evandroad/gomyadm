import { API_URL } from "@/api";
import type { DatabaseRepository } from "../database";
import type { Result } from "..";

export class ApiDatabaseRepository implements DatabaseRepository {
  async getAll(): Promise<Result<any>> {
    const response = await fetch(`${API_URL}/api/database`)

    if (!response.ok) {
      const error = await response.json().catch(() => null)
      return {
        ok: false, error: error?.message ?? "Erro ao buscar databases"
      }
    }

    return {
      ok: true, data: await response.json()
    }
  }

  async select(database: string): Promise<Result<void>> {
    const response = await fetch(`${API_URL}/api/database/select`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ name: database }),
    })

    if (!response.ok) {
      const error = await response.json().catch(() => null)
      return {
        ok: false, error: error?.message ?? "Erro ao selecionar database"
      }
    }

    return {
      ok: true, data: undefined
    }
  }

  async create(database: any): Promise<Result<void>> {
    const response = await fetch(`${API_URL}/api/database`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(database),
    })

    if (!response.ok) {
      const error = await response.json().catch(() => null)
      return {
        ok: false, error: error?.message ?? "Erro ao criar database"
      }
    }

    return {
      ok: true, data: undefined
    }
  }

  async update(database: any): Promise<Result<void>> {
    const response = await fetch(`${API_URL}/api/database`, {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(database),
    })

    if (!response.ok) {
      const error = await response.json().catch(() => null)
      return {
        ok: false, error: error?.message ?? "Erro ao editar database"
      }
    }

    return {
      ok: true, data: undefined
    }
  }

  async delete(database: string): Promise<Result<void>> {
    const response = await fetch(`${API_URL}/api/database/${database}`, {
      method: "DELETE",
      headers: { "Content-Type": "application/json" },
    })

    if (!response.ok) {
      const error = await response.json().catch(() => null)
      return {
        ok: false, error: error?.message ?? "Erro ao editar database"
      }
    }

    return {
      ok: true, data: undefined
    }
  }
}