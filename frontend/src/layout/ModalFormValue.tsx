import { API_URL } from "@/api"
import { Button } from "@/components/button"
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from "@/components/card"
import { Input } from "@/components/input"
import { Label } from "@/components/label"
import ModalBase from "@/components/modalBase"
import { useSchema } from "@/contexts/SchemaProvider"
import type { ColumnSchema, Values } from "@/models"
import { castValue, getInputType } from "@/tableUtils"
import { useEffect, useState } from "react"

type Props = {
  open: boolean
  onClose: () => void
  data: Values
}

export function ModalFormValue({ open, onClose, data }: Props) {
  const [formData, setFormData] = useState<Values>({})
  const { activeSchema } = useSchema()

  async function handleSubmit() {
    if (!activeSchema) return
    const primaryKeys = activeSchema.columns.filter((col: any) => col.key === "PRI")
    const payload = {
      table: activeSchema.name,
      key: Object.fromEntries(primaryKeys.map((column: any) => [column.name, castValue(formData[column.name], column.type)])),
      values: Object.fromEntries(
        activeSchema.columns
          .filter((col: any) => col.key !== "PRI")
          .map((column: any) => [column.name, castValue(formData[column.name], column.type)])
      )
    }
    
    try {
      const res = await fetch(`${API_URL}/api/tables`, {
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
      
          </div>
        </CardContent>
        <CardFooter>
          <Button onClick={handleSubmit} variant="primary">Salvar</Button>
          <Button onClick={onClose}>Cancelar</Button>
        </CardFooter>
      </Card>
    </ModalBase>
  )
}