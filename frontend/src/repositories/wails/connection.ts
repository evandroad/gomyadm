import type { Connection } from '@/models';
import { GetAll } from '../../../wailsjs/go/bindings/ConnectionBinding'
import type { ConnectionRepository } from "../connection";

export class WailsConnectionRepository implements ConnectionRepository {
  async getAll(): Promise<Connection[]> {
    return GetAll()
  }
}