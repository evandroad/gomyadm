import { API_URL } from "@/api"
import type { Table } from "@/models"
import { createContext, useContext, useEffect, useState, type ReactNode } from "react"
import { useDatabase } from "./DatabaseProvider"

type TableContextType = {
  activeTable: Table | null
  setActiveTable: (table: Table | null) => void
  load: (table: string) => void
}

const TableContext = createContext<TableContextType | undefined>(undefined)

export function TableProvider({ children }: { children: ReactNode }) {
  const { activeDatabase } = useDatabase()
  const [activeTable, setActiveTable] = useState<Table | null>(null)

  async function load(table: string) {
    if (!activeDatabase || !activeTable) return
    if (table === activeTable.name) return

    const res = await fetch(`${API_URL}/api/table/column/${table}`)
    if (!res.ok) {
      setActiveTable(null)
      return
    }
    const data = await res.json()
    setActiveTable(data)
  }

  useEffect(() => {
    if (!activeDatabase) setActiveTable(null)
  }, [activeDatabase])

  return (
    <TableContext.Provider value={{ activeTable, setActiveTable, load }}>
      {children}
    </TableContext.Provider>
  )
}

export function useTable() {
  const context = useContext(TableContext)

  if (!context) {
    throw new Error("useTable deve ser usado dentro do TableProvider")
  }

  return context
}