import { useEffect, useState } from "react"
import TablePreview from "./TablePreview"
import { API_URL } from "./api"
import { useNavigate } from "react-router-dom"

type Connection = {
  id: string
  name: string
  database: string
  databases: string[]
}

export default function MainPage() {
  const [connection, setConnection] = useState<Connection | null>(null)
  const [tables, setTables] = useState<string[]>([])
  const [selectedTable, setSelectedTable] = useState<string | null>(null)
  const [selectedDatabase, setSelectedDatabase] = useState<string | null>(null)
  const [loading, setLoading] = useState(true)
  const navigate = useNavigate()

  useEffect(() => {
    async function loadConnections() {
      try {
        const res = await fetch(`${API_URL}/api/connection`)
        const data = await res.json()
        
        if (data) {
          setConnection(data)
          if (data.database != '') {
            setSelectedDatabase(data.database)
          }
        }
      } finally {
        setLoading(false)
      }
    }

    loadConnections()
  }, [])

  useEffect(() => {
    if (!connection) {
      setSelectedTable(null)
      setTables([])
      return
    }

    async function switchDatabase() {
      setLoading(true)
      const res = await fetch(`${API_URL}/api/connection/database/select`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ database: selectedDatabase }),
      })

      if (res.ok) {
        setSelectedTable(null)
        setTables([])
        loadTables()
      }

      setLoading(false)
    }

    switchDatabase()
  }, [selectedDatabase])

  async function loadTables() {
    if (!connection || !selectedDatabase) return

    const res = await fetch(`${API_URL}/api/connection/tables`)
    const data = await res.json()
    setTables(data)
    setSelectedTable(null)
  }

  async function disconnect() {
    const res = await fetch(`${API_URL}/api/connection/disconnect`, { method: "POST" })

    if (res.ok) {
      setConnection(null)
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

        <div className="p-3 text-zinc-400 border-b border-zinc-800">
          { connection?.name || 'Conexão 1' }
        </div>

        <div className="p-3 font-bold border-b border-zinc-800">
          Banco de dados
          <div className="space-y-4 mt-2">
            <select className="w-full bg-zinc-800 p-2 rounded-md cursor-pointer" value={selectedDatabase || ''} onChange={(e) => setSelectedDatabase(e.target.value)}>
              <option value="">Selecione</option>
              {connection?.databases.map((db) => (
                <option key={db} value={db}>{db}</option>
              ))}
            </select>
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

        {connection && (
          <button onClick={() => disconnect()} className="mt-auto p-3 bg-red-600 hover:bg-red-700">
            Desconectar
          </button>
        )}
      </aside>

      <main className="flex-1 flex">
        <div className="flex-1 p-4">
          {!selectedTable ? (
            <div className="text-zinc-500">Selecione uma tabela</div>
          ) : (
            <TablePreview table={selectedTable}/>
          )}
        </div>
      </main>
    </div>
  )
}