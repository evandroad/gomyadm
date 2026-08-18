import type { AppRepository } from "../app"
import { Version } from '../../../wailsjs/go/bindings/AppBinding'

export class WailsAppRepository implements AppRepository {
  async version(): Promise<string> {
    return Version()
  }
}