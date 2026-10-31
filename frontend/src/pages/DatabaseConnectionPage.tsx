import { useState } from "react"
import { useNavigate } from "react-router-dom"
import { Button } from "@/components/button"
import { Card, CardContent, CardTitle } from "@/components/card"
import { Label } from "@/components/label"
import { Input } from "@/components/input"
import { Select } from "@/components/select"
import { createConnection, type Connection } from "@/models"
import { useConnections } from "@/contexts/ConnectionsContext"
import { useDatabase } from "@/contexts/DatabaseContext";
import { useConnection } from "@/contexts/ConectionContext";
import { repositories } from "@/repositories";
import { OpenSQLiteFile } from "../../wailsjs/go/bindings/FileDialogBinding"

const DB_DRIVERS = [
  { value: "mysql", text: "MySQL" },
  { value: "postgres", text: "PostgreSQL" },
  { value: "sqlite", text: "SQLite" }
]

export default function DatabaseConnectionPage() {
  const [connection, setConnection] = useState<Connection>(createConnection())
  const [sqliteFileName, setSqliteFileName] = useState("")
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)
  const navigate = useNavigate()
  const { setActiveConnection } = useConnection()
  const { connections } = useConnections()
  const { getDatabases } = useDatabase()

  async function handleConnect(connArg?: Connection) {
    setLoading(true)
    setError(null)

    const conn = connArg || {
      ...connection,
      port: parseInt(connection.port as string, 10),
    }

    const res = await repositories.session.connect(conn)

    if (!res.ok) {
      setError(res.error)
    } else {
      res.data.username = conn.username
      res.data.password = conn.password
      setActiveConnection(res.data)
      getDatabases()
      navigate("/app", { replace: true })
    }

    setLoading(false)
  }

  function connectFromSaved(conn: Connection) {
    setConnection(conn)
    setSqliteFileName(conn.filePath.split(/[\\/]/).pop() ?? "")
    handleConnect(conn)
  }

  function handleDriverChange(driver: string) {
    setConnection({
      ...connection,
      driver,
      filePath: driver === "sqlite" ? connection.filePath : "",
    })
    if (driver !== "sqlite") setSqliteFileName("")
  }

  async function selectSQLiteFile() {
    setError(null)
    try {
      const filePath = await OpenSQLiteFile()
      if (!filePath) return

      setConnection({ ...connection, filePath })
      setSqliteFileName(filePath.split(/[\\/]/).pop() ?? filePath)
    } catch (error) {
      setError(error instanceof Error ? error.message : "Erro ao selecionar arquivo SQLite")
    }
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
                    <strong className="text-base">{conn.name}</strong> - {conn.driver === "sqlite" ? conn.filePath : `${conn.host}:${conn.port}`} - {conn.driver.toUpperCase()}
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
          <Select value={connection.driver} onChange={(e) => handleDriverChange(e.target.value)} options={DB_DRIVERS}/>

          {connection.driver === "sqlite" ? (
            <div className="space-y-2">
              {import.meta.env.VITE_RUNTIME === "wails" ? (
                <>
                  <Button type="button" className="w-full" onClick={selectSQLiteFile}>
                    Selecionar arquivo SQLite
                  </Button>
                  {sqliteFileName && <p className="text-sm text-zinc-400 break-all">{sqliteFileName}</p>}
                </>
              ) : (
                <p className="text-sm text-zinc-400">
                  A seleção de arquivos locais está disponível apenas na versão desktop.
                </p>
              )}
            </div>
          ) : (
            <>
              <Label>Host</Label>
              <Input value={connection.host} onChange={(e) => setConnection({...connection, host: e.target.value})} />

              <Label>Porta</Label>
              <Input value={connection.port} onChange={(e) => setConnection({...connection, port: e.target.value})} />

              <Label>Usuário</Label>
              <Input value={connection.username} onChange={(e) => setConnection({...connection, username: e.target.value})} />

              <Label>Senha</Label>
              <Input value={connection.password} onChange={(e) => setConnection({...connection, password: e.target.value})} type="password" />
            </>
          )}

          {error && <div className="text-red-400 text-sm bg-red-950 p-2 rounded-md">{error}</div>}

          <Button className="w-full" onClick={() => handleConnect()} disabled={loading}>{loading ? "Conectando..." : "Conectar"}</Button>
        </CardContent>
      </Card>
    </div>
  )
}