import { API_URL } from "@/api"
import { DatabaseContext } from "@/contexts/DatabaseContext";
import { repositories } from "@/repositories";
import { useCallback, useEffect, useState, type ReactNode } from "react"

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

  const getDatabases = useCallback(async () => {
    const data = await repositories.database.getAll()

    if (data) {
      setActiveDatabase(data.active)
      setDatabases(data.databases)
    }
  }, [])

  useEffect(() => {
    // eslint-disable-next-line react-hooks/set-state-in-effect
    getDatabases()
  }, [getDatabases])

  return (
    <DatabaseContext.Provider value={{ activeDatabase, databases, setActiveDatabase, changeDatabase, getDatabases }}>
      {children}
    </DatabaseContext.Provider>
  )
}