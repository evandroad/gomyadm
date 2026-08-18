import { useConnection } from "@/contexts/ConectionContext";
import { Navigate, Outlet } from "react-router-dom"

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