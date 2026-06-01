import { useEffect, useState } from "react"
import { API_URL } from "../api";
import { useDatabase } from "@/contexts/DatabaseProvider";
import { Th } from "@/components/th";
import { Td } from "@/components/td";

export default function TablePreview({ table }: { table: string }) {
  const [schema, setSchema] = useState<any>(null)
  const { activeDatabase } = useDatabase()

  useEffect(() => {load()}, [table])
  useEffect(() => {setSchema(null)}, [activeDatabase])

  async function load() {
    if (!activeDatabase) return
    const res = await fetch(`${API_URL}/api/connection/tables/${table}`)
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
    <div>
      <table className="w-full text-sm text-left">
        <thead className="bg-zinc-900 border-b border-zinc-800">
          <tr>
            {schema.columns.map((column: string) => <Th key={column}>{column}</Th>)}
          </tr>
        </thead>

        <tbody>
          {schema.rows.map((row: Record<string, any>, index: number) => (
              <tr key={index} className="border-b border-zinc-800 hover:bg-zinc-900/50">
                {schema.columns.map((column: string) => <Td key={column}>{String(row[column])}</Td>)}
              </tr>
          ))}
        </tbody>
      </table>
    </div>
  )
}