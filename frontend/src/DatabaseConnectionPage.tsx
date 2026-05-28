import { useState } from "react"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"
import { Label } from "@/components/ui/label"

type Driver = "mysql" | "postgres"

export default function DatabaseConnectionPage() {
  const [driver, setDriver] = useState<Driver>("postgres")
  const [host, setHost] = useState("localhost")
  const [port, setPort] = useState("")
  const [name, setName] = useState("")
  const [username, setUsername] = useState("")
  const [password, setPassword] = useState("")
  const [database, setDatabase] = useState("")
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  async function handleConnect() {
    setLoading(true)
    setError(null)

    try {
      const res = await fetch("http://localhost:8181/api/connections/connect", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ name, driver, host, port: Number(port), username, password, database }),
      })

      if (!res.ok) {
        const data = await res.json()
        throw new Error(data.message || "Falha ao conectar")
      }

      const data = await res.json()
      console.log("Conectado:", data)

    } catch (err: any) {
      setError(err.message)
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="min-h-screen flex items-start justify-center bg-zinc-950 p-2">
      <Card className="w-full max-w-md bg-zinc-900 border-zinc-800 text-white mt-16">
        <CardHeader>
          <CardTitle className="text-xl">Conectar ao banco de dados</CardTitle>
        </CardHeader>

        <CardContent className="space-y-6">
          <div className="space-y-2">
            <Label>Nome</Label>
            <Input value={name} onChange={(e) => setName(e.target.value)} placeholder="Nome da conexão" />
          </div>
          <div className="space-y-4">
            <Label>Banco de dados</Label>
            <select className="w-full bg-zinc-800 p-2 rounded-md" value={driver} onChange={(e) => setDriver(e.target.value as Driver)}>
              <option value="mysql">MySQL</option>
              <option value="postgres">PostgreSQL</option>
            </select>
          </div>
          <div className="space-y-2">
            <Label>Host</Label>
            <Input value={host} onChange={(e) => setHost(e.target.value)} placeholder="localhost" />
          </div>
          <div className="space-y-2">
            <Label>Porta</Label>
            <Input value={port} onChange={(e) => setPort(e.target.value)} placeholder="5432 / 3306" />
          </div>
          <div className="space-y-2">
            <Label>Usuário</Label>
            <Input value={username} onChange={(e) => setUsername(e.target.value)} placeholder="postgres / root" />
          </div>
          <div className="space-y-2">
            <Label>Senha</Label>
            <Input type="password" value={password} onChange={(e) => setPassword(e.target.value)} />
          </div>
          <div className="space-y-2">
            <Label>Database</Label>
            <Input value={database} onChange={(e) => setDatabase(e.target.value)} placeholder="nome_do_banco" />
          </div>

          {error && (
            <div className="text-red-400 text-sm bg-red-950 p-2 rounded-md">
              {error}
            </div>
          )}

          <Button className="w-full" onClick={handleConnect} disabled={loading} >
            {loading ? "Conectando..." : "Conectar"}
          </Button>
        </CardContent>
      </Card>
    </div>
  )
}