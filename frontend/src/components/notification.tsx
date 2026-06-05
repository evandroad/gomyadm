import { setNotifyHandler } from '@/utils'
import { useEffect, useRef, useState } from 'react'

export type NotificationType = 'success' | 'error' | 'warning'

export default function Notification() {
  const [message, setMessage] = useState<string>('')
  const [type, setType]  = useState<NotificationType>('success')
  const timeoutRef = useRef<any | null>(null)

  const closeNotification = () => {
    setMessage("")
    if (timeoutRef.current) {
      clearTimeout(timeoutRef.current)
      timeoutRef.current = null
    }
  }

  useEffect(() => {
    setNotifyHandler((msg: string, tp: NotificationType = 'success') => {
      setMessage(msg)
      setType(tp)
      if (timeoutRef.current) clearTimeout(timeoutRef.current)
      timeoutRef.current = setTimeout(closeNotification, 3000)
    })
  }, [])

  if (!message) return null

  const typeStyles: Record<NotificationType, string> = {
    success: "bg-green-500",
    error: "bg-red-500",
    warning: "bg-yellow-500 text-black",
  }

  return(
    <>
      <div className={`${typeStyles[type]} fixed z-9999 top-10 left-1/2 -translate-x-1/2 -translate-y-1/2 text-white px-4 py-2 rounded-md shadow-xl text-center animate-fade-in flex items-center justify-between gap-4 min-w-75`}>
        <span className='flex-1 text-left'>{message}</span>
        <button onClick={closeNotification} className="text-white hover:text-zinc-200 font-bold text-lg leading-none cursor-pointer">
          x
        </button>
      </div>
    </>
  )
}