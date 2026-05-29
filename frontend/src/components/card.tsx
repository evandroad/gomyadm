import * as React from "react"
import { useState } from "react"
import { ChevronDown } from "lucide-react"

type CardProps = React.ComponentProps<"div"> & {
  size?: "default" | "sm"
  title?: React.ReactNode
  collapsible?: boolean
  defaultOpen?: boolean
}

function Card({ className = "", size = "default", title, collapsible = false, defaultOpen = true, children, ...props }: CardProps) {
  const [open, setOpen] = useState(defaultOpen)
  return (
    <div
      data-slot="card"
      data-size={size}
      className={`group/card flex flex-col gap-4 overflow-hidden rounded-xl bg-card text-sm ring-1 ring-foreground/10 has-data-[slot=card-footer]:pb-0 has-[>img:first-child]:pt-0 data-[size=sm]:gap-3 data-[size=sm]:py-3 data-[size=sm]:has-data-[slot=card-footer]:pb-0 *:[img:first-child]:rounded-t-xl *:[img:last-child]:rounded-b-xl text-white ${className}`}
      {...props}
    >
      {(title || collapsible) && (
        <button
          type="button"
          onClick={() => collapsible && setOpen(!open)}
          className="flex w-full items-center justify-between px-4 py-3 font-medium hover:bg-white/5 transition cursor-pointer"
        >
          <span>{title}</span>
          {collapsible && <ChevronDown className={`h-4 w-4 transition-transform ${open ? "rotate-180" : ""}`}/>}
        </button>
      )}

      {open && children}
    </div>
  )
}

function CardHeader({ className = "", ...props }: React.ComponentProps<"div">) {
  return (
    <div
      data-slot="card-header"
      className={`group/card-header @container/card-header grid auto-rows-min items-start gap-1 rounded-t-xl px-4 mt-2 group-data-[size=sm]/card:px-3 has-data-[slot=card-action]:grid-cols-[1fr_auto] has-data-[slot=card-description]:grid-rows-[auto_auto] [.border-b]:pb-4 group-data-[size=sm]/card:[.border-b]:pb-3 ${className}`}
      {...props}
    />
  )
}

function CardTitle({ className = "", ...props }: React.ComponentProps<"div">) {
  return (
    <div
      data-slot="card-title"
      className={`font-heading text-base leading-snug font-medium group-data-[size=sm]/card:text-sm ${className}`}
      {...props}
    />
  )
}

function CardDescription({ className = "", ...props }: React.ComponentProps<"div">) {
  return (
    <div
      data-slot="card-description"
      className={`text-sm text-muted-foreground ${className}`}
      {...props}
    />
  )
}

function CardAction({ className = "", ...props }: React.ComponentProps<"div">) {
  return (
    <div
      data-slot="card-action"
      className={`col-start-2 row-span-2 row-start-1 self-start justify-self-end ${className}`}
      {...props}
    />
  )
}

function CardContent({ className = "", ...props }: React.ComponentProps<"div">) {
  return (
    <div
      data-slot="card-content"
      className={`p-4 group-data-[size=sm]/card:px-3 space-y-6 ${className}`}
      {...props}
    />
  )
}

function CardFooter({ className = "", ...props }: React.ComponentProps<"div">) {
  return (
    <div
      data-slot="card-footer"
      className={`flex items-center rounded-b-xl px-4 mb-2 group-data-[size=sm]/card:p-3 bg-zinc-900 justify-end space-x-2 ${className}`}
      {...props}
    />
  )
}

export { Card, CardHeader, CardFooter, CardTitle, CardAction, CardDescription, CardContent }