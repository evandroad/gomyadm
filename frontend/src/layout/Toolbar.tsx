export function Toolbar({ children }: { children?: React.ReactNode }) {
  return (
    <div className="flex items-center justify-between w-full h-16 px-4 bg-zinc-900 text-white">
      {children}
    </div>
  )
}