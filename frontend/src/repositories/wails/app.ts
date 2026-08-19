import type { AppRepository } from "../app"
import { Version } from '../../../wailsjs/go/bindings/AppBinding'
import type { Result } from "..";

export class WailsAppRepository implements AppRepository {
  async version(): Promise<Result<string>> {
    try {
      return {
        ok: true, data: await Version()
      }
    } catch (error) {
      return {
        ok: false, error: error instanceof Error ? error.message : "Erro ao conectar"
      }
    }
  }
}