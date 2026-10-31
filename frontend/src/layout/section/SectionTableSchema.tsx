import { Th } from "@/components/th";
import { Td } from "@/components/td";
import type { Column } from "@/models";
import { Button } from "@/components/button";
import { Pencil, Trash } from "lucide-react";
import { useState } from "react";
import { ModalFormColumn } from "../modal/ModalFormColumn";
import { ModalDeleteColumn } from "../modal/ModalDeleteColumn";
import { COLUMNS_LABEL, getStatus } from "@/utils";
import { useDatabase } from "@/contexts/DatabaseContext";
import { useTable } from "@/contexts/TableContext";

export default function SectionTableSchema() {
  const { activeDatabase } = useDatabase()
  const { activeTable } = useTable()
  const [selectedRow, setSelectedRow] = useState<Column | null>(null)
  const [openForm, setOpenForm] = useState(false)
  const [openDelete, setOpenDelete] = useState(false)

  if (!activeTable) {
    return <div className="text-zinc-500">{ activeDatabase ? 'Carregando schema...' : 'Selecione uma base de dados' }</div>
  }

  return (
    <>
      <ModalFormColumn open={openForm} toApi={true} onClose={() => setOpenForm(false)} onSave={() => {}} data={selectedRow} />
      <ModalDeleteColumn open={openDelete} onClose={() => setOpenDelete(false)} data={selectedRow} />

      <div>
        <table className="w-full text-sm text-left">
          <thead className="bg-zinc-900 border-b border-zinc-800">
            <tr>
              {COLUMNS_LABEL.map((column, index) => <Th key={String(index)}>{column}</Th>)}
            </tr>
          </thead>
          <tbody>
            {activeTable?.columns.map((row: Column, index: number) => (
              <tr key={index} className="border-b border-zinc-800 hover:bg-zinc-900/50">
                <Td>{row.name}</Td>
                <Td>{row.type}</Td>
                <Td>{row.length ?? ''}</Td>
                <Td>{getStatus(row.nullable)}</Td>
                <Td>{getStatus(row.primary)}</Td>
                <Td>{getStatus(row.unique)}</Td>
                <Td>{getStatus(row.autoIncrement)}</Td>
                <Td>{row.defaultValue}</Td>
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
    </>
  )
}