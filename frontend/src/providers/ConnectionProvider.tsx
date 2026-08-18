import { useEffect, useState, type ReactNode } from "react"
import type { Connection } from "../models"
import { ConnectionContext } from "@/contexts/ConectionContext";
import { repositories } from "@/repositories";

export function ConnectionProvider({ children }: { children: ReactNode }) {
  const [activeConnection, setActiveConnection] = useState<Connection | null>(null)
  const [loading, setLoading] = useState(true)
  
  useEffect(() => {
    async function checkActiveConnection() {
      setLoading(true)
      try {
        const data = await repositories.session.active()
        if (data) setActiveConnection(data)
      } catch (err) {
        setActiveConnection(null)
        console.error(err)
      } finally {
        setLoading(false)
      }
    }

    checkActiveConnection()
  }, [])

  return (
    <ConnectionContext.Provider value={{ activeConnection, loading, setActiveConnection }}>
      {children}
    </ConnectionContext.Provider>
  )
}