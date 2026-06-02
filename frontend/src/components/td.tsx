type Props = {
  className?: string
  children: React.ReactNode
}

export function Td({ className = "", children, ...props }: React.ComponentProps<"td"> & Props) {
  return (
    <td className={`px-3 py-2 whitespace-nowrap ${className}`} {...props}>
      {children}
    </td>
  )
}