import { API_URL } from "@/api"
import { Button } from "@/components/button"
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from "@/components/card"
import { Input } from "@/components/input"
import { Label } from "@/components/label"
import ModalBase from "@/components/modalBase"
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
    try {
      const res = await fetch(`${API_URL}/api/table/${tableName}/${newName}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" }
      })

      if (!res.ok) {
        const data = await res.json()
        notify(`Erro: ${data.message || 'Falha ao alterar o dado'}`, 'error')
        return
      }

      notify("Dado alterado com sucesso!")
      setNewName('')
    } catch (err: any) {
      console.error(err)
      notify(`Erro: ${err.message || 'Falha ao alterar o dado'}`, 'error')
    } finally {
      onClose()
    }
  }

  useEffect(() => {
    if (!open) return
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