import { Button } from "@/components/button"
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from "@/components/card"
import { Label } from "@/components/label"
import ModalBase from "@/components/modalBase"
import { useTable } from "@/contexts/TableContext";
import type { Column, Values } from "@/models"
import { repositories } from "@/repositories"
import { castValue } from "@/tableUtils"
import { notify } from "@/utils"
import { useEffect } from "react"

type Props = {
  open: boolean
  onClose: () => void
  data: Values
}

export function ModalDeleteItem({ open, onClose, data }: Props) {
  const { activeTable } = useTable()

  async function saveConnection() {
    if (!activeTable) return

    const primaryKeys = activeTable.columns.filter((col: Column) => col.primary)
    const payload = {
      table: activeTable.name,
      key: Object.fromEntries(
        primaryKeys.map((column: any) => [column.name, castValue(data[column.name], column.type)])
      )
    }

    const res = await repositories.item.delete(payload)

    if (!res.ok) {
      console.error(res.error)
      notify(`Erro: ${res.error || 'Falha ao remover o dado'}`, 'error')
      return
    } else {
      notify("Dado removido com sucesso!")
    }

    onClose()
  }

  useEffect(() => {
    if (!open) return
  }, [open])

  return (
    <ModalBase open={open} onClose={onClose} className="p-2 w-96">
      <Card className="bg-zinc-900">
        <CardHeader>
          <CardTitle onClose={onClose}>Excluir Item</CardTitle>
        </CardHeader>
        <CardContent>
          {Object.entries(data).map(([key, value]) => <Label key={key}><span>{key}</span><span>{value}</span></Label>)}
        </CardContent>
        <CardFooter>
          <Button onClick={onClose}>Cancelar</Button>
          <Button onClick={saveConnection} variant="danger">Excluir</Button>
        </CardFooter>
      </Card>
    </ModalBase>
  )
}