import { Button } from "@/components/button"
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from "@/components/card"
import ModalBase from "@/components/modalBase"
import { repositories } from "@/repositories"
import { notify } from "@/utils"
import { useEffect } from "react"

type Props = {
  open: boolean
  onClose: () => void
  tableName: string
}

export function ModalDeleteTable({ open, onClose, tableName }: Props) {
  async function saveConnection() {
    const res = await repositories.table.delete(tableName)

    if (!res.ok) {
      console.error(res.error)
      notify(`Erro: ${res.error || 'Falha ao remover o dado'}`, 'error')
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