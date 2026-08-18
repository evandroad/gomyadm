import { useConnection } from "@/contexts/ConectionContext";
import { Navigate, Outlet } from "react-router-dom"

export default function GuestGuard() {
  const { activeConnection, loading } = useConnection()

  if (loading) {
    return (
      <div className="min-h-screen flex items-center justify-center bg-zinc-950 text-white">
        Carregando...
      </div>
    )
  }

  if (activeConnection) {
    return <Navigate to="/app" replace />
  }

  return <Outlet />
}