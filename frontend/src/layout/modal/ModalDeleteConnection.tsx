import { API_URL } from "@/api"
import { Button } from "@/components/button"
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from "@/components/card"
import { Label } from "@/components/label"
import ModalBase from "@/components/modalBase"
import { useConnection } from "@/contexts/ConnectionProvider"
import { useConnections } from "@/contexts/ConnectionsContext"
import type { Connection } from "@/models"
import { useEffect } from "react"

type Props = {
  open: boolean
  onClose: () => void
  data: Connection
}

export function ModalDeleteConnection({ open, onClose, data }: Props) {
  const { setActiveConnection } = useConnection()
  const { removeConnection } = useConnections()
  const id = data.id

  async function saveConnection() {
    try {
      const res = await fetch(`${API_URL}/api/connection/${id}`, {
        method: "DELETE",
        headers: { "Content-Type": "application/json" }
      })

      if (res.ok) {
        removeConnection(id)
        setActiveConnection({ ...data, name: "" })
      }
    } catch (err) {
      console.error(err)
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
          <CardTitle onClose={onClose}>Excluir conexão</CardTitle>
        </CardHeader>
        <CardContent>
          <Label><span>Nome da conexão:</span><span>{data.name}</span></Label>
          <Label><span>Host:</span><span>{data.host}</span></Label>
          <Label><span>Porta:</span><span>{data.port}</span></Label>
          <Label><span>Usuário:</span><span>{data.username}</span></Label>
        </CardContent>
        <CardFooter>
          <Button onClick={onClose}>Cancelar</Button>
          <Button onClick={saveConnection} variant="danger">Excluir</Button>
        </CardFooter>
      </Card>
    </ModalBase>
  )
}