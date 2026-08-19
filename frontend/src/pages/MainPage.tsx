import { useEffect, useState } from "react"
import SectionTablePreview from "@/layout/section/SectionTablePreview"
import SidebarConnection from "@/layout/sidebar/SidebarConnection"
import { SidebarDatabase } from "@/layout/sidebar/SidebarDatabase"
import { SidebarTables } from "@/layout/sidebar/SidebarTables"
import { Toolbar } from "@/layout/Toolbar"
import SectionTableSchema from "@/layout/section/SectionTableSchema"
import { SidebarDisconnect } from "@/layout/sidebar/SidebarDisconnect"
import { SectionContentSQL } from "@/layout/section/SectionContentSQL"
import SectionFormColumn from "@/layout/section/SectionFormColumn"
import SectionFormTable from "@/layout/section/SectionFormTable"
import { Button } from "@/components/button"
import { Plus, RefreshCcw } from "lucide-react"
import { ModalFormItem } from "@/layout/modal/ModalFormItem"
import { repositories } from '@/repositories'
import { useDatabase } from "@/contexts/DatabaseContext";
import { useTable } from "@/contexts/TableContext";
import { useConnection } from "@/contexts/ConectionContext";

export type View = "data" | "struct" | "formColumn" | "formTable" | "sql"

export default function MainPage() {
  const { loading } = useConnection()
  const { activeDatabase } = useDatabase()
  const { load, loadTableData } = useTable()
  const [selectedTable, setSelectedTable] = useState<string | null>(null)
  const [view, setView] = useState<View>("data")
  const [sidebarWidth, setSidebarWidth] = useState(256)
  const [openFormItem, setOpenFormItem] = useState<boolean>(false)
  const [version, setVersion] = useState<string>('')

  useEffect(() => {load(selectedTable)}, [selectedTable, load])
  useEffect(() => {loadTableData(selectedTable)} ,[selectedTable, loadTableData])

  useEffect(() => {
    async function fetchVersion() {
      const res = await repositories.app.version()
      if (res.ok) setVersion(res.data)
    }

    fetchVersion()
  }, [])

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
      case "struct":
      case "formColumn":
        if (!selectedTable) {
          return (
            <div className="text-zinc-500">{!activeDatabase ? 'Selecione uma base de dados' : 'Selecione uma tabela'}</div>
          )
        }
        return renderTab()
      case "sql":
        return <SectionContentSQL />
      case "formTable":
        return <SectionFormTable />
      default:
        return null
    }
  }

  function renderTab() {
    if (!selectedTable) return null

    const tabs = (
      <div className="flex flex-row justify-between items-center">
        <div className="flex w-fit border-b border-zinc-800 mb-2">
          <button onClick={() => setView("data")} className={getTabClass("data")}>Dados</button>
          <button onClick={() => setView("struct")} className={getTabClass("structure")}>Estrutura</button>
          {(view == 'struct' || view == 'formColumn') && 
            <button onClick={() => setView("formColumn")} className={getTabClass("formColumn")}>Nova Coluna</button>
          }
        </div>
        <div className="space-x-1">
          {view == 'data' &&
            <>
              <Button variant="success" onClick={() => setOpenFormItem(true)}><Plus size={16}/></Button>
              <Button onClick={() => loadTableData(selectedTable)}><RefreshCcw size={16}/></Button>
            </>
          }
          {view == 'struct' &&
            <>
              <Button variant="success" onClick={() => setView("struct")}><Plus size={16}/></Button>
              <Button><RefreshCcw size={16}/></Button>
            </>
          }
        </div>
      </div>
    )

    let content = <div>Conteúdo não encontrado</div>

    switch (view) {
      case "data":
        content = <SectionTablePreview />
        break
      case "struct":
        content = <SectionTableSchema />
        break
      case "formColumn":
        content = <SectionFormColumn />
        break
    }

    return (<>
      {tabs}
      {content}
    </>)
  }

  function getTabClass(tab: string) {
    return `px-4 py-2 text-sm border-b-2 cursor-pointer ${
      view === tab ? "border-blue-500 text-blue-400" : "border-transparent text-zinc-400 hover:text-zinc-200"
    }`
  }

  const startResize = (e: React.MouseEvent) => {
    const startX = e.clientX
    const startWidth = sidebarWidth

    const onMouseMove = (e: MouseEvent) => {
      const newWidth = startWidth + (e.clientX - startX)
      setSidebarWidth(Math.min(Math.max(newWidth, 150), 800))
    }

    const onMouseUp = () => {
      document.removeEventListener("mousemove", onMouseMove)
      document.removeEventListener("mouseup", onMouseUp)
    }

    document.addEventListener("mousemove", onMouseMove)
    document.addEventListener("mouseup", onMouseUp)
  }

  return (
    <div className="min-h-screen flex bg-zinc-950 text-white">
      <ModalFormItem open={openFormItem} onClose={() => setOpenFormItem(false)} data={null} />

      <aside className="bg-zinc-900 border-r border-zinc-800 flex flex-col relative" style={{ width: sidebarWidth }}>
        <div className="py-4 px-3 font-semibold text-lg border-b border-zinc-800">
          Gomyadm
          <p className="text-sm text-zinc-400 font-light">{version}</p>
        </div>
        <SidebarConnection />
        <SidebarDatabase />
        <SidebarTables selectedTable={selectedTable} setSelectedTable={setSelectedTable} setView={setView} />
        <SidebarDisconnect />
        <div
          onMouseDown={startResize}
          className="absolute top-0 right-0 w-1 h-full cursor-col-resize hover:bg-blue-500"
        />
      </aside>

      <main className="flex-1 flex flex-col">
        <Toolbar view={view} table={selectedTable} setView={setView} setSelectedTable={setSelectedTable} />
        {activeDatabase && (
          <div className="flex flex-col h-full px-2">
            { renderView() }
          </div>
        )}
      </main>
    </div>
  )
}