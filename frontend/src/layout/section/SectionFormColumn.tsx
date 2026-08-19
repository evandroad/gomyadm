import { useState } from "react"
import { Input } from "@/components/input";
import { Label } from "@/components/label";
import { Button } from "@/components/button";
import { notify } from "@/utils";
import { Select } from "@/components/select";
import { ColumnTypes, createColumn, type Column } from "@/models";
import { useDatabase } from "@/contexts/DatabaseContext";
import { useTable } from "@/contexts/TableContext";
import { repositories } from "@/repositories";

export default function SectionFormColumn() {
  const { activeDatabase } = useDatabase()
  const { activeTable } = useTable()
  const [column, setColumn] = useState<Column>(createColumn())

  async function handleSubmit() {
    if (!activeTable) return

    const values = { ...column, length: column.length === "" ? undefined : Number(column.length) }
    const payload = { table: activeTable.name, column: values }
    const res = await repositories.column.create(payload)

    if (!res.ok) {
      notify(`Erro: ${res.error || 'Falha ao inserir dados'}`, 'error')
      return
    }

    notify("Dados inseridos com sucesso!")
    setColumn(createColumn())
  }

  if (!activeDatabase) {
    return <div className="text-zinc-500">{ activeDatabase ? 'Carregando schema...' : 'Selecione uma base de dados' }</div>
  }

  return (
    <div className="p-2 w-100 space-y-4">
      <Label>Name</Label>
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

      <Button onClick={handleSubmit} variant="primary">Salvar</Button>
    </div>
  )
}