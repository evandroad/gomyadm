import { useEffect, useState } from "react"
import { API_URL } from "../api";
import { useDatabase } from "@/contexts/DatabaseProvider";
import { Input } from "@/components/input";
import { Label } from "@/components/label";
import { Button } from "@/components/button";
import { castValue, getInputType } from "@/tableUtils";

export default function TableForm({ table }: { table: string }) {
  const { activeDatabase } = useDatabase()
  const [schema, setSchema] = useState<any>(null)
  const [formData, setFormData] = useState<Record<string, string>>({})

  useEffect(() => {load()}, [table])
  useEffect(() => {setSchema(null)}, [activeDatabase])

  async function load() {
    if (!activeDatabase) return
    const res = await fetch(`${API_URL}/api/tables/struct/${table}`)
    if (!res.ok) {
      setSchema(null)
      return
    }
    const data = await res.json()
    setSchema(data)
  }

  async function handleSubmit() {
    const payload = {
      table: schema.name,
      values: Object.fromEntries(
        schema.columns.map((column: any) => [column.name, castValue(formData[column.name], column.type)])
      )
    }

    try {
      const res = await fetch(`${API_URL}/api/tables`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload),
      })

      if (!res.ok) {
        const data = await res.json()
        alert(`Erro: ${data.message || 'Falha ao inserir dados'}`)
        return
      }

      alert("Dados inseridos com sucesso!")
      setFormData({})
    } catch (err: any) {
      alert(`Erro: ${err.message || 'Falha ao inserir dados'}`)
    }
  }

  if (!schema) {
    return <div className="text-zinc-500">{ activeDatabase ? 'Carregando schema...' : 'Selecione uma base de dados' }</div>
  }

  return (
    <div className="p-2 w-100 space-y-4">
      {schema.columns.map((column: any) => (
        <div key={column.name}>
          <Label>{column.name}</Label>

          <Input
            type={getInputType(column.type)}
            required={!column.nullable}
            value={formData[column.name] || ""}
            onChange={(e) => setFormData(prev => ({ ...prev, [column.name]: e.target.value }))}
          />
        </div>
      ))}

      <Button onClick={handleSubmit} variant="primary">Salvar</Button>
    </div>
  )
}