import { API_URL } from "@/api";
import type { ConnectionRepository } from "../connection";
import type { Connection } from "@/models";

export class ApiConnectionRepository implements ConnectionRepository {
  async getAll(): Promise<Connection[]> {
    const response = await fetch(`${API_URL}/api/connection`)

    if (!response.ok) {
      throw new Error("Erro ao buscar versão")
    }

    const data = await response.json()

    return data
  }
}