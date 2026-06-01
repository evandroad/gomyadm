import { Button } from "@/components/button"
import { useDatabase } from "@/contexts/DatabaseProvider"

type Props = {
  selectedTable: string | null
  view: "data" | "structure"
  setView: (view: "data" | "structure") => void
}

export function Toolbar({ selectedTable, view, setView }: Props) {
  const { activeDatabase } = useDatabase()

  return (
    <div className="flex items-center justify-between w-full h-16 px-4 bg-zinc-900 text-white">
      <h1 className="text-lg font-semibold">{(selectedTable && activeDatabase) ? selectedTable : "Visão geral"}</h1>
        {selectedTable && activeDatabase && (
          <div className="space-x-2">
            <Button onClick={() => setView("data")} bg={view === "data" ? "bg-zinc-800" : "bg-zinc-900"}>Dados</Button>
            <Button onClick={() => setView("structure")} bg={view === "structure" ? "bg-zinc-800" : "bg-zinc-900"}>Estrutura</Button>
          </div>
        )}
    </div>
  )
}