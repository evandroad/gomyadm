import { API_URL } from "@/api"
import { Button } from "@/components/button"
import { useConnection } from "@/contexts/ConnectionProvider"
import { useDatabase } from "@/contexts/DatabaseProvider"
import type { View } from "@/pages/MainPage"
import { Plus } from "lucide-react"
import { useEffect, useState } from "react"

type Props = {
  selectedTable: string | null
  setSelectedTable: (table: string | null) => void
  setView: (view: View) => void
}

export function SidebarTables({ selectedTable, setSelectedTable, setView }: Props) {
  const { activeDatabase } = useDatabase()
  const { activeConnection } = useConnection()
  const [tables, setTables] = useState<string[]>([])

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

    const res = await fetch(`${API_URL}/api/tables`)
    const data = await res.json()
    setTables(data)
    setSelectedTable(null)
  }

  return (
    <div className="border-r border-zinc-800">
      <div className="overflow-auto">
        <div className="flex justify-end">
          <Button sm className="m-1" variant="primary"><Plus size={16}/></Button>
        </div>
        {activeDatabase && tables?.length > 0 ? (
          tables.map((table) => (
            <button
              key={table}
              onClick={() => { setSelectedTable(table); setView("data") }}
              className={`w-full text-left px-3 py-2 hover:bg-zinc-950 cursor-pointer ${selectedTable === table ? "bg-zinc-800" : ""}`}
            >
              {table}
            </button>
          ))
        ) : (
          <div className="p-3 text-zinc-500 text-sm">Nenhuma tabela encontrada</div>
        )}
      </div>
    </div>
  )
}