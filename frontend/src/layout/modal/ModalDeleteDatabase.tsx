import { Button } from "@/components/button"
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from "@/components/card"
import ModalBase from "@/components/modalBase"
import { repositories } from "@/repositories"
import { notify } from "@/utils"

type Props = {
  open: boolean
  onClose: () => void
  name: string
}

export function ModalDeleteDatabase({ open, onClose, name }: Props) {
  async function saveConnection() {
    const res = await repositories.database.delete(name)

    if (!res.ok) {
      console.error(res.error)
      notify(`Erro: ${res.error || 'Falha ao remover o dado'}`, 'error')
    } else {
      notify("Dado removido com sucesso!")
    }

    onClose()
  }

  return (
    <ModalBase open={open} onClose={onClose} className="p-2 w-96">
      <Card className="bg-zinc-900">
        <CardHeader>
          <CardTitle onClose={onClose}>Excluir Item</CardTitle>
        </CardHeader>
        <CardContent>
          Tem certeza que deseja remover o banco de dados <b>{name}</b>?
        </CardContent>
        <CardFooter>
          <Button onClick={onClose}>Cancelar</Button>
          <Button onClick={saveConnection} variant="danger">Excluir</Button>
        </CardFooter>
      </Card>
    </ModalBase>
  )
}