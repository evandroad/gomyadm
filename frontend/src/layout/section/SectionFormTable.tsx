import { Input } from "@/components/input";
import { Label } from "@/components/label";

export default function SectionFormTable() {
  return (
    <div className="p-2 w-100 space-y-4">
      <Label>Name</Label>
      <Input value={""} /*onChange={(e) => setColumn({ ...column, name: e.target.value })}*/ />
    </div>
  )
}