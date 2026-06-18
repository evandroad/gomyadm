import { BrowserRouter as Router, Navigate, Route, Routes } from "react-router-dom"
import DatabaseConnectionPage from "./pages/DatabaseConnectionPage"
import MainPage from "./pages/MainPage"
import ConnectionGuard from "./guard/ConnectionGuard"
import GuestGuard from "./guard/GuestGuard"
import Notification from "./components/notification"
import NotFound from "./pages/NotFound"

export default function App() {
  return (
    <Router>
      <Notification />
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

        <Route path="*" element={<NotFound />} />
      </Routes>
    </Router>
  )
}