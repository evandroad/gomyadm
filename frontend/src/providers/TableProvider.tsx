import { API_URL } from "@/api"
import type { Table } from "@/models"
import { useEffect, useState, type ReactNode } from "react"
import { TableContext } from "@/contexts/TableContext";
import { useDatabase } from "@/contexts/DatabaseContext";

export function TableProvider({ children }: { children: ReactNode }) {
  const { activeDatabase } = useDatabase()
  const [activeTable, setActiveTable] = useState<Table | null>(null)

  async function load(table: string) {
    if (!activeDatabase) return
    if (table === activeTable?.name) return
    if (table == '') return

    const res = await fetch(`${API_URL}/api/table/column/${table}`)
    if (!res.ok) {
      setActiveTable(null)
      return
    }
    const data = await res.json()
    setActiveTable(data)
  }

  useEffect(() => {
    // eslint-disable-next-line react-hooks/set-state-in-effect
    if (!activeDatabase) setActiveTable(null)
  }, [activeDatabase])

  return (
    <TableContext.Provider value={{ activeTable, setActiveTable, load }}>
      {children}
    </TableContext.Provider>
  )
}