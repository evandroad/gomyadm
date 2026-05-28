import { useEffect, useState } from "react"
import TablePreview from "./TablePreview"
import { API_URL } from "./api"

type Connection = {
  id: string
  name?: string
}

export default function MainPage() {
  const [connections, setConnections] = useState<Connection[]>([])
  const [activeConnection, setActiveConnection] = useState<string | null>(null)
  const [tables, setTables] = useState<string[]>([])
  const [selectedTable, setSelectedTable] = useState<string | null>(null)
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    async function loadConnections() {
      try {
        const res = await fetch(`${API_URL}/api/connections`)
        const data = await res.json()

        setConnections(data)

        if (data.length > 0) {
          setActiveConnection(data[0].id)
        }
      } finally {
        setLoading(false)
      }
    }

    loadConnections()
  }, [])

  useEffect(() => {
    if (!activeConnection) return

    async function loadTables() {
      const res = await fetch(
        `${API_URL}/api/connections/${activeConnection}/tables`
      )
      const data = await res.json()

      setTables(data)
      setSelectedTable(null)
    }

    loadTables()
  }, [activeConnection])

  if (loading) {
    return (
      <div className="h-screen flex items-center justify-center bg-zinc-950 text-white">
        Carregando...
      </div>
    )
  }

  return (
    <div className="h-screen flex bg-zinc-950 text-white">

      {/* SIDEBAR */}
      <aside className="w-64 bg-zinc-900 border-r border-zinc-800 flex flex-col">
        
        <div className="p-3 font-bold border-b border-zinc-800">
          Gomyadm
        </div>

        {/* connections */}
        <div className="p-2 text-xs text-zinc-400">
          CONEXÕES
        </div>

        <div className="flex-1 overflow-auto">
          {connections.map((conn) => (
            <button
              key={conn.id}
              onClick={() => setActiveConnection(conn.id)}
              className={`w-full text-left px-3 py-2 hover:bg-zinc-800 ${
                activeConnection === conn.id ? "bg-zinc-800" : ""
              }`}
            >
              { conn.name || conn.id }
            </button>
          ))}
        </div>

        {activeConnection && (
          <button
            onClick={async () => {
              const res = await fetch(
                `${API_URL}/api/connections/${activeConnection}/disconnect`,
                { method: "POST" }
              )

              if (res.ok) {
                setActiveConnection(null)
                setTables([])
              }
            }}
            className="p-3 bg-red-600 hover:bg-red-700"
          >
            Desconectar
          </button>
        )}
      </aside>

      {/* MAIN */}
      <main className="flex-1 flex">

        {/* TABLE LIST */}
        <div className="w-72 border-r border-zinc-800 bg-zinc-950">
          <div className="p-3 border-b border-zinc-800 text-sm">
            Tabelas
          </div>

          <div className="overflow-auto">
            {tables.map((table) => (
              <button
                key={table}
                onClick={() => setSelectedTable(table)}
                className={`w-full text-left px-3 py-2 hover:bg-zinc-900 ${
                  selectedTable === table ? "bg-zinc-900" : ""
                }`}
              >
                {table}
              </button>
            ))}
          </div>
        </div>

        {/* CONTENT */}
        <div className="flex-1 p-4">
          {!selectedTable ? (
            <div className="text-zinc-500">
              Selecione uma tabela
            </div>
          ) : (
            <TablePreview
              connectionId={activeConnection!}
              table={selectedTable}
            />
          )}
        </div>

      </main>
    </div>
  )
}