import { API_URL } from "@/api";
import type { DatabaseRepository } from "../database";

export class ApiDatabaseRepository implements DatabaseRepository {
  async getAll(): Promise<any> {
    const response = await fetch(`${API_URL}/api/database`)

    if (!response.ok) {
      console.error("Erro ao buscar databases")
      return
    }

    const data = await response.json()

    return data
  }
}