import { useDatabase } from "@/contexts/DatabaseProvider";
import { Th } from "@/components/th";
import { Td } from "@/components/td";
import type { Values } from "@/models";
import { useSchema } from "@/contexts/SchemaProvider";
import { Button } from "@/components/button";
import { Pencil, Trash } from "lucide-react";
import { useState } from "react";

export default function TableStructure() {
  const { activeDatabase } = useDatabase()
  const { activeSchema } = useSchema()
  const [/*selectedRow*/, setSelectedRow] = useState<Values | null>(null)
  const [/*openForm*/, setOpenForm] = useState(false)
  const [/*openDelete*/, setOpenDelete] = useState(false)

  const COLUMNS_LABEL = ['Nome', 'Tipo', 'Nulo', 'Chave', 'Padrão', 'Extra', 'Ações']
  const COLUMNS = ['name', 'type', 'nullable', 'key', 'default', 'extra']

  if (!activeSchema) {
    return <div className="text-zinc-500">{ activeDatabase ? 'Carregando schema...' : 'Selecione uma base de dados' }</div>
  }

  return (
    <>
      <div>
        <table className="w-full text-sm text-left">
          <thead className="bg-zinc-900 border-b border-zinc-800">
            <tr>
              {COLUMNS_LABEL.map((column, index) => <Th key={String(index)}>{column}</Th>)}
            </tr>
          </thead>

          <tbody>
            {activeSchema?.columns.map((row: Values, index: number) => (
              <tr key={index} className="border-b border-zinc-800 hover:bg-zinc-900/50">
                {COLUMNS.map((column) => <Td key={column}>{String(row[column])}</Td>)}
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