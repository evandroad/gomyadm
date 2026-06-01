import { API_URL } from "@/api"
import { Button } from "@/components/button"
import { useConnection } from "@/contexts/ConnectionProvider"
import { useNavigate } from "react-router-dom"

export function SidebarDisconnect() {
  const { activeConnection, setActiveConnection } = useConnection()
  const navigate = useNavigate()

  async function disconnect() {
    const res = await fetch(`${API_URL}/api/connection/disconnect`, { method: "POST" })

    if (res.ok) {
      setActiveConnection(null)
      navigate("/connect", { replace: true })
    }
  }

  return (
    <>
      {activeConnection && <Button onClick={() => disconnect()} className="mt-auto mb-2 mx-8" variant="danger">Desconectar</Button>}
    </>
  )
}