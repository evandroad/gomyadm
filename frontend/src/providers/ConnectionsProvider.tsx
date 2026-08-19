import { ConnectionsContext } from "@/contexts/ConnectionsContext";
import type { Connection } from "@/models";
import { repositories } from "@/repositories";
import { useEffect, useState, type ReactNode } from "react";

export function ConnectionsProvider({ children }: { children: ReactNode }) {
  const [connections, setConnections] = useState<Connection[]>([])
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    async function loadConnections() {
      try {
        const res = await repositories.connection.getAll()
        if (res.ok) setConnections(res.data)
      } catch (err) {
        console.error(err)
      } finally {
        setLoading(false)
      }
    }

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