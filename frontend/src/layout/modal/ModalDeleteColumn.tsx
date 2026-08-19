import { Button } from "@/components/button"
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from "@/components/card"
import { Label } from "@/components/label"
import ModalBase from "@/components/modalBase"
import { useTable } from "@/contexts/TableContext";
import type { Column } from "@/models"
import { repositories } from "@/repositories"
import { notify } from "@/utils"
import { useEffect } from "react"

type Props = {
  open: boolean
  onClose: () => void
  data: Column | null
}

export function ModalDeleteColumn({ open, onClose, data }: Props) {
  const { activeTable } = useTable()

  async function saveConnection() {
    if (!activeTable) return

    const res = await repositories.column.delete(activeTable.name, data?.name ?? '')

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
          <Label><span>Nome:</span><span>{data?.name}</span></Label>
          <Label><span>Tipo:</span><span>{data?.type}</span></Label>
          <Label><span>Tamanho:</span><span>{data?.length}</span></Label>
          <Label><span>Nulo:</span><span>{data?.nullable ? 'Sim' : 'Não'}</span></Label>
          <Label><span>Chave:</span><span>{data?.primary ? 'Sim' : 'Não'}</span></Label>
          <Label><span>Exclusivo:</span><span>{data?.unique ? 'Sim' : 'Não'}</span></Label>
          <Label><span>Auto Inc.:</span><span>{data?.autoIncrement ? 'Sim' : 'Não'}</span></Label>
          <Label><span>Padrão:</span><span>{data?.defaultValue}</span></Label>
        </CardContent>
        <CardFooter>
          <Button onClick={onClose}>Cancelar</Button>
          <Button onClick={saveConnection} variant="danger">Excluir</Button>
        </CardFooter>
      </Card>
    </ModalBase>
  )
}