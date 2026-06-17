import { API_URL } from "@/api"
import { Button } from "@/components/button"
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from "@/components/card"
import { Label } from "@/components/label"
import ModalBase from "@/components/modalBase"
import { useTable } from "@/contexts/TableProvider"
import type { Column, Values } from "@/models"
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

    try {
      const res = await fetch(`${API_URL}/api/table/item`, {
        method: "DELETE",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload)
      })

      if (!res.ok) {
        const data = await res.json()
        notify(`Erro: ${data.message || 'Falha ao remover o dado'}`, 'error')
        return
      }

      notify("Dado removido com sucesso!")
    } catch (err: any) {
      console.error(err)
      notify(`Erro: ${err.message || 'Falha ao remover o dado'}`, 'error')
    } finally {
      onClose()
    }
  }

  useEffect(() => {
    if (!open) return
  }, [open])

  return (
    <ModalBase open={open} onClose={onClose} className="p-2 w-96">
      <Card className="bg-zinc-900">
        <CardHeader>
          <CardTitle>Excluir Item</CardTitle>
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