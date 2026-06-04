import { Button } from "@/components/button"
import { useDatabase } from "@/contexts/DatabaseProvider"
import type { View } from "@/pages/MainPage"

type Props = {
  view: View
  setView: (view: View) => void
  setSelectedTable: (table: string | null) => void
}

export function Toolbar({ view, setView, setSelectedTable }: Props) {
  const { activeDatabase } = useDatabase()

  return (
    <div className="flex items-center justify-between w-full h-16 px-4 bg-zinc-900 text-white">
      <h1 className="text-lg font-semibold">{activeDatabase ?? "Visão geral"}</h1>
        {activeDatabase && (
          <div className="space-x-2">
            <Button onClick={() => {setSelectedTable(null); setView("sql")}} bg={view === "sql" ? "bg-zinc-800" : "outline"}>SQL</Button>
          </div>
        )}
    </div>
  )
}