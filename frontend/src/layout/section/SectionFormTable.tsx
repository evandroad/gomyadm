import { Button } from "@/components/button";
import { Input } from "@/components/input";
import { Label } from "@/components/label";
import { createColumn, createTable, type Column, type Table } from "@/models";
import { useState } from "react";
import { COLUMNS_LABEL } from "./SectionTableSchema";
import { Th } from "@/components/th";
import { Td } from "@/components/td";
import { ModalFormColumn } from "../modal/ModalFormColumn";
import { Trash } from "lucide-react";
import { getStatus, notify } from "@/utils";
import { API_URL } from "@/api";

export default function SectionFormTable() {
  const [table, setTable] = useState<Table>(createTable())
  const [openForm, setOpenForm] = useState<boolean>(false)

  async function handleSubmit() {
    try {
      const res = await fetch(`${API_URL}/api/table`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(table),
      })

      if (!res.ok) {
        const data = await res.json()
        notify(`Erro: ${data.message || 'Falha ao alterar o dado'}`, 'error')
        return
      }

      notify("Dado alterado com sucesso!")
    } catch (err: any) {
      console.error(err)
      notify(`Erro: ${err.message || 'Falha ao alterar o dado'}`, 'error')
    } finally {
      setTable(createTable())
    }
  }

  function addCol(column: Column) {
    setTable(prev => ({ ...prev, columns: [...prev.columns, column] }))
    setOpenForm(false)
  }

  function delCol(column: Column) {
    setTable(prev => ({
      ...prev,
      columns: prev.columns.filter(
        col => !(col.name === column.name && col.type === column.type)
      )
    }))
  }

  return (
    <>
      <ModalFormColumn open={openForm} toApi={false} onClose={() => setOpenForm(false)} onSave={addCol} data={createColumn()} />

      <div className="p-2 w-100 space-y-4">
        <Label>Name</Label>
        <Input value={table.name} onChange={(e) => setTable({ ...table, name: e.target.value })} />

        <Button variant="primary" onClick={() => setOpenForm(true)} className="mr-1">Inserir Coluna</Button>
        <Button variant="success" onClick={handleSubmit}>Salvar Tabela</Button>
      </div>

      <table className="w-full text-sm text-left">
        <thead className="bg-zinc-900 border-b border-zinc-800">
          <tr>
            {COLUMNS_LABEL.map((column, index) => <Th key={String(index)}>{column}</Th>)}
          </tr>
        </thead>
        <tbody>
          {table.columns.map((row: Column, index: number) => (
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
                  <Button variant="danger" sm onClick={() => delCol(row) }>
                    <Trash size={16} />
                  </Button>
                </div>
              </Td>
            </tr>
          ))}
        </tbody>
      </table>
    </>
  )
}