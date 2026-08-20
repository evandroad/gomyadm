import { Button } from "@/components/button"
import { repositories } from "@/repositories"
import { useState } from "react"

export function SectionContentSQL() {
  const [query, setQuery] = useState("")
  const [result, setResult] = useState("")

  async function executeQuery() {
    const res = await repositories.query.query(query)
    
    if (res.ok) {
      setResult(JSON.stringify(res.data, null, 2))
    } else {
      setResult("Erro ao executar consulta: " + res.error)
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