import type { Connection } from "@/models"
import { createContext, useContext } from "react"

type ConnectionsContextType = {
  connections: Connection[]
  loading: boolean
  insertConnection: (connection: Connection) => void
  updateConnection: (id: string, connection: Partial<Connection>) => void
  removeConnection: (id: string) => void
  getConnectionById: (id: string) => Connection | undefined
  clearConnections: () => void
}

export const ConnectionsContext = createContext<ConnectionsContextType | undefined>(undefined)

export function useConnections() {
  const context = useContext(ConnectionsContext)

  if (!context) {
    throw new Error("useConnections deve ser usado dentro do ConnectionsProvider")
  }

  return context
}