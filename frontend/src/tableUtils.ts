export function castValue(value: string, sqlType: string): any {
  const type = sqlType.toLowerCase()

  if (type.includes("int") || type.includes("decimal") || type.includes("float") || type.includes("double")) {
    return value === "" ? null : Number(value)
  }

  if (type.includes("bool") || type.includes("boolean")) {
    return value === "true"
  }

  return value
}

export function getInputType(type: string) {
  if (!type) return "text"
  type = type.toLowerCase()

  if (type.includes("int") || type.includes("decimal") || type.includes("float") || type.includes("double")) {
    return "number"
  }

  if (type.includes("date") || type.includes("timestamp")) {
    return "date"
  }

  return "text"
}