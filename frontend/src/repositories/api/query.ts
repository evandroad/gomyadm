import { API_URL } from "@/api";
import type { Result } from "..";
import type { QueryRepository } from "../query";

export class ApiQueryRepository implements QueryRepository {
  async query(query: string): Promise<Result<any>> {
    const response = await fetch(`${API_URL}/api/query`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ query }),
    })

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
}