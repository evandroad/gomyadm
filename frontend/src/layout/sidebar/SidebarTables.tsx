import { API_URL } from "@/api"
import { Button } from "@/components/button"
import type { View } from "@/pages/MainPage"
import { Pencil, Plus, RefreshCcw, Trash } from "lucide-react"
import { useCallback, useEffect, useState } from "react"
import { ModalFormTable } from "../modal/ModalFormTable"
import { ModalDeleteTable } from "../modal/ModalDeleteTable"
import { useDatabase } from "@/contexts/DatabaseContext";
import { useConnection } from "@/contexts/ConectionContext";

type Props = {
  selectedTable: string | null
  setSelectedTable: (table: string | null) => void
  setView: (view: View) => void
}

export function SidebarTables({ selectedTable, setSelectedTable, setView }: Props) {
  const { activeDatabase } = useDatabase()
  const { activeConnection } = useConnection()
  const [tables, setTables] = useState<string[]>([])
  const [tableToUpdate, setTableToUpdate] = useState<string | null>(null)
  const [tableToDelete, setTableToDelete] = useState<string | null>(null)

  const loadTables = useCallback(async () => {
    if (!activeConnection || !activeDatabase) return

    const res = await fetch(`${API_URL}/api/table`)
    const data = await res.json()
    setTables(data)
    setSelectedTable(null)
  }, [activeConnection, activeDatabase, setSelectedTable])

  useEffect(() => {
    if (!activeConnection) {
      setSelectedTable(null)
      // eslint-disable-next-line react-hooks/set-state-in-effect
      setTables([])
      return
    }

    loadTables()
  }, [activeDatabase, activeConnection, loadTables, setSelectedTable])

  function showFormNewTable() {
    if (!activeDatabase) {
      alert("Selecione um banco de dados!")
      return
    }

    setSelectedTable(null)
    setView("formTable")
  }

  return (
    <>
      <ModalFormTable open={tableToUpdate != null} onClose={() => setTableToUpdate(null)} tableName={tableToUpdate ?? ''} />
      <ModalDeleteTable open={tableToDelete != null} onClose={() => setTableToDelete(null)} tableName={tableToDelete ?? ''} />

      <div className="border-r border-zinc-800 p-3">
        <div className="overflow-auto">
          <div className="flex flex-row justify-between mb-3">
            <span className="font-bold">Tabelas</span>
            <div className="space-x-1">
              <Button sm variant="success" onClick={showFormNewTable}><Plus size={16}/></Button>
              <Button sm onClick={loadTables}><RefreshCcw size={16}/></Button>
            </div>
          </div>
          {activeDatabase && tables?.length > 0 ? (
            tables.map((table) => (
              <div key={table} className="flex items-center gap-2 justify-between mb-2">
                <button
                  key={table}
                  onClick={() => { setSelectedTable(table); setView("data") }}
                  className={`text-left px-3 rounded-md hover:bg-zinc-950 cursor-pointer ${selectedTable === table ? "bg-zinc-800" : ""}`}
                >
                  {table}
                </button>
                <div className="space-x-1">
                  <Button sm variant="primary" onClick={() => setTableToUpdate(table)}><Pencil size={16} /></Button>
                  <Button sm variant="danger" onClick={() => setTableToDelete(table)}><Trash size={16} /></Button>
                </div>
              </div>
            ))
          ) : (
            <div className="text-zinc-500 text-sm">Nenhuma tabela encontrada</div>
          )}
        </div>
      </div>
    </>
  )
}