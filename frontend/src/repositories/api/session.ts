import { API_URL } from "@/api";
import type { SessionRepository } from "../session";
import type { Result } from "..";

export class ApiSessionRepository implements SessionRepository {
  async active(): Promise<Result<string>> {
    const response = await fetch(`${API_URL}/api/session`)

    if (!response.ok) {
      const error = await response.json().catch(() => null)
      return {
        ok: false, error: error?.error ?? "Erro ao buscar session"
      }
    }

    return {
      ok: true, data: await response.json()
    } 
  }

  async connect(conn: any): Promise<Result<any>> {
    const response = await fetch( API_URL + "/api/session", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(conn),
    })

    if (!response.ok) {
      const error = await response.json().catch(() => null)
      return {
        ok: false, error: error?.error ?? "Erro ao conectar session"
      }
    }

    return {
      ok: true, data: await response.json()
    }
  }

  async disconnect(): Promise<Result<void>> {
    const response = await fetch( API_URL + "/api/session", {
      method: "DELETE",
      headers: { "Content-Type": "application/json" }
    })

    if (!response.ok) {
      const error = await response.json().catch(() => null)
      return {
        ok: false, error: error?.error ?? "Erro ao desconectar session"
      }
    }

    return {
      ok: true, data: undefined
    }
  }
}