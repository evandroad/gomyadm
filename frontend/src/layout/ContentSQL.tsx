import { API_URL } from "@/api"
import { Button } from "@/components/button"
import { useState } from "react"

export function ContentSQL() {
  const [query, setQuery] = useState("")
  const [result, setResult] = useState("")

  async function executeQuery() {
    try {
      const res = await fetch(`${API_URL}/api/query`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ query }),
      })
      const data = await res.json()
      setResult(JSON.stringify(data, null, 2))
    } catch (error) {
      setResult("Erro ao executar consulta: " + error)
    }
  }

  return (
    <div className="p-4 space-y-2 space-x-2">
      <textarea 
        className="w-full h-64 p-2 bg-zinc-800 text-zinc-300 font-mono rounded resize-none focus:outline-none focus:ring-none"
        value={query}
        onChange={(e) => setQuery(e.target.value)}
        placeholder="Escreva sua consulta SQL aqui..."
      />

      <Button variant="primary" onClick={executeQuery}>Executar</Button>
      <Button variant="default" onClick={() => { setResult(""); setQuery("") }}>Limpar Resultado</Button>

      {result && <pre className="p-4 bg-zinc-800 text-white rounded">{ result }</pre>}
    </div>
  )
}