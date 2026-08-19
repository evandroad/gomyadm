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
      
      const res = await repositories.session.active()
      
      if (res.ok) {
        setActiveConnection(res.data)
      } else {
        setActiveConnection(null)
        console.error(res.error)
      }

      setLoading(false)
    }

    checkActiveConnection()
  }, [])

  return (
    <ConnectionContext.Provider value={{ activeConnection, loading, setActiveConnection }}>
      {children}
    </ConnectionContext.Provider>
  )
}