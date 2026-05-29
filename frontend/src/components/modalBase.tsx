import { useEffect, type ReactNode } from "react"

type Props = {
  open?: boolean
  onClose: () => void
  children: ReactNode
  className?: string
  zindex?: string
}

export default function ModalBase({ open = true, onClose, children, className = 'p-6 w-96', zindex = 'z-50' }: Props) {
  useEffect(() => {
    if (open) {
      document.body.style.overflow = 'hidden'
      const handleKeyDown = (e: KeyboardEvent) => { if (e.key === 'Escape') onClose() }
      window.addEventListener('keydown', handleKeyDown)
      return () => { document.body.style.overflow = ''; window.removeEventListener('keydown', handleKeyDown) }
    }
  }, [open])

  if (!open) return null

  return (
    <div className={`fixed inset-0 ${zindex} flex items-center justify-center bg-zinc-600/30`} onClick={onClose}>
      <div className={`bg-zinc-900 rounded-lg ${className}`} onClick={(e) => e.stopPropagation()}>
        {children}
      </div>
    </div>
  )
}