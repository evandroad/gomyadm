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

export type Column = {
  name: string
  type: string
  length?: string
  nullable: boolean
  primary: boolean
  unique: boolean
  autoIncrement: boolean
  defaultValue: string
}

export function createColumn(data?: Partial<Column>): Column {
  return {
    name: "",
    type: '',
    length: '',
    nullable: false,
    primary: false,
    unique: false,
    autoIncrement: false,
    defaultValue: '',
    ...data
  }
}

export const ColumnTypes = ["VARCHAR", "TEXT", "INT", "BIGINT", "DECIMAL", "BOOLEAN", "DATE", "DATETIME"]

export type TableSchema = {
  name:    string
  columns: Column[]
}

export type TableData = {
  columns: string[]
  rows: Values
}