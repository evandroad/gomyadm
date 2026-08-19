import { Button } from "@/components/button"
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from "@/components/card"
import { Input } from "@/components/input"
import { Label } from "@/components/label"
import ModalBase from "@/components/modalBase"
import { repositories } from "@/repositories"
import { notify } from "@/utils"
import { useEffect, useState } from "react"

type Props = {
  open: boolean
  onClose: () => void
  tableName: string
}

export function ModalFormTable({ open, onClose, tableName }: Props) {
  const [newName, setNewName] = useState<string>('')

  async function handleSubmit() {  
    const res = await repositories.table.update(tableName, newName)

    if (!res.ok) {
      console.error(res.error)
      notify(`Erro: ${res.error || 'Falha ao alterar o dado'}`, 'error')
    } else {
      notify("Dado alterado com sucesso!")
      setNewName('')
    }

     onClose()
  }

  useEffect(() => {
    if (!open) return
    // eslint-disable-next-line react-hooks/set-state-in-effect
    setNewName('')
  }, [open])

  return (
    <ModalBase open={open} onClose={onClose} className="p-2 w-96">
      <Card className="bg-zinc-900">
        <CardHeader>
          <CardTitle onClose={onClose}>Salvar Dados</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="p-2 w-full space-y-4">
            <Label>Nome</Label>
            <Input value={newName || ''} onChange={(e) => setNewName(e.target.value)} />
          </div>
        </CardContent>
        <CardFooter>
          <Button onClick={handleSubmit} variant="primary">Salvar</Button>
          <Button onClick={onClose}>Cancelar</Button>
        </CardFooter>
      </Card>
    </ModalBase>
  )
}