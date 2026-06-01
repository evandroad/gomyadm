import { API_URL } from "@/api"
import { createContext, useContext, useState, type ReactNode } from "react"

type DatabaseContextType = {
  activeDatabase: string | null
  setActiveDatabase: (database: string | null) => void
  changeDatabase: (database: string) => Promise<void>
}

const DatabaseContext = createContext<DatabaseContextType | undefined>(undefined)

export function DatabaseProvider({ children }: { children: ReactNode }) {
  const [activeDatabase, setActiveDatabase] = useState<string | null>(null)

  async function changeDatabase(database: string) {
    const res = await fetch(`${API_URL}/api/connection/database/select`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ database }),
    })

    if (!res.ok) return

    setActiveDatabase(database)
  }

  return (
    <DatabaseContext.Provider value={{ activeDatabase, setActiveDatabase, changeDatabase }}>
      {children}
    </DatabaseContext.Provider>
  )
}

export function useDatabase() {
  const context = useContext(DatabaseContext)

  if (!context) {
    throw new Error("useDatabase deve ser usado dentro do DatabaseProvider")
  }

  return context
}