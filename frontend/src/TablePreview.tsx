import { useEffect, useState } from "react"
import { API_URL } from "./api";

export default function TablePreview({ connectionId, table }: { connectionId: string; table: string }) {
  const [schema, setSchema] = useState<any>(null)

  useEffect(() => {
    async function load() {
      const res = await fetch(
        `${API_URL}/api/connections/${connectionId}/tables/${table}`
      )

      const data = await res.json()
      setSchema(data)
    }

    load()
  }, [connectionId, table])

  if (!schema) {
    return <div className="text-zinc-500">Carregando schema...</div>
  }

  return (
    <div>
      <h2 className="text-lg font-bold mb-4">{table}</h2>

      <pre className="text-xs bg-zinc-900 p-3 rounded">
        {JSON.stringify(schema, null, 2)}
      </pre>
    </div>
  )
}