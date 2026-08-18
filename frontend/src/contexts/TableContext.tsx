import type { Table } from "@/models";
import { createContext, useContext } from "react";

type TableContextType = {
  activeTable: Table | null
  setActiveTable: (table: Table | null) => void
  load: (table: string) => void
}

export const TableContext = createContext<TableContextType | undefined>(undefined)

export function useTable() {
  const context = useContext(TableContext)

  if (!context) {
    throw new Error("useTable deve ser usado dentro do TableProvider")
  }

  return context
}