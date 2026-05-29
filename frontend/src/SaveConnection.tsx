import { useState } from "react"

import { Input } from "@/components/ui/input"
import ModalBase from "./components/modalBase"

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
      <ModalBase open={open} onClose={() => setOpen(false)}>
        <h1>Salvar conexão</h1>

        <Input placeholder="Nome da conexão" value={connectionName} onChange={(e) => setConnectionName(e.target.value)} className="bg-zinc-950 border-zinc-800"/>
        <button type="button" onClick={saveConnection} className="bg-zinc-800 hover:bg-zinc-700 px-3 py-2 text-sm">
          Cancelar
        </button>
      </ModalBase>
      
      <div className="p-3 border-b border-zinc-800 flex items-center justify-between">
        <div className="text-zinc-400">
          {connectionName || "Conexão 1"}
        </div>

        <button type="button" onClick={() => setOpen(true)} className="bg-zinc-800 hover:bg-zinc-700 px-3 py-2 text-sm">
          Salvar
        </button>
      </div>
    </>
  )
}