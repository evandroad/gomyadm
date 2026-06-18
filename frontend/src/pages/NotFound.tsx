import { useNavigate } from "react-router-dom"

export default function NotFound() {
  const navigate = useNavigate()

  return (
    <div className="flex flex-col items-center justify-center min-h-screen text-center px-6">

      <p className="text-[8rem] font-medium leading-none opacity-5 font-mono select-none">
        404
      </p>

      <h1 className="text-2xl font-medium -mt-8 mb-2 text-gray-900 dark:text-gray-100">
        Página não encontrada
      </h1>
      <p className="text-gray-500 max-w-sm leading-relaxed mb-8">
        A rota que você tentou acessar não existe ou foi movida.
      </p>

      <div className="flex gap-3 flex-wrap justify-center">
        <button
          onClick={() => navigate(-1)}
          className="flex items-center gap-2 px-5 py-2.5 rounded-lg border border-gray-200 text-gray-600 text-sm hover:bg-gray-50 transition-colors dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-800"
        >
          Voltar
        </button>
      </div>
    </div>
  )
}