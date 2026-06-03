import { useState } from "react"
import TablePreview from "@/layout/TablePreview"
import SidebarConnection from "@/layout/SidebarConnection"
import { useConnection } from "@/contexts/ConnectionProvider"
import { SidebarDatabase } from "@/layout/SidebarDatabase"
import { SidebarTables } from "@/layout/SidebarTables"
import { Toolbar } from "@/layout/Toolbar"
import TableStructure from "@/layout/TableStructure"
import { SidebarDisconnect } from "@/layout/SidebarDisconnect"
import { useDatabase } from "@/contexts/DatabaseProvider"
import { ContentSQL } from "@/layout/ContentSQL"
import TableForm from "@/layout/TableForm"

export type View = "data" | "structure" | "form" | "sql"

export default function MainPage() {
  const { loading } = useConnection()
  const { activeDatabase } = useDatabase()
  const [selectedTable, setSelectedTable] = useState<string | null>(null)
  const [view, setView] = useState<View>("data")

  if (loading) {
    return (
      <div className="h-screen flex items-center justify-center bg-zinc-950 text-white">
        Carregando...
      </div>
    )
  }

  function renderView() {
    switch (view) {
      case "data":
      case "structure":
      case "form":
        if (!selectedTable) {
          return (
            <div className="text-zinc-500">{!activeDatabase ? 'Selecione uma base de dados' : 'Selecione uma tabela'}</div>
          )
        }
        return renderTab()
      case "sql":
        return <ContentSQL />
      default:
        return null
    }
  }

  function renderTab() {
    if (!selectedTable) return null

    const tabs = (<div className="flex w-fit border-b border-zinc-800 mb-2">
      <button onClick={() => setView("data")} className={getTabClass("data")}>Dados</button>
      <button onClick={() => setView("structure")} className={getTabClass("structure")}>Estrutura</button>
      <button onClick={() => setView("form")} className={getTabClass("form")}>Formulário</button>
    </div>)

    switch (view) {
      case "data":
        return (<>
          {tabs}
          <TablePreview table={selectedTable} />
        </>)
      case "structure":
        return (<>
          {tabs}
          <TableStructure table={selectedTable} />
        </>)
      case "form":
        return (<>
          {tabs}
          <TableForm table={selectedTable} />
        </>)
      default:
        return null
    }
  }

  function getTabClass(tab: string) {
    return `px-4 py-2 text-sm border-b-2 cursor-pointer ${
      view === tab ? "border-blue-500 text-blue-400" : "border-transparent text-zinc-400 hover:text-zinc-200"
    }`
  }

  return (
    <div className="min-h-screen flex bg-zinc-950 text-white">
      <aside className="w-64 bg-zinc-900 border-r border-zinc-800 flex flex-col">        
        <div className="p-3 font-bold border-b border-zinc-800">Gomyadm</div>
        <SidebarConnection />
        <SidebarDatabase />
        <SidebarTables selectedTable={selectedTable} setSelectedTable={setSelectedTable} setView={setView} />
        <SidebarDisconnect />
      </aside>

      <main className="flex-1 flex flex-col">
        <Toolbar view={view} setView={setView} setSelectedTable={setSelectedTable} />
        {activeDatabase && (
          <div className="flex flex-col h-full p-2">
            { renderView() }
          </div>
        )}
      </main>
    </div>
  )
}