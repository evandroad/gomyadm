import { API_URL } from "@/api"
import { Button } from "@/components/button"
import { useConnection } from "@/contexts/ConectionContext";
import { useDatabase } from "@/contexts/DatabaseContext";
import { useTable } from "@/contexts/TableContext";
import { useNavigate } from "react-router-dom"

export function SidebarDisconnect() {
  const { activeConnection, setActiveConnection } = useConnection()
  const { setActiveDatabase } = useDatabase()
  const { setActiveTable } = useTable()
  const navigate = useNavigate()

  async function disconnect() {
    const res = await fetch(`${API_URL}/api/session`, { method: "DELETE" })

    if (res.ok) {
      setActiveConnection(null)
      setActiveDatabase(null)
      setActiveTable(null)
      navigate("/connect", { replace: true })
    }
  }

  return (
    <>
      {activeConnection && <Button onClick={() => disconnect()} className="mt-auto mb-2 mx-8" variant="danger">Desconectar</Button>}
    </>
  )
}