export type Values = Record<string, any>

export type Connection = {
  id:        string
  name:      string
  driver:    string
  host:      string
  port:      number | string
  username:  string
  password:  string
  database:  string
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

export type ColumnSchema = {
	name:       string
	type:       string
	nullable:   boolean
	key:        string
	default:    string
	extra:      string
	autoNumber: boolean
}

export type TableStructure = {
  name:    string
  columns: ColumnSchema[]
}