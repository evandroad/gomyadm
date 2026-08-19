import { API_URL } from "@/api";
import type { AppRepository } from "../app"
import type { Result } from "..";

export class ApiAppRepository implements AppRepository {
  async version(): Promise<Result<string>> {
    const response = await fetch(`${API_URL}/api/version`)

    if (!response.ok) {
      return {
        ok: false, error: "Erro ao buscar versão"
      }
    }

    const data = await response.json()
    return {
      ok: true, data: data.version,
    }
  }
}