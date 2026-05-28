import { useEffect, useState } from "react"
import { Navigate } from "react-router-dom"
import { API_URL } from "./api"

export default function ConnectionGuard({ children }: { children: React.ReactNode }) {
  const [loading, setLoading] = useState(true)
  const [connected, setConnected] = useState(false)

  useEffect(() => {
    async function checkConnection() {
      try {
        const res = await fetch(`${API_URL}/api/connections`)

        if (!res.ok) {
          setConnected(false)
          return
        }

        const data = await res.json()
        setConnected(Array.isArray(data) && data.length > 0)
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
      <div className="min-h-screen flex items-center justify-center text-white">
        Conectando...
      </div>
    )
  }

  if (!connected) {
    return <Navigate to="/connect" replace />
  }

  return <>{children}</>
}