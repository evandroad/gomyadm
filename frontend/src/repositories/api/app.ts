import { API_URL } from "@/api";
import type { AppRepository } from "../app"

export class ApiAppRepository implements AppRepository {
  async version(): Promise<string> {
    const response = await fetch(`${API_URL}/api/version`)

    if (!response.ok) {
      throw new Error("Erro ao buscar versão")
    }

    const data = await response.json()

    return data.version
  }
}