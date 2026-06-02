export function ContentSQL() {
  return (
    <div className="p-4 text-zinc-500">
      <textarea className="w-full h-64 p-2 bg-zinc-800 text-white rounded" placeholder="Escreva sua consulta SQL aqui..." disabled />
      <div className="mt-2 text-sm text-zinc-400">Resultados</div>
    </div>
  )
}