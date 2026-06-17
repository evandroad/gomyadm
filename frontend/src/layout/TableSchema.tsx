import { useDatabase } from "@/contexts/DatabaseProvider";
import { Th } from "@/components/th";
import { Td } from "@/components/td";
import type { Column } from "@/models";
import { useTable } from "@/contexts/TableProvider";
import { Button } from "@/components/button";
import { Pencil, Trash } from "lucide-react";
import { useState } from "react";
import { ModalFormColumn } from "./ModalFormColumn";
import { ModalDeleteColumn } from "./ModalDeleteColumn";

export default function TableSchema() {
  const { activeDatabase } = useDatabase()
  const { activeTable } = useTable()
  const [selectedRow, setSelectedRow] = useState<Column | null>(null)
  const [openForm, setOpenForm] = useState(false)
  const [openDelete, setOpenDelete] = useState(false)

  const COLUMNS_LABEL = ['Nome', 'Tipo', 'Tamanho', 'Nulo', 'Chave', 'Exclusivo', 'Auto Inc.','Padrão', 'Ações']

  if (!activeTable) {
    return <div className="text-zinc-500">{ activeDatabase ? 'Carregando schema...' : 'Selecione uma base de dados' }</div>
  }

  return (
    <>
      <ModalFormColumn open={openForm} onClose={() => setOpenForm(false)} data={selectedRow} />
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
                <Td>{row.nullable ? 'V' : ''}</Td>
                <Td>{row.primary ? 'V' : ''}</Td>
                <Td>{row.unique ? 'V' : ''}</Td>
                <Td>{row.autoIncrement ? 'V' : ''}</Td>
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