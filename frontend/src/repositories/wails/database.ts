import { GetAll } from '../../../wailsjs/go/bindings/DatabaseBinding'
import type { DatabaseRepository } from '../database';

export class WailsDatabaseRepository implements DatabaseRepository {
  async getAll(): Promise<any> {
    return GetAll()
  }
}