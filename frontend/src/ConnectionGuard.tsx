import { Navigate, Outlet } from "react-router-dom"
import { useConnection } from "./contexts/ConnectionProvider"

export default function ConnectionGuard() {
  const { activeConnection, loading } = useConnection()

  if (loading) {
    return (
      <div className="min-h-screen flex items-center justify-center text-white">
        Conectando...
      </div>
    )
  }

  if (!activeConnection) {
    return <Navigate to="/connect" replace />
  }

  return <Outlet />
}