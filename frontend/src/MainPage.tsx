import { useEffect, useState } from "react"
import TablePreview from "./layout/TablePreview"
import { API_URL } from "./api"
import { useNavigate } from "react-router-dom"
import SidebarConnection from "./layout/SidebarConnection"
import { Button } from "./components/button"
import { useConnection } from "./contexts/ConnectionProvider"
import { SidebarDatabase } from "./layout/SidebarDatabase"
import { useDatabase } from "./contexts/DatabaseProvider"
import { SidebarTables } from "./layout/SidebarTables"
import { Toolbar } from "./layout/Toolbar"
import TableStructure from "./layout/TableStructure"

export default function MainPage() {
  const navigate = useNavigate()
  const { activeConnection, setActiveConnection, loading } = useConnection()
  const { activeDatabase } = useDatabase()
  
  const [tables, setTables] = useState<string[]>([])
  const [selectedTable, setSelectedTable] = useState<string | null>(null)
  const [view, setView] = useState<"data" | "structure">("data")

  useEffect(() => {
    if (!activeConnection) {
      setSelectedTable(null)
      setTables([])
      return
    }

    loadTables()
  }, [activeDatabase])

  async function loadTables() {
    if (!activeConnection || !activeDatabase) return

    const res = await fetch(`${API_URL}/api/connection/tables`)
    const data = await res.json()
    setTables(data)
    setSelectedTable(null)
  }

  async function disconnect() {
    const res = await fetch(`${API_URL}/api/connection/disconnect`, { method: "POST" })

    if (res.ok) {
      setActiveConnection(null)
      setTables([])
      navigate("/connect", { replace: true })
    }
  }

  if (loading) {
    return (
      <div className="h-screen flex items-center justify-center bg-zinc-950 text-white">
        Carregando...
      </div>
    )
  }

  function renderView() {
    if (!selectedTable) {
      return (
        <div className="text-zinc-500">Selecione uma tabela</div>
      )
    }

    switch (view) {
      case "data":
        return <TablePreview table={selectedTable} />
      case "structure":
        return <TableStructure table={selectedTable} />
      default:
        return null
    }
  }

  return (
    <div className="min-h-screen flex bg-zinc-950 text-white">
      <aside className="w-64 bg-zinc-900 border-r border-zinc-800 flex flex-col">        
        <div className="p-3 font-bold border-b border-zinc-800">Gomyadm</div>
        <SidebarConnection />
        <SidebarDatabase />
        <SidebarTables tables={tables} selectedTable={selectedTable} setSelectedTable={setSelectedTable}/>

        {activeConnection && <Button onClick={() => disconnect()} className="mt-auto mb-2 mx-8" variant="danger">Desconectar</Button>}
      </aside>

      <main className="flex-1 flex flex-col">
        <Toolbar>
          <h1 className="text-lg font-semibold">{selectedTable || "Visão geral"}</h1>
          <div className="space-x-2">
            <Button onClick={() => setView("data")} bg={view === "data" ? "bg-zinc-800" : "bg-zinc-900"}>Dados</Button>
            <Button onClick={() => setView("structure")} bg={view === "structure" ? "bg-zinc-800" : "bg-zinc-900"}>Estrutura</Button>
          </div>
        </Toolbar>
        <div className="flex-1 p-4">
          { renderView() }
        </div>
      </main>
    </div>
  )
}