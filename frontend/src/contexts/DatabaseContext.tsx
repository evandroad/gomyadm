import { createContext, useContext } from "react";

type DatabaseContextType = {
  activeDatabase: string | null
  databases: string[]
  setActiveDatabase: (database: string | null) => void
  changeDatabase: (database: string) => Promise<void>
  getDatabases: () => Promise<void>
}

export const DatabaseContext = createContext<DatabaseContextType | undefined>(undefined)

export function useDatabase() {
  const context = useContext(DatabaseContext)

  if (!context) {
    throw new Error("useDatabase deve ser usado dentro do DatabaseProvider")
  }

  return context
}