import DatabaseConnectionPage from "./DatabaseConnectionPage"
import MainPage from "./MainPage"
import { Navigate, Route, Routes } from "react-router-dom"
import ConnectionGuard from "./ConnectionGuard"

export default function App() {
   return (
    <Routes>
      <Route path="/" element={<Navigate to="/app" replace />} />
      <Route path="/connect" element={<DatabaseConnectionPage />} />

      <Route
        path="/app"
        element={
          <ConnectionGuard>
            <MainPage />
          </ConnectionGuard>
        }
      />
    </Routes>
  )
}