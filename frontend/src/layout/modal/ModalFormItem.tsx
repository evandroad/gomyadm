import { Button } from "@/components/button"
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from "@/components/card"
import { Input } from "@/components/input"
import { Label } from "@/components/label"
import ModalBase from "@/components/modalBase"
import { useTable } from "@/contexts/TableContext";
import type { Column, Values } from "@/models"
import { repositories } from "@/repositories"
import { castValue, getInputType } from "@/tableUtils"
import { notify } from "@/utils"
import { useEffect, useState } from "react"

type Props = {
  open: boolean
  onClose: () => void
  data: Values | null
}

export function ModalFormItem({ open, onClose, data }: Props) {
  const [formData, setFormData] = useState<Values | null>(null)
  const { activeTable } = useTable()

  async function handleSubmit() {
    if (!activeTable) return
    const primaryKeys = activeTable.columns.filter((col: Column) => col.primary)
    const payload = {
      table: activeTable.name,
      key: Object.fromEntries(primaryKeys.map((column: Column) => [column.name, castValue(formData?.[column.name], column.type)])),
      values: Object.fromEntries(
        activeTable.columns
          .filter((col: any) => col.key !== "PRI")
          .map((column: any) => [column.name, castValue(formData?.[column.name], column.type)])
      )
    }
    
    const res = data == null ?
      await repositories.item.create(payload) :
      await repositories.item.update(payload)

    if (!res.ok) {
      console.error(res.error)
      notify(`Erro: ${res.error || 'Falha ao alterar o dado'}`, 'error')
      return
    } else {
      notify("Item salvo com sucesso")
      setFormData({})
    }

    onClose()
  }

  useEffect(() => {
    if (!open) return
    // eslint-disable-next-line react-hooks/set-state-in-effect
    setFormData(data)
  }, [open, data])

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
                  value={formData?.[column.name] || ""}
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