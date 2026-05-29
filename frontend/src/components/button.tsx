const variants = {
  default: "bg-zinc-800 hover:bg-zinc-700",
  success: "bg-green-600 hover:bg-green-700",
  danger: "bg-red-700 hover:bg-red-800",
  primary: "bg-blue-700 hover:bg-blue-800",
}

type ButtonProps = React.ComponentProps<"button"> & {
  variant?: "default" | "success" | "danger" | "primary"
}

export function Button({ children, className = "", variant = "default", ...props }: ButtonProps) {
  return (
    <button {...props} className={`px-3 py-2 text-sm rounded-md cursor-pointer ${variants[variant]} ${className}`}>
      {children}
    </button>
  )
}