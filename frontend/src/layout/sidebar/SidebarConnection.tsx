import { useState } from "react"
import { Button } from "@/components/button"
import { useConnection } from "@/contexts/ConnectionProvider"
import { useConnections } from "@/contexts/ConnectionsProvider"
import { Pencil, Trash } from "lucide-react"
import { ModalFormConnection } from "../modal/ModalFormConnection"
import { ModalDeleteConnection } from "../modal/ModalDeleteConnection"
import { createConnection } from "@/models"

export default function SidebarConnection() {
  const { activeConnection } = useConnection()
  const [openForm, setOpenForm] = useState(false)
  const [openDelete, setOpenDelete] = useState(false)
  const { getConnectionById } = useConnections()

  const alreadySaved = !!activeConnection && !!getConnectionById(activeConnection.id)

  return (
    <>
      <ModalFormConnection open={openForm} onClose={() => setOpenForm(false)} data={activeConnection ?? createConnection()} edit={alreadySaved} />
      <ModalDeleteConnection open={openDelete} onClose={() => setOpenDelete(false)} data={activeConnection ?? createConnection()} />

      <div className="p-3 border-b border-zinc-800 flex items-center justify-between">
        <div className="text-zinc-400">{activeConnection?.name || "Nova Conexão"}</div>
        {alreadySaved ?
          <div className="space-x-1">
            <Button onClick={() => setOpenForm(true)} variant="primary" sm><Pencil size={16} /></Button>
            <Button onClick={() => setOpenDelete(true)} variant="danger" sm><Trash size={16} /></Button>
          </div> :
          <Button onClick={() => setOpenForm(true)}>Salvar</Button>
        }
      </div>
    </>
  )
}