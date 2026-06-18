import { API_URL } from "@/api"
import { Button } from "@/components/button"
import { useConnection } from "@/contexts/ConnectionProvider"
import { useDatabase } from "@/contexts/DatabaseProvider"
import type { View } from "@/pages/MainPage"
import { Pencil, Plus, Trash } from "lucide-react"
import { useEffect, useState } from "react"
import { ModalFormTable } from "../modal/ModalFormTable"

type Props = {
  selectedTable: string | null
  setSelectedTable: (table: string | null) => void
  setView: (view: View) => void
}

export function SidebarTables({ selectedTable, setSelectedTable, setView }: Props) {
  const { activeDatabase } = useDatabase()
  const { activeConnection } = useConnection()
  const [tables, setTables] = useState<string[]>([])
  const [tableToEdit, setTableToEdit] = useState<string | null>(null)

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

    const res = await fetch(`${API_URL}/api/table`)
    const data = await res.json()
    setTables(data)
    setSelectedTable(null)
  }

  function showFormNewTable() {
    if (!activeDatabase) {
      alert("Selecione um banco de dados!")
      return
    }

    setSelectedTable(null)
    setView("formTable")
  }

  function showUpdate(tableName: string) {
    setTableToEdit(tableName)
  }

  return (
    <>
      <ModalFormTable open={tableToEdit != null} onClose={() => setTableToEdit(null)} oldName={tableToEdit ?? ''} />
      <div className="border-r border-zinc-800">
        <div className="overflow-auto">
          <div className="flex justify-end">
            <Button sm className="m-1" variant="primary" onClick={showFormNewTable}><Plus size={16}/></Button>
          </div>
          {activeDatabase && tables?.length > 0 ? (
            tables.map((table) => (
              <div key={table} className="flex items-center m-2 gap-2 justify-between">
                <button
                  key={table}
                  onClick={() => { setSelectedTable(table); setView("data") }}
                  className={`w-30 text-left px-1 py-1 rounded-md hover:bg-zinc-950 cursor-pointer ${selectedTable === table ? "bg-zinc-800" : ""}`}
                >
                  {table}
                </button>
                <div className="space-x-1">
                  <Button sm variant="primary" onClick={() => showUpdate(table)}><Pencil size={16} /></Button>
                  <Button sm variant="danger"><Trash size={16} /></Button>
                </div>
              </div>
            ))
          ) : (
            <div className="p-3 text-zinc-500 text-sm">Nenhuma tabela encontrada</div>
          )}
        </div>
      </div>
    </>
  )
}