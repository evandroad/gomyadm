import { useEffect, useState } from "react"
import { API_URL } from "./api";

export default function TablePreview({ table }: { table: string }) {
  const [schema, setSchema] = useState<any>(null)

  useEffect(() => {
    load()
  }, [table])
    
  async function load() {
    const res = await fetch(`${API_URL}/api/connection/tables/${table}`)
    if (!res.ok) {
      setSchema(null)
      return
    }
    const data = await res.json()
    setSchema(data)
  }

  if (!schema) {
    return <div className="text-zinc-500">Carregando schema...</div>
  }

  return (
    <div>
      <h2 className="text-lg font-bold mb-4">{table}</h2>

      <table className="w-full text-sm text-left">
        <thead className="bg-zinc-900 border-b border-zinc-800">
          <tr>
            {schema.columns.map((column: string) => (
                <th key={column} className="px-4 py-3 font-medium text-zinc-300 whitespace-nowrap">
                  {column}
                </th>
              ))}
          </tr>
        </thead>

        <tbody>
          {schema.rows.map((row: Record<string, any>, index: number) => (
              <tr key={index} className="border-b border-zinc-800 hover:bg-zinc-900/50">
                {schema.columns.map((column: string) => (
                  <td key={column} className="px-4 py-3 whitespace-nowrap">
                    {String(row[column])}
                  </td>
                ))}
              </tr>
          ))}
        </tbody>
      </table>
    </div>
  )
}