import { API_URL } from "@/api";
import type { DatabaseRepository } from "../database";

export class ApiConnectionRepository implements DatabaseRepository {
  async getAll(): Promise<any> {
    const response = await fetch(`${API_URL}/api/database`)

    if (!response.ok) {
      throw new Error("Erro ao buscar versão")
    }

    const data = await response.json()

    return data
  }
}