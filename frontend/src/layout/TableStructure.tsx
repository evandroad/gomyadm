import { useDatabase } from "@/contexts/DatabaseProvider";
import { Th } from "@/components/th";
import { Td } from "@/components/td";
import type { Values } from "@/models";
import { useSchema } from "@/contexts/SchemaProvider";

export default function TableStructure() {
  const { activeDatabase } = useDatabase()
  const { activeSchema } = useSchema()

  const COLUMNS_LABEL = ['Nome', 'Tipo', 'Nulo', 'Chave', 'Padrão', 'Extra']
  const COLUMNS = ['name', 'type', 'nullable', 'key', 'default', 'extra']

  if (!activeSchema) {
    return <div className="text-zinc-500">{ activeDatabase ? 'Carregando schema...' : 'Selecione uma base de dados' }</div>
  }

  return (
    <>
      <div>
        <table className="w-full text-sm text-left">
          <thead className="bg-zinc-900 border-b border-zinc-800">
            <tr>
              {COLUMNS_LABEL.map((column, index) => <Th key={String(index)}>{column}</Th>)}
            </tr>
          </thead>

          <tbody>
            {activeSchema?.columns.map((row: Values, index: number) => (
              <tr key={index} className="border-b border-zinc-800 hover:bg-zinc-900/50">
                {COLUMNS.map((column) => <Td key={column}>{String(row[column])}</Td>)}
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </>
  )
}