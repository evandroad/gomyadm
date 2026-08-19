import { API_URL } from "@/api";
import type { ConnectionRepository } from "../connection";
import type { Connection } from "@/models";
import type { Result } from "..";

export class ApiConnectionRepository implements ConnectionRepository {
  async getAll(): Promise<Result<Connection[]>> {
    const response = await fetch(`${API_URL}/api/connection`)

    if (!response.ok) {
      const error = await response.json().catch(() => null)
      return {
        ok: false, error: error?.message ?? "Erro ao buscar conexões"
      }
    }

    return {
      ok: true, data: await response.json()
    }
  }

  async create(con: any): Promise<Result<void>> {
    const response = await fetch(`${API_URL}/api/connection`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(con),
    })

    if (!response.ok) {
      const error = await response.json().catch(() => null)
      return {
        ok: false, error: error?.message ?? "Erro ao criar conexão"
      }
    }

    return {
      ok: true, data: await response.json()
    }
  }
}