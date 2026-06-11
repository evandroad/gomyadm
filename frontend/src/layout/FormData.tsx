import { useState } from "react"
import { API_URL } from "../api";
import { useDatabase } from "@/contexts/DatabaseProvider";
import { Input } from "@/components/input";
import { Label } from "@/components/label";
import { Button } from "@/components/button";
import { castValue, getInputType } from "@/tableUtils";
import type { ColumnSchema } from "@/models";
import { useSchema } from "@/contexts/SchemaProvider";
import { notify } from "@/utils";

export default function FormData() {
  const { activeDatabase } = useDatabase()
  const { activeSchema } = useSchema()
  const [formData, setFormData] = useState<Record<string, string>>({})

  async function handleSubmit() {
    if (!activeSchema) return
    const payload = {
      table: activeSchema.name,
      values: Object.fromEntries(
        activeSchema.columns.map((column: any) => [column.name, castValue(formData[column.name], column.type)])
      )
    }

    try {
      const res = await fetch(`${API_URL}/api/tables/item`, {
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
      {activeSchema?.columns.map((column: ColumnSchema) => (
        <div key={column.name}>
          <Label>{column.name}</Label>

          <Input
            type={getInputType(column.type)}
            required={!column.nullable}
            disabled={column.autoNumber}
            value={formData[column.name] || ""}
            onChange={(e) => setFormData(prev => ({ ...prev, [column.name]: e.target.value }))}
          />
        </div>
      ))}

      <Button onClick={handleSubmit} variant="primary">Salvar</Button>
    </div>
  )
}