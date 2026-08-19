import { API_URL } from "@/api";
import type { Result } from "..";
import type { TableRepository } from "../table";

export class ApiTableRepository implements TableRepository {
  async getAll(): Promise<Result<string[]>> {
    const response = await fetch(`${API_URL}/api/table`)

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