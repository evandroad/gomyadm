import { API_URL } from "@/api"
import { Button } from "@/components/button"
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from "@/components/card"
import { Input } from "@/components/input"
import { Label } from "@/components/label"
import ModalBase from "@/components/modalBase"
import { Select } from "@/components/select"
import { useSchema } from "@/contexts/SchemaProvider"
import { ColumnTypes, createColumn, type Column } from "@/models"
import { notify } from "@/utils"
import { useEffect, useState } from "react"

type Props = {
  open: boolean
  onClose: () => void
  data: Column | null
}

export function ModalFormColumn({ open, onClose, data }: Props) {
  const { activeSchema } = useSchema()
  const [column, setColumn] = useState<Column>(createColumn())
  const [oldName, setOldName] = useState<string>('')

  async function handleSubmit() {
    if (!activeSchema) return
    const value = { ...column, length: Number(column.length) }
    const payload = { table: activeSchema.name, oldName, column: value }
    
    try {
      const res = await fetch(`${API_URL}/api/table/column`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload),
      })

      if (!res.ok) {
        const data = await res.json()
        notify(`Erro: ${data.message || 'Falha ao alterar o dado'}`, 'error')
        return
      }

      notify("Dado alterado com sucesso!")
    } catch (err: any) {
      console.error(err)
      notify(`Erro: ${err.message || 'Falha ao alterar o dado'}`, 'error')
    } finally {
      onClose()
    }
  }

  useEffect(() => {
    if (!open) return
    if (data) {
      setColumn(data)
      setOldName(data.name)
    }
  }, [open])

  return (
    <ModalBase open={open} onClose={onClose} className="p-2 w-96">
      <Card className="bg-zinc-900">
        <CardHeader>
          <CardTitle>Salvar Dados</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="p-2 w-full space-y-4">
            <Label>Nome</Label>
            <Input value={column.name} onChange={(e) => setColumn({ ...column, name: e.target.value })} />
      
            <Label>Tipo</Label>
            <Select value={column.type} onChange={(e) => setColumn({...column, type: e.target.value})} options={ColumnTypes}/>
      
            <Label>Tamanho</Label>
            <Input value={column.length} onChange={(e) => setColumn({ ...column, length: e.target.value })} />
      
            <Label>Padrão</Label>
            <Input value={column.defaultValue} onChange={(e) => setColumn({ ...column, defaultValue: e.target.value })} />
      
            <div className="flex justify-between">
              <div>
                <label htmlFor="nullable" className="mr-2 cursor-pointer">Nulo</label>
                <input
                  type="checkbox"
                  id="nullable"
                  className="scale-115 cursor-pointer"
                  checked={!!column.nullable}
                  onChange={(e) => setColumn({ ...column, nullable: e.target.checked })}
                />
              </div>
      
              <div>
                <label htmlFor="primary" className="mr-2 cursor-pointer">Chave Primária</label>
                <input
                  type="checkbox"
                  id="primary"
                  className="scale-115 cursor-pointer"
                  checked={!!column.primary}
                  onChange={(e) => setColumn({ ...column, primary: e.target.checked })}
                />
              </div>
            </div>
      
            <div className="flex justify-between">
              <div>
                <label htmlFor="unique" className="mr-2 cursor-pointer">Exclusivo</label>
                <input
                  type="checkbox"
                  id="unique"
                  className="scale-115 cursor-pointer"
                  checked={!!column.unique}
                  onChange={(e) => setColumn({ ...column, unique: e.target.checked })}
                />
              </div>
              <div>
                <label htmlFor="autoIncrement" className="mr-2 cursor-pointer">Auto Incremento</label>
                <input
                  type="checkbox"
                  id="autoIncrement"
                  className="scale-115 cursor-pointer"
                  checked={!!column.autoIncrement}
                  onChange={(e) => setColumn({ ...column, autoIncrement: e.target.checked })}
                />
              </div>
            </div>
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