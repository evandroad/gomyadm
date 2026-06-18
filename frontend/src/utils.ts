import type { NotificationType } from "./components/notification"

export let showFn: (message: string, type?: NotificationType) => void

export const notify = (msg: string, type: NotificationType = 'success') => {
  if (showFn) showFn(msg, type)
}

export const setNotifyHandler = (fn: typeof showFn) => showFn = fn

export function getStatus(value: boolean): string {
  return value ? '✅' : '❌'
}