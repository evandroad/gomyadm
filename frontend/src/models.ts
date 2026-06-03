export type Connection = {
  id: string
  name: string
  driver: string
  host: string
  port: number | string
  username: string
  password: string
  database: string
  databases: string[]
}

export function createConnection(data?: Partial<Connection>): Connection {
  return {
    id: "",
    name: "",
    driver: "",
    host: "",
    port: "",
    username: "",
    password: "",
    database: "",
    databases: [],
    ...data
  }
}