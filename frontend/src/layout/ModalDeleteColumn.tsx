import { API_URL } from "@/api"
import { Button } from "@/components/button"
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from "@/components/card"
import { Label } from "@/components/label"
import ModalBase from "@/components/modalBase"
import { useSchema } from "@/contexts/SchemaProvider"
import type { Column } from "@/models"
import { notify } from "@/utils"
import { useEffect } from "react"

type Props = {
  open: boolean
  onClose: () => void
  data: Column | null
}

export function ModalDeleteColumn({ open, onClose, data }: Props) {
  const { activeSchema } = useSchema()

  async function saveConnection() {
    if (!activeSchema) return

    try {
      const res = await fetch(`${API_URL}/api/tables/column/${activeSchema.name}/${data?.name}`, {
        method: "DELETE",
        headers: { "Content-Type": "application/json" }
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