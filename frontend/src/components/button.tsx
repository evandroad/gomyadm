const variants = {
  default: "bg-zinc-800 hover:bg-zinc-700",
  success: "bg-green-600 hover:bg-green-700",
  danger: "bg-red-700 hover:bg-red-800",
  primary: "bg-blue-700 hover:bg-blue-800",
}

type ButtonProps = React.ComponentProps<"button"> & {
  variant?: "default" | "success" | "danger" | "primary"
  sm?: boolean
  bg?: string
}

export function Button({ children, className = "", variant = "default", sm = false, bg = "", ...props }: ButtonProps) {
  return (
    <button {...props} className={`rounded-md cursor-pointer ${bg ? bg : variants[variant]} ${sm ? 'p-1 text-xs' : 'p-2 text-sm'} ${className}`}>
      {children}
    </button>
  )
}