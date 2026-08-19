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
}