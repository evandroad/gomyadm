type Props = {
  tables: string[]
  selectedTable: string | null
  setSelectedTable: (table: string) => void
}

export function SidebarTables({ tables, selectedTable, setSelectedTable }: Props) {
  return (
    <div className="border-r border-zinc-800">
      <div className="overflow-auto">
        {tables.length > 0 ? (
          tables.map((table) => (
            <button key={table} onClick={() => setSelectedTable(table)} className={`w-full text-left px-3 py-2 hover:bg-zinc-950 cursor-pointer ${selectedTable === table ? "bg-zinc-900" : ""}`}>{ table }</button>
          ))
        ) : (
          <div className="p-3 text-zinc-500 text-sm">Nenhuma tabela encontrada</div>
        )}
      </div>
    </div>
  )
}