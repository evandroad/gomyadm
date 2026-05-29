import { useEffect, useState } from "react"

import ModalBase from "./components/modalBase"
import { Input } from "./components/input"
import { Button } from "./components/button"
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from "./components/card"
import { Label } from "./components/label"
import { createConnection, type Connection } from "./models"
import { API_URL } from "./api"
import { useConnection } from "./contexts/ConnectionProvider"

export default function SidebarConnection() {
  const { activeConnection, setActiveConnection } = useConnection()
  const [open, setOpen] = useState(false)
  const [connection, setConnection] = useState<Connection>(createConnection())

  useEffect(() => {
    if (activeConnection) setConnection(activeConnection)
  }, [activeConnection])

  async function saveConnection() {
    try {
      const res = await fetch(`${API_URL}/api/connection/save`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(connection)
      })

      const data = await res.json()
      if (data) {
        setConnection(data)
        setActiveConnection(data)
      }
    } catch (err) {
      console.error(err)
    } finally {
      setOpen(false)
    }
  }

  return (
    <>
      <ModalBase open={open} onClose={() => setOpen(false)} className="p-2 w-96">
        <Card className="bg-zinc-900">
          <CardHeader>
            <CardTitle>Salvar conexão</CardTitle>
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
            <Button onClick={() => setOpen(false)}>Cancelar</Button>
            <Button onClick={saveConnection} variant="success">Salvar</Button>
          </CardFooter>
        </Card>
      </ModalBase>
      
      <div className="p-3 border-b border-zinc-800 flex items-center justify-between">
        <div className="text-zinc-400">{activeConnection?.name || "Nova Conexão"}</div>

        <Button onClick={() => setOpen(true)}>Salvar</Button>
      </div>
    </>
  )
}