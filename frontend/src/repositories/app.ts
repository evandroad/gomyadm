export interface AppRepository {
  version(): Promise<string>
}