const variants = {
  default: "bg-zinc-800 text-gray-300 hover:bg-zinc-700",
  success: "bg-green-600 hover:bg-green-700 dark:bg-green-600/20 dark:text-green-400 dark:hover:bg-green-600/30",
  danger: "bg-red-700 hover:bg-red-800 dark:bg-red-600/20 dark:text-red-400 dark:hover:bg-red-600/30",
  primary: "bg-blue-700 hover:bg-blue-800 dark:bg-blue-600/20 dark:text-blue-400 dark:hover:bg-blue-600/30",
}

type variantsType = "default" | "success" | "danger" | "primary"

type ButtonProps = React.ComponentProps<"button"> & {
  variant?: variantsType 
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