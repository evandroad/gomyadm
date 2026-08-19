import { API_URL } from "@/api";
import type { DatabaseRepository } from "../database";
import type { Result } from "..";

export class ApiDatabaseRepository implements DatabaseRepository {
  async getAll(): Promise<Result<any>> {
    const response = await fetch(`${API_URL}/api/database`)

    if (!response.ok) {
      return {
        ok: false, error: "Erro ao buscar databases"
      }
    }


    const data = await response.json()
    return {
      ok: true, data
    }
  }
}