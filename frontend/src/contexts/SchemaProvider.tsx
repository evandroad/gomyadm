import { API_URL } from "@/api"
import type { TableSchema } from "@/models"
import { createContext, useContext, useEffect, useState, type ReactNode } from "react"
import { useDatabase } from "./DatabaseProvider"

type SchemaContextType = {
  activeSchema: TableSchema | null
  setActiveSchema: (schema: TableSchema | null) => void
  load: (table: string) => void
}

const SchemaContext = createContext<SchemaContextType | undefined>(undefined)

export function SchemaProvider({ children }: { children: ReactNode }) {
  const { activeDatabase } = useDatabase()
  const [activeSchema, setActiveSchema] = useState<TableSchema | null>(null)
  const [thisTable, setThisTable] = useState<string | null>(null)

  async function load(table: string) {
    if (!activeDatabase) return
    if (table === thisTable) return
    setThisTable(table)
    const res = await fetch(`${API_URL}/api/tables/struct/${table}`)
    if (!res.ok) {
      setActiveSchema(null)
      return
    }
    const data = await res.json()
    setActiveSchema(data)
  }

  useEffect(() => {
    if (!activeDatabase) setActiveSchema(null)
  }, [activeDatabase])

  return (
    <SchemaContext.Provider value={{ activeSchema, setActiveSchema, load }}>
      {children}
    </SchemaContext.Provider>
  )
}

export function useSchema() {
  const context = useContext(SchemaContext)

  if (!context) {
    throw new Error("useDatabase deve ser usado dentro do DatabaseProvider")
  }

  return context
}