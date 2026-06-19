import { API_URL } from "@/api"
import { Button } from "@/components/button"
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from "@/components/card"
import ModalBase from "@/components/modalBase"
import { notify } from "@/utils"
import { useEffect } from "react"

type Props = {
  open: boolean
  onClose: () => void
  tableName: string
}

export function ModalDeleteTable({ open, onClose, tableName }: Props) {
  async function saveConnection() {
    try {
      const res = await fetch(`${API_URL}/api/table/${tableName}`, {
        method: "DELETE",
        headers: { "Content-Type": "application/json" },
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
          <CardTitle onClose={onClose}>Excluir Item</CardTitle>
        </CardHeader>
        <CardContent>
          Tem certeza que deseja remover a tabela <b>{tableName}</b>?
        </CardContent>
        <CardFooter>
          <Button onClick={onClose}>Cancelar</Button>
          <Button onClick={saveConnection} variant="danger">Excluir</Button>
        </CardFooter>
      </Card>
    </ModalBase>
  )
}