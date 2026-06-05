import { API_URL } from "@/api"
import { Button } from "@/components/button"
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from "@/components/card"
import { Input } from "@/components/input"
import { Label } from "@/components/label"
import ModalBase from "@/components/modalBase"
import type { Values } from "@/models"
import { castValue, getInputType } from "@/tableUtils"
import { useEffect, useState } from "react"

type Props = {
  open: boolean
  onClose: () => void
  data: Values
  table: string
}

export function ModalFormValue({ open, onClose, data, table }: Props) {
  const [formData, setFormData] = useState<Values>({})
  const [schema, setSchema] = useState<any>(null)

  useEffect(() => {load()}, [table])

  async function load() {
    const res = await fetch(`${API_URL}/api/connection/tables/struct/${table}`)
    if (!res.ok) {
      setSchema(null)
      return
    }
    const data = await res.json()
    setSchema(data)
  }

  async function handleSubmit() {
    const primaryKeys = schema.columns.filter((col: any) => col.key === "PRI")
    const payload = {
      table: schema.name,
      key: Object.fromEntries(primaryKeys.map((column: any) => [column.name, castValue(formData[column.name], column.type)])),
      values: Object.fromEntries(
        schema.columns
          .filter((col: any) => col.key !== "PRI")
          .map((column: any) => [column.name, castValue(formData[column.name], column.type)])
      )
    }
    
    try {
      const res = await fetch(`${API_URL}/api/connection/tables`, {
        method: "PUT",
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
      console.error(err)
      alert(`Erro: ${err.message || 'Falha ao inserir dados'}`)
    } finally {
      onClose()
    }
  }

  useEffect(() => {
    if (!open) return
    setFormData(data)
  }, [open])

  return (
    <ModalBase open={open} onClose={onClose} className="p-2 w-96">
      <Card className="bg-zinc-900">
        <CardHeader>
          <CardTitle>Salvar Dados</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="p-2 w-full space-y-4">
            {schema?.columns.map((column: any) => (
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
      
          </div>
        </CardContent>
        <CardFooter>
          <Button onClick={handleSubmit}>Salvar</Button>
          <Button onClick={onClose}>Cancelar</Button>
        </CardFooter>
      </Card>
    </ModalBase>
  )
}