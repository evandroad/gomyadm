import { useState } from "react"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"
import { Label } from "@/components/ui/label"
import { API_URL } from "./api"
import { useNavigate } from "react-router-dom"

type Driver = "mysql" | "postgres"

type ConnectionRequest = {
  driver: Driver
  host: string
  port: number
  username: string
  password: string
}

export default function DatabaseConnectionPage() {
  const [connection, setConnection] = useState<ConnectionRequest>({
    driver: "mysql",
    host: "",
    port: 0,
    username: "",
    password: ""
  })
  const [loading, setLoading] = useState(false)
  const [port, setPort] = useState<string>("")
  const [error, setError] = useState<string | null>(null)
  const navigate = useNavigate()

  async function handleConnect() {
    setLoading(true)
    setError(null)

    try {
      connection.port = parseInt(port, 0)
      const res = await fetch( API_URL + "/api/connection/connect", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(connection),
      })

      if (!res.ok) {
        const data = await res.json()
        throw new Error(data.message || "Falha ao conectar")
      }

      navigate("/app", { replace: true })
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
          <div className="space-y-4">
            <Label>Banco de dados</Label>
            <select className="w-full bg-zinc-800 p-2 rounded-md" value={connection.driver} onChange={(e) => setConnection({...connection, driver: e.target.value as Driver})}>
              <option value="mysql">MySQL</option>
              <option value="postgres">PostgreSQL</option>
            </select>
          </div>
          <div className="space-y-2">
            <Label>Host</Label>
            <Input value={connection.host} onChange={(e) => setConnection({...connection, host: e.target.value})} placeholder="localhost" />
          </div>
          <div className="space-y-2">
            <Label>Porta</Label>
            <Input value={port} onChange={(e) => setPort(e.target.value)} placeholder="5432 / 3306" />
          </div>
          <div className="space-y-2">
            <Label>Usuário</Label>
            <Input value={connection.username} onChange={(e) => setConnection({...connection, username: e.target.value})} placeholder="postgres / root" />
          </div>
          <div className="space-y-2">
            <Label>Senha</Label>
            <Input type="password" value={connection.password} onChange={(e) => setConnection({...connection, password: e.target.value})} />
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