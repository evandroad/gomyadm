import { useEffect, useState } from "react"
import { API_URL } from "./api"
import { useNavigate } from "react-router-dom"
import { Button } from "./components/button"
import { Card, CardContent, CardHeader, CardTitle } from "./components/card"
import { Label } from "./components/label"
import { Input } from "./components/input"

type Driver = "mysql" | "postgres"

type ConnectionRequest = {
  driver: Driver
  host: string
  port: number
  username: string
  password: string
  name: string
}

export default function DatabaseConnectionPage() {
  const [connection, setConnection] = useState<ConnectionRequest>({
    driver: "mysql",
    host: "",
    port: 0,
    username: "",
    password: "",
    name: ""
  })
  const [connections, setConnections] = useState<ConnectionRequest[]>([])
  const [loading, setLoading] = useState(false)
  const [port, setPort] = useState<string>("")
  const [error, setError] = useState<string | null>(null)
  const navigate = useNavigate()

  useEffect(() => {
    loadConnections()
  }, [])

  async function loadConnections() {
    try {
      const res = await fetch(`${API_URL}/api/connection/list`)
      const data = await res.json()
      
      if (data) {
        setConnections(data)
        if (data.port) {
          setPort(data.port.toString())
        }
      }
    } catch (err) {
      console.error(err)
    }
  }

  async function handleConnect(connArg?: ConnectionRequest) {
    setLoading(true)
    setError(null)

    try {
      const conn = connArg || {
        ...connection,
        port: parseInt(port, 10),
      }
      const res = await fetch( API_URL + "/api/connection/connect", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(conn),
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

  function connectFromSaved(conn: ConnectionRequest) {
    setConnection(conn)
    setPort(conn.port.toString())
    handleConnect(conn)
  }

  return (
    <div className="min-h-screen flex flex-col items-center bg-zinc-950 p-2">
      <Card className="w-full max-w-md bg-zinc-900 border-zinc-800 mt-6">
        { connections.length > 0 && (
          <CardContent className="space-y-4">
            <CardTitle className="text-lg">Conexões salvas</CardTitle>

            <div className="space-y-2">
              {connections.map((conn, index) => (
                <div key={index} className="p-2 bg-zinc-950 rounded-md flex justify-between items-center">
                  <span>
                    <strong className="text-base">{conn.name}</strong> - {conn.host}:{conn.port} - {conn.driver.toUpperCase()}
                  </span>
                  <Button onClick={() => connectFromSaved(conn)}>Conectar</Button>
                </div>
              ))}
            </div>
          </CardContent>
        )}
      </Card>

      <Card className="w-full max-w-md bg-zinc-900 border-zinc-800 text-white mt-6">
        <CardHeader>
          <CardTitle className="text-xl">Conectar ao banco de dados</CardTitle>
        </CardHeader>

        <CardContent>
          <Label>Banco de dados</Label>
          <select className="w-full bg-zinc-800 p-2 rounded-md" value={connection.driver} onChange={(e) => setConnection({...connection, driver: e.target.value as Driver})}>
            <option value="mysql">MySQL</option>
            <option value="postgres">PostgreSQL</option>
          </select>

          <Label>Host</Label>
          <Input value={connection.host} onChange={(e) => setConnection({...connection, host: e.target.value})} />

          <Label>Porta</Label>
          <Input value={port} onChange={(e) => setPort(e.target.value)} />

          <Label>Usuário</Label>
          <Input value={connection.username} onChange={(e) => setConnection({...connection, username: e.target.value})} />

          <Label>Senha</Label>
          <Input type="password" value={connection.password} onChange={(e) => setConnection({...connection, password: e.target.value})} />

          {error && (
            <div className="text-red-400 text-sm bg-red-950 p-2 rounded-md">
              {error}
            </div>
          )}

          <Button className="w-full" onClick={() => handleConnect()} disabled={loading} >
            {loading ? "Conectando..." : "Conectar"}
          </Button>
        </CardContent>
      </Card>
    </div>
  )
}