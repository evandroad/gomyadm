import type { Table, TableData } from "@/models"
import { useCallback, useEffect, useState, type ReactNode } from "react"
import { TableContext } from "@/contexts/TableContext";
import { useDatabase } from "@/contexts/DatabaseContext";
import { repositories } from "@/repositories";

export function TableProvider({ children }: { children: ReactNode }) {
  const { activeDatabase } = useDatabase()
  const [activeTable, setActiveTable] = useState<Table | null>(null)
  const [tableData, setTableData] = useState<TableData | null>(null)

  const load = useCallback(async (table: string | null) => {
    if (!activeDatabase) return
    if (table == '' || table == null) return

    const res = await repositories.column.getAll(table)
    if (!res.ok) {
      setActiveTable(null)
      return
    }
    
    setActiveTable(res.data)
  }, [activeDatabase])

  const loadTableData = useCallback(async (table: string | null) => {
    if (!activeDatabase) return
    if (table == '' || table == null) return

    const res = await repositories.item.getAll(table)
    
    if (!res.ok) {
      setTableData(null)
      return
    }

    setTableData(res.data)
  }, [activeDatabase])

  useEffect(() => {
    // eslint-disable-next-line react-hooks/set-state-in-effect
    if (!activeDatabase) setActiveTable(null)
  }, [activeDatabase])

  return (
    <TableContext.Provider value={{ activeTable, setActiveTable, load, tableData, setTableData, loadTableData }}>
      {children}
    </TableContext.Provider>
  )
}