import { useEffect, useState } from "react"
import { API_URL } from "@/api";
import { useDatabase } from "@/contexts/DatabaseProvider";
import { Th } from "@/components/th";
import { Td } from "@/components/td";
import { Button } from "@/components/button";
import { Pencil, Trash } from "lucide-react";
import { ModalFormItem } from "../modal/ModalFormItem";
import type { TableData, Values } from "@/models";
import { ModalDeleteItem } from "../modal/ModalDeleteItem";

export default function SectionTablePreview({ table }: { table: string }) {
  const [tableData, setTableData] = useState<TableData | null>(null)
  const [selectedRow, setSelectedRow] = useState<Values | null>(null)
  const [openForm, setOpenForm] = useState(false)
  const [openDelete, setOpenDelete] = useState(false)
  const { activeDatabase } = useDatabase()

  useEffect(() => {loadTableData()}, [table])
  useEffect(() => {setTableData(null)}, [activeDatabase])

  async function loadTableData() {
    if (!activeDatabase) return
    const res = await fetch(`${API_URL}/api/table/item/${table}`)
    if (!res.ok) {
      setTableData(null)
      return
    }
    const data = await res.json()
    setTableData(data)
  }

  if (!tableData) {
    return <div className="text-zinc-500">{ activeDatabase ? 'Carregando schema...' : 'Selecione uma base de dados' }</div>
  }

  return (
    <div>
      <ModalFormItem open={openForm} onClose={() => setOpenForm(false)} data={selectedRow ?? {}} />
      <ModalDeleteItem open={openDelete} onClose={() => setOpenDelete(false)} data={selectedRow ?? {}} />

      <table className="w-full text-sm text-left">
        <thead className="bg-zinc-900 border-b border-zinc-800">
          <tr>
            {tableData.columns.map((column: string) => <Th key={column}>{column}</Th>)}
            <Th>Ações</Th>
          </tr>
        </thead>

        <tbody>
          {tableData.rows.map((row: Values, index: number) => (
              <tr key={index} className="border-b border-zinc-800 hover:bg-zinc-900/50">
                {tableData.columns.map((column: string) => <Td key={column}>{String(row[column])}</Td>)}
                <Td>
                  <div className="space-x-2">
                    <Button variant="primary" sm onClick={() => { setSelectedRow(row); setOpenForm(true) }}>
                      <Pencil size={16} />
                    </Button>
                    <Button variant="danger" sm onClick={() => { setSelectedRow(row); setOpenDelete(true) }}>
                      <Trash size={16} />
                    </Button>
                  </div>
                </Td>
              </tr>
          ))}
        </tbody>
      </table>
    </div>
  )
}