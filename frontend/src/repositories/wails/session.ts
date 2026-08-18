import { Active } from '../../../wailsjs/go/bindings/SessionBinding'
import type { SessionRepository } from "../session";

export class WailsSessionRepository implements SessionRepository {
  async active(): Promise<any> {
    return Active()
  }
}