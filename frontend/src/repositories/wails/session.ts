import type { Result } from '..';
import { Active, Connect, Disconnect } from '../../../wailsjs/go/bindings/SessionBinding'
import type { SessionRepository } from "../session";

export class WailsSessionRepository implements SessionRepository {
  async active(): Promise<Result<any>> {
    try {
      return {
        ok: true, data: await Active()
      }
    } catch (error) {
      return {
        ok: false, error: error instanceof Error ? error.message : "Erro ao obter conexão ativa"
      }
    }
  }

  async connect(conn: any): Promise<Result<any>> {
    try {
      return {
        ok: true, data: await Connect(conn)
      }
    } catch (error) {
      return {
        ok: false, error: error instanceof Error ? error.message : "Erro ao conectar"
      }
    }
  }

  async disconnect(): Promise<Result<void>> {
    try {
      await Disconnect()
      return {
        ok: true, data: undefined
      }
    } catch (error) {
      return {
        ok: false, error: error instanceof Error ? error.message : "Erro ao desconectar"
      }
    }
  }
}