import { useEffect, useState } from "react"
import { Navigate, Outlet } from "react-router-dom"
import { API_URL } from "./api"

export default function GuestGuard() {
  const [loading, setLoading] = useState(true)
  const [connected, setConnected] = useState(false)

  useEffect(() => {
    async function checkConnection() {
      try {
        const res = await fetch(`${API_URL}/api/connection`)

        if (!res.ok) {
          setConnected(false)
          return
        }

        const data = await res.json()
        const conected = data && data != undefined ? true : false
        setConnected(conected)
      } catch {
        setConnected(false)
      } finally {
        setLoading(false)
      }
    }

    checkConnection()
  }, [])

  if (loading) {
    return (
      <div className="min-h-screen flex items-center justify-center bg-zinc-950 text-white">
        Carregando...
      </div>
    )
  }

  if (connected) {
    return <Navigate to="/app" replace />
  }

  return <Outlet />
}