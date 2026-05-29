import { API_URL } from "@/api"
import type { Connection } from "@/models"
import { createContext, useContext, useEffect, useState, type ReactNode } from "react"

type ConnectionsContextType = {
  connections: Connection[]
  loading: boolean
  insertConnection: (connection: Connection) => void
  updateConnection: (id: string, connection: Partial<Connection>) => void
  removeConnection: (id: string) => void
  getConnectionById: (id: string) => Connection | undefined
  clearConnections: () => void
}

const ConnectionsContext = createContext<ConnectionsContextType | undefined>(undefined)

export function ConnectionsProvider({ children }: { children: ReactNode }) {
  const [connections, setConnections] = useState<Connection[]>([])
  const [loading, setLoading] = useState(true)

  async function loadConnections() {
    try {
      const res = await fetch(`${API_URL}/api/connection`)

      if (!res.ok) {
        setConnections([])
        return
      }

      const data = await res.json()
      if (data) setConnections(data)
    } catch (err) {
      console.error(err)
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    loadConnections()
  }, [])

  function insertConnection(connection: Connection) {
    setConnections((prev) => {
      const exists = prev.some(c => c.id === connection.id)
      if (exists) return prev
      return [...prev, connection]
    })
  }

  function updateConnection(id: string, updated: Partial<Connection>) {
    setConnections(prev => prev.map((conn) => conn.id === id ? { ...conn, ...updated } : conn))
  }

  function removeConnection(id: string) {
    setConnections(prev => prev.filter((conn) => conn.id !== id))
  }

  function getConnectionById(id: string) {
    return connections.find(conn => conn.id === id)
  }

  function clearConnections() {
    setConnections([])
  }

  return (
    <ConnectionsContext.Provider value={{ connections, loading, insertConnection, updateConnection, removeConnection, getConnectionById, clearConnections }}>
      {children}
    </ConnectionsContext.Provider>
  )
}

export function useConnections() {
  const context = useContext(ConnectionsContext)

  if (!context) {
    throw new Error("useConnections deve ser usado dentro do ConnectionsProvider")
  }

  return context
}