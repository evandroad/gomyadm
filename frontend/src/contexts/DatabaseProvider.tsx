import { API_URL } from "@/api"
import { createContext, useContext, useEffect, useState, type ReactNode } from "react"

type DatabaseContextType = {
  activeDatabase: string | null
  databases: string[]
  setActiveDatabase: (database: string | null) => void
  changeDatabase: (database: string) => Promise<void>
  getDatabases: () => Promise<void>
}

const DatabaseContext = createContext<DatabaseContextType | undefined>(undefined)

export function DatabaseProvider({ children }: { children: ReactNode }) {
  const [activeDatabase, setActiveDatabase] = useState<string | null>(null)
  const [databases, setDatabases] = useState<string[]>([])

  async function changeDatabase(database: string) {
    const res = await fetch(`${API_URL}/api/database/select`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ name: database }),
    })

    if (!res.ok) return

    setActiveDatabase(database)
  }

  async function getDatabases() {
    const res = await fetch(`${API_URL}/api/database`)

    if (!res.ok) {
      return
    }

    const data = await res.json()
    if (data) {
      setActiveDatabase(data.active)
      setDatabases(data.databases)
    }
  }

  useEffect(() => {
    getDatabases()
  }, [])

  return (
    <DatabaseContext.Provider value={{ activeDatabase, databases, setActiveDatabase, changeDatabase, getDatabases }}>
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