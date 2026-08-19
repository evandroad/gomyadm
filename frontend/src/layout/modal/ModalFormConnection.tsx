import { Button } from "@/components/button"
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from "@/components/card"
import { Input } from "@/components/input"
import { Label } from "@/components/label"
import ModalBase from "@/components/modalBase"
import { useConnection } from "@/contexts/ConectionContext";
import { useConnections } from "@/contexts/ConnectionsContext"
import type { Connection } from "@/models"
import { repositories } from "@/repositories";
import { useEffect, useState } from "react"

type Props = {
  open: boolean
  onClose: () => void
  data: Connection
  edit?: boolean
}

export function ModalFormConnection({ open, onClose, data, edit }: Props) {
  const { setActiveConnection } = useConnection()
  const { insertConnection, updateConnection } = useConnections()
  const [connection, setConnection] = useState<Connection>(data)

  async function saveConnection() {
    const res = edit ? 
      await repositories.connection.update(connection) : 
      await repositories.connection.create(connection)

    if (res.ok) {
      setActiveConnection(connection)

      if (edit) {
        updateConnection(connection.id, connection)
      } else {
        insertConnection(connection)
      }
    } else {
      console.error(res.error)
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
          <CardTitle onClose={onClose}>Salvar conexão</CardTitle>
        </CardHeader>
        <CardContent>
          <Label>Nome da conexão</Label>
          <Input placeholder="Nome da conexão" value={connection.name ?? ''} onChange={(e) => setConnection({...connection, name: e.target.value})} />

          <Label>Host</Label>
          <Input value={connection.host ?? ''} onChange={(e) => setConnection({...connection, host: e.target.value})} />

          <Label>Porta</Label>
          <Input value={connection.port ?? ''} onChange={(e) => setConnection({...connection, port: e.target.value})} />

          <Label>Usuário</Label>
          <Input value={connection.username ?? ''} onChange={(e) => setConnection({...connection, username: e.target.value})} />

          <Label>Senha</Label>
          <Input value={connection.password ?? ''} onChange={(e) => setConnection({...connection, password: e.target.value})} type="password" />
        </CardContent>
        <CardFooter>
          <Button onClick={onClose}>Cancelar</Button>
          <Button onClick={saveConnection} variant="success">Salvar</Button>
        </CardFooter>
      </Card>
    </ModalBase>
  )
}