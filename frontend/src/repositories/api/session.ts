import { API_URL } from "@/api";
import type { SessionRepository } from "../session";

export class ApiSessionRepository implements SessionRepository {
  async active(): Promise<string> {
    const response = await fetch(`${API_URL}/api/session`)

    if (!response.ok) {
      console.error("Erro ao buscar session")
      return ''
    }

    const data = await response.json()

    return data
  }
}