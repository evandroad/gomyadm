type Props = {
  className?: string
  children: React.ReactNode
}

export function Th({ className = "", children, ...props }: React.ComponentProps<"th"> & Props) {
  return (
    <th className={`px-3 py-2 whitespace-nowrap font-medium text-zinc-300 text-base ${className}`} {...props}>
      {children}
    </th>
  )
}