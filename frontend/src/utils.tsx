import type { NotificationType } from "./components/notification"
import { SquareCheck, XSquare } from "lucide-react"

export let showFn: (message: string, type?: NotificationType) => void

export const notify = (msg: string, type: NotificationType = 'success') => {
  if (showFn) showFn(msg, type)
}

export const setNotifyHandler = (fn: typeof showFn) => showFn = fn

export function getStatus(value: boolean) {
  return value ? <SquareCheck className="text-green-600/50" /> : <XSquare className="text-red-600/60" />
}

export const COLUMNS_LABEL = ['Nome', 'Tipo', 'Tamanho', 'Nulo', 'Chave', 'Exclusivo', 'Auto Inc.','Padrão', 'Ações']