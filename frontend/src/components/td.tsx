type Props = {
  key?: string
  className?: string
  children: React.ReactNode
}

export function Td({ className = "", children, key = "", ...props }: React.ComponentProps<"td"> & Props) {
  return (
    <td key={key} className={`px-3 py-2 whitespace-nowrap ${className}`} {...props}>
      {children}
    </td>
  )
}