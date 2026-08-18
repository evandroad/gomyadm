type Option = {
  value: string
  text: string
}

type SelectProps = React.ComponentProps<"select"> & { options: Option[] | string[] }

export function Select({ className = "", options, ...props }: SelectProps) {
  return (
    <select {...props} className={`w-full min-w-0 h-8 bg-zinc-800 border border-zinc-700 outline-none px-1 rounded-md cursor-pointer ${className}`}>
      <option value="">Selecione</option>
      {options.map((option) => (
        typeof option === "string" ? (
          <option key={option} value={option}>
            {option}
          </option>
        ) : (
          <option key={option.value} value={option.value}>
            {option.text}
          </option>
        )
      ))}
      </select>
  )
}