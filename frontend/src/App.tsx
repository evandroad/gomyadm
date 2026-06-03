import DatabaseConnectionPage from "./pages/DatabaseConnectionPage"
import MainPage from "./pages/MainPage"
import { Navigate, Route, Routes } from "react-router-dom"
import ConnectionGuard from "./ConnectionGuard"
import GuestGuard from "./GuestGuard"

export default function App() {
  return (
    <Routes>
      {/* rotas públicas somente sem conexão */}
      <Route element={<GuestGuard />}>
        <Route path="/connect" element={<DatabaseConnectionPage />} />
      </Route>

      {/* rotas protegidas */}
      <Route element={<ConnectionGuard />}>
        <Route path="/" element={<Navigate to="/app" replace />} />
        <Route path="/app" element={<MainPage />} />
      </Route>
    </Routes>
  )
}