type Props = {
  key?: string
  className?: string
  children: React.ReactNode
}

export function Th({ className = "", children, key = "", ...props }: React.ComponentProps<"th"> & Props) {
  return (
    <th key={key} className={`px-3 py-2 whitespace-nowrap font-medium text-zinc-300 text-base ${className}`} {...props}>
      {children}
    </th>
  )
}