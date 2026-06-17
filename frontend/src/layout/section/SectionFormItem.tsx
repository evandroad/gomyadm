import { useState } from "react"
import { API_URL } from "@/api";
import { useDatabase } from "@/contexts/DatabaseProvider";
import { Input } from "@/components/input";
import { Label } from "@/components/label";
import { Button } from "@/components/button";
import { castValue, getInputType } from "@/tableUtils";
import { useTable } from "@/contexts/TableProvider";
import { notify } from "@/utils";
import type { Column } from "@/models";

export default function SectionFormItem() {
  const { activeDatabase } = useDatabase()
  const { activeTable } = useTable()
  const [formData, setFormData] = useState<Record<string, string>>({})

  async function handleSubmit() {
    if (!activeTable) return
    const payload = {
      table: activeTable.name,
      values: Object.fromEntries(
        activeTable.columns.map((column: any) => [column.name, castValue(formData[column.name], column.type)])
      )
    }

    try {
      const res = await fetch(`${API_URL}/api/table/item`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload),
      })

      if (!res.ok) {
        const data = await res.json()
        notify(`Erro: ${data.message || 'Falha ao inserir dados'}`, 'error')
        return
      }

      notify("Dados inseridos com sucesso!")
      setFormData({})
    } catch (err: any) {
      notify(`Erro: ${err.message || 'Falha ao inserir dados'}`, 'error')
    }
  }

  if (!activeDatabase) {
    return <div className="text-zinc-500">{ activeDatabase ? 'Carregando schema...' : 'Selecione uma base de dados' }</div>
  }

  return (
    <div className="p-2 w-100 space-y-4">
      {activeTable?.columns.map((column: Column) => (
        <div key={column.name}>
          <Label>{column.name}</Label>

          <Input
            type={getInputType(column.type)}
            required={!column.nullable}
            disabled={column.autoIncrement}
            value={formData[column.name] || ""}
            onChange={(e) => setFormData(prev => ({ ...prev, [column.name]: e.target.value }))}
          />
        </div>
      ))}

      <Button onClick={handleSubmit} variant="primary">Salvar</Button>
    </div>
  )
}