import { useState } from "react"
import { API_URL } from "./api"
import { useNavigate } from "react-router-dom"
import { Button } from "./components/button"
import { Card, CardContent, CardTitle } from "./components/card"
import { Label } from "./components/label"
import { Input } from "./components/input"
import { Select } from "./components/select"
import { createConnection, type Connection } from "./models"
import { useConnection } from "./contexts/ConnectionProvider"
import { useConnections } from "./contexts/ConnectionsProvider"

const DB_DRIVERS = [
  { value: "mysql", text: "MySQL" },
  { value: "postgres", text: "PostgreSQL" },
  { value: "sqlite", text: "SQLite" }
]

export default function DatabaseConnectionPage() {
  const [connection, setConnection] = useState<Connection>(createConnection())
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)
  const navigate = useNavigate()
  const { setActiveConnection } = useConnection()
  const { connections } = useConnections()

  async function handleConnect(connArg?: Connection) {
    setLoading(true)
    setError(null)

    try {
      const conn = connArg || {
        ...connection,
        port: parseInt(connection.port as string, 10),
      }
      const res = await fetch( API_URL + "/api/connection/connect", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(conn),
      })

      const data = await res.json()

      if (!res.ok) {
        throw new Error(data.message || "Falha ao conectar")
      }

      setActiveConnection(data)
      navigate("/app", { replace: true })
    } catch (err: any) {
      setError(err.message)
    } finally {
      setLoading(false)
    }
  }

  function connectFromSaved(conn: Connection) {
    setConnection(conn)
    handleConnect(conn)
  }

  return (
    <div className="min-h-screen flex flex-col items-center bg-zinc-950 p-2">
      <Card className="w-full max-w-md bg-zinc-900 border-zinc-800 mt-6">
        {connections.length > 0 ? (
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
        ) : (
          <CardContent>
            <CardTitle className="text-lg">Nenhuma conexão salva</CardTitle>
            <p className="text-zinc-500">Configure uma nova conexão abaixo</p>
          </CardContent>
        )}
      </Card>

      <Card className="w-full max-w-md bg-zinc-900 border-zinc-800 text-white mt-6" title="Conectar ao banco de dados" collapsible>
        <CardContent>
          <Label>Banco de dados</Label>
          <Select value={connection.driver} onChange={(e) => setConnection({...connection, driver: e.target.value})} options={DB_DRIVERS}/>

          <Label>Host</Label>
          <Input value={connection.host} onChange={(e) => setConnection({...connection, host: e.target.value})} />

          <Label>Porta</Label>
          <Input value={connection.port} onChange={(e) => setConnection({...connection, port: e.target.value})} />

          <Label>Usuário</Label>
          <Input value={connection.username} onChange={(e) => setConnection({...connection, username: e.target.value})} />

          <Label>Senha</Label>
          <Input value={connection.password} onChange={(e) => setConnection({...connection, password: e.target.value})} type="password" />

          {error && <div className="text-red-400 text-sm bg-red-950 p-2 rounded-md">{error}</div>}

          <Button className="w-full" onClick={() => handleConnect()} disabled={loading}>{loading ? "Conectando..." : "Conectar"}</Button>
        </CardContent>
      </Card>
    </div>
  )
}