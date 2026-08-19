import { Button } from "@/components/button"
import { useConnection } from "@/contexts/ConectionContext";
import { useDatabase } from "@/contexts/DatabaseContext";
import { useTable } from "@/contexts/TableContext";
import { repositories } from "@/repositories";
import { useNavigate } from "react-router-dom"

export function SidebarDisconnect() {
  const { activeConnection, setActiveConnection } = useConnection()
  const { setActiveDatabase } = useDatabase()
  const { setActiveTable } = useTable()
  const navigate = useNavigate()

  async function disconnect() {
    try {
      await repositories.session.disconnect()
      setActiveConnection(null)
      setActiveDatabase(null)
      setActiveTable(null)
      navigate("/connect", { replace: true })
    } catch (error) {
      console.error(error)
    }
  }

  return (
    <>
      {activeConnection && <Button onClick={() => disconnect()} className="mt-auto mb-2 mx-8" variant="danger">Desconectar</Button>}
    </>
  )
}