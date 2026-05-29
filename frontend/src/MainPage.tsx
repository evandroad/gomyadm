import { useEffect, useState } from "react"
import TablePreview from "./TablePreview"
import { API_URL } from "./api"
import { useNavigate } from "react-router-dom"
import SidebarConnection from "./SidebarConnection"
import { Button } from "./components/button"
import { Select } from "./components/select"
import { useConnection } from "./contexts/ConnectionProvider"

export default function MainPage() {
  const [tables, setTables] = useState<string[]>([])
  const [selectedTable, setSelectedTable] = useState<string | null>(null)
  const [selectedDatabase, setSelectedDatabase] = useState<string | null>(null)
  const [loading, setLoading] = useState(true)
  const navigate = useNavigate()
  const { activeConnection, setActiveConnection } = useConnection()

  useEffect(() => {
    if (!activeConnection) {
      setSelectedTable(null)
      setTables([])
      return
    }

    switchDatabase()
  }, [selectedDatabase])

  async function switchDatabase() {
    setLoading(true)
    const res = await fetch(`${API_URL}/api/connection/database/select`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ database: selectedDatabase }),
    })

    if (res.ok) {
      setSelectedTable(null)
      setTables([])
      loadTables()
    }

    setLoading(false)
  }

  async function loadTables() {
    if (!activeConnection || !selectedDatabase) return

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

  return (
    <div className="min-h-screen flex bg-zinc-950 text-white">
      <aside className="w-64 bg-zinc-900 border-r border-zinc-800 flex flex-col">        
        <div className="p-3 font-bold border-b border-zinc-800">
          Gomyadm
        </div>

        <SidebarConnection />

        <div className="p-3 font-bold border-b border-zinc-800">
          Banco de dados
          <div className="space-y-4 mt-2">
            <Select value={selectedDatabase || ''} onChange={(e) => setSelectedDatabase(e.target.value)} options={activeConnection?.databases || []}/>
          </div>
        </div>

        <div className="border-r border-zinc-800">
          <div className="overflow-auto">
            {tables.length > 0 ? (
              tables.map((table) => (
                <button key={table} onClick={() => setSelectedTable(table)} className={`w-full text-left px-3 py-2 hover:bg-zinc-950 cursor-pointer ${selectedTable === table ? "bg-zinc-900" : ""}`}>{ table }</button>
              ))
            ) : (
              <div className="p-3 text-zinc-500 text-sm">Nenhuma tabela encontrada</div>
            )}
          </div>
        </div>

        {activeConnection && <Button onClick={() => disconnect()} className="mt-auto mb-2 mx-8" variant="danger">Desconectar</Button>}
      </aside>

      <main className="flex-1 flex">
        <div className="flex-1 p-4">
          {!selectedTable ? <div className="text-zinc-500">Selecione uma tabela</div> : <TablePreview table={selectedTable}/>}
        </div>
      </main>
    </div>
  )
}