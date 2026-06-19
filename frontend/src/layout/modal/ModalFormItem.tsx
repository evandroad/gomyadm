import { API_URL } from "@/api"
import { Button } from "@/components/button"
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from "@/components/card"
import { Input } from "@/components/input"
import { Label } from "@/components/label"
import ModalBase from "@/components/modalBase"
import { useTable } from "@/contexts/TableProvider"
import type { Column, Values } from "@/models"
import { castValue, getInputType } from "@/tableUtils"
import { notify } from "@/utils"
import { useEffect, useState } from "react"

type Props = {
  open: boolean
  onClose: () => void
  data: Values
}

export function ModalFormItem({ open, onClose, data }: Props) {
  const [formData, setFormData] = useState<Values>({})
  const { activeTable } = useTable()

  async function handleSubmit() {
    if (!activeTable) return
    const primaryKeys = activeTable.columns.filter((col: Column) => col.primary)
    const payload = {
      table: activeTable.name,
      key: Object.fromEntries(primaryKeys.map((column: Column) => [column.name, castValue(formData[column.name], column.type)])),
      values: Object.fromEntries(
        activeTable.columns
          .filter((col: any) => col.key !== "PRI")
          .map((column: any) => [column.name, castValue(formData[column.name], column.type)])
      )
    }
    
    try {
      const res = await fetch(`${API_URL}/api/table/item`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload),
      })

      if (!res.ok) {
        const data = await res.json()
        notify(`Erro: ${data.message || 'Falha ao alterar o dado'}`, 'error')
        return
      }

      notify("Dado alterado com sucesso!")
      setFormData({})
    } catch (err: any) {
      console.error(err)
      notify(`Erro: ${err.message || 'Falha ao alterar o dado'}`, 'error')
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
          <CardTitle onClose={onClose}>Salvar Dados</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="p-2 w-full space-y-4">
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