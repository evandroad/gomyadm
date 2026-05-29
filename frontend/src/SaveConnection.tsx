import { useState } from "react"

import ModalBase from "./components/modalBase"
import { Input } from "./components/input"
import { Button } from "./components/button"
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from "./components/card"
import { Label } from "./components/label"

export default function SidebarConnection() {
  const [open, setOpen] = useState(false)
  const [connectionName, setConnectionName] = useState("")

  async function saveConnection() {
    try {
      // exemplo:
      // await fetch(`${API_URL}/api/connections/save`, {...})

      console.log("Salvar conexão:", connectionName)

      setOpen(false)
      setConnectionName("")
    } catch (err) {
      console.error(err)
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
            <Input placeholder="Nome da conexão" value={connectionName} onChange={(e) => setConnectionName(e.target.value)} />
          </CardContent>
          <CardFooter>
            <Button onClick={saveConnection}>Cancelar</Button>
            <Button onClick={saveConnection} variant="success">Salvar</Button>
          </CardFooter>
        </Card>
      </ModalBase>
      
      <div className="p-3 border-b border-zinc-800 flex items-center justify-between">
        <div className="text-zinc-400">
          {connectionName || "Conexão 1"}
        </div>

        <Button onClick={() => setOpen(true)}>Salvar</Button>
      </div>
    </>
  )
}