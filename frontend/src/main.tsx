import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.tsx'
import { BrowserRouter } from 'react-router-dom'
import { ConnectionProvider } from './contexts/ConnectionProvider.tsx'
import { ConnectionsProvider } from './contexts/ConnectionsProvider.tsx'
import { DatabaseProvider } from './contexts/DatabaseProvider.tsx'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <BrowserRouter>
      <ConnectionProvider>
        <ConnectionsProvider>
          <DatabaseProvider>
            <App />
          </DatabaseProvider>
        </ConnectionsProvider>
      </ConnectionProvider>
    </BrowserRouter>
  </StrictMode>,
)
