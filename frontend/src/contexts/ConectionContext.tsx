import type { Connection } from "@/models";
import { createContext, useContext } from "react";

type ConnectionContextType = {
  activeConnection: Connection | null
  loading: boolean
  setActiveConnection: (connection: Connection | null) => void
}

export const ConnectionContext = createContext<ConnectionContextType | undefined>(undefined)

export function useConnection() {
  const context = useContext(ConnectionContext)

  if (!context) {
    throw new Error("useConnection deve ser usado dentro do ConnectionProvider")
  }

  return context
}