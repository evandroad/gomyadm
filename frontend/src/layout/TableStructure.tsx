import { useEffect, useState } from "react"
import { API_URL } from "../api";
import { useDatabase } from "@/contexts/DatabaseProvider";
import { Th } from "@/components/th";
import { Td } from "@/components/td";
import type { TableStructure, Values } from "@/models";

export default function TableStructure({ table }: { table: string }) {
  const [schema, setSchema] = useState<TableStructure | null>(null)
  const { activeDatabase } = useDatabase()

  useEffect(() => {load()}, [table])
  useEffect(() => {setSchema(null)}, [activeDatabase])

  const COLUMNS_LABEL = ['Nome', 'Tipo', 'Nulo', 'Chave', 'Padrão', 'Extra']
  const COLUMNS = ['name', 'type', 'nullable', 'key', 'default', 'extra']

  async function load() {
    if (!activeDatabase) return
    const res = await fetch(`${API_URL}/api/tables/struct/${table}`)
    if (!res.ok) {
      setSchema(null)
      return
    }
    const data = await res.json()
    setSchema(data)
  }

  if (!schema) {
    return <div className="text-zinc-500">{ activeDatabase ? 'Carregando schema...' : 'Selecione uma base de dados' }</div>
  }

  return (
    <>
      <div>
        <table className="w-full text-sm text-left">
          <thead className="bg-zinc-900 border-b border-zinc-800">
            <tr>
              {COLUMNS_LABEL.map((column, index) => <Th key={String(index)}>{column}</Th>)}
            </tr>
          </thead>

          <tbody>
            {schema.columns.map((row: Values, index: number) => (
              <tr key={index} className="border-b border-zinc-800 hover:bg-zinc-900/50">
                {COLUMNS.map((column) => <Td key={column}>{String(row[column])}</Td>)}
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </>
  )
}