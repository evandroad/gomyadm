import { Button } from "@/components/button";
import { Select } from "@/components/select";
import { Pencil, Plus, RefreshCcw, Trash } from "lucide-react";
import { useState } from "react";
import { ModalFormDatabase } from "../modal/ModalFormDatabase";
import { ModalDeleteDatabase } from "../modal/ModalDeleteDatabase";
import { useDatabase } from "@/contexts/DatabaseContext";

export function SidebarDatabase() {
  const { activeDatabase, databases, changeDatabase, getDatabases } = useDatabase()
  const [openForm, setOpenForm] = useState<boolean>(false)
  const [openDelete, setopenDelete] = useState<boolean>(false)
  const [oldName, setOldName] = useState<string | undefined>(undefined)

  const handleCreate = () => {
    setOldName(undefined)
    setOpenForm(true)
  }

  const handleEdit = () => {
    setOldName(activeDatabase ?? undefined)
    setOpenForm(true)
  }

  const handleCloseForm = () => {
    setOpenForm(false)
    setOldName(undefined)
  }

  return (
    <>
      <ModalFormDatabase open={openForm} onClose={handleCloseForm} oldName={oldName} />
      <ModalDeleteDatabase open={openDelete} onClose={() => setopenDelete(false)} name={activeDatabase ?? ''} />

      <div className="p-3 border-b border-zinc-800">
        <div className="flex flex-row justify-between">
          <span className="font-bold">Banco de dados</span>
          <div className="space-x-1">
            {activeDatabase &&
              <>
                <Button sm variant="primary" onClick={handleEdit}><Pencil size={16} /></Button>
                <Button sm variant="danger"  onClick={() => setopenDelete(true)}><Trash size={16} /></Button>
              </>
            }
            <Button sm variant="success" onClick={handleCreate}><Plus size={16} /></Button>
            <Button sm onClick={getDatabases}><RefreshCcw size={16} /></Button>
          </div>
        </div>
        <div className="space-y-4 mt-2">
          <Select value={activeDatabase || ''} onChange={(e) => changeDatabase(e.target.value)} options={databases}/>
        </div>
      </div>
    </>
  )
}