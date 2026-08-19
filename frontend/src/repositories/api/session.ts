import { API_URL } from "@/api";
import type { SessionRepository } from "../session";

export class ApiSessionRepository implements SessionRepository {
  async active(): Promise<string> {
    const response = await fetch(`${API_URL}/api/session`)

    if (!response.ok) {
      console.error("Erro ao buscar session")
      return ''
    }

    return await response.json()
  }

  async connect(conn: any): Promise<any> {
    const response = await fetch( API_URL + "/api/session", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(conn),
    })

    if (!response.ok) {
      console.error("Erro ao conectar session")
      return null
    }

    return await response.json()
  }

  async disconnect(): Promise<any> {
    const response = await fetch( API_URL + "/api/session", {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
      }
    })

    if (!response.ok) {
      console.error("Erro ao desconectar session")
      return null
    }

    return await response.json()
  }
}