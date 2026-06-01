import { Select } from "@/components/select";
import { useConnection } from "@/contexts/ConnectionProvider";
import { useDatabase } from "@/contexts/DatabaseProvider";

export function SidebarDatabase() {
  const { activeConnection } = useConnection()
  const { activeDatabase, changeDatabase } = useDatabase()
  const databases = activeConnection?.databases || []

  return (
    <div className="p-3 font-bold border-b border-zinc-800">
      Banco de dados
      <div className="space-y-4 mt-2">
        <Select value={activeDatabase || ''} onChange={(e) => changeDatabase(e.target.value)} options={databases}/>
      </div>
    </div>
  )
}