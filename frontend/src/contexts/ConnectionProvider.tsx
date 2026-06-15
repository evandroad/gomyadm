import { createContext, useContext, useEffect, useState, type ReactNode } from "react"
import type { Connection } from "../models"
import { API_URL } from "@/api"

type ConnectionContextType = {
  activeConnection: Connection | null
  loading: boolean
  setActiveConnection: (connection: Connection | null) => void
}

const ConnectionContext = createContext<ConnectionContextType | undefined>(undefined)

export function ConnectionProvider({ children }: { children: ReactNode }) {
  const [activeConnection, setActiveConnection] = useState<Connection | null>(null)
  const [loading, setLoading] = useState(true)

  async function checkActiveConnection() {
    setLoading(true)
    try {
      const res = await fetch(`${API_URL}/api/session`)

      if (!res.ok) {
        setActiveConnection(null)
        return
      }

      const data = await res.json()
      if (data) setActiveConnection(data)
    } catch (err) {
      console.error(err)
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
      checkActiveConnection()
    }, [])

  return (
    <ConnectionContext.Provider value={{ activeConnection, loading, setActiveConnection }}>
      {children}
    </ConnectionContext.Provider>
  )
}

export function useConnection() {
  const context = useContext(ConnectionContext)

  if (!context) {
    throw new Error("useConnection deve ser usado dentro do ConnectionProvider")
  }

  return context
}