import { Active, Connect, Disconnect } from '../../../wailsjs/go/bindings/SessionBinding'
import type { SessionRepository } from "../session";

export class WailsSessionRepository implements SessionRepository {
  async active(): Promise<any> {
    return Active()
  }

  async connect(conn: any) {
    return Connect(conn)
  }

  async disconnect() {
    return Disconnect()
  }
}