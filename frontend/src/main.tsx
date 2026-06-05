import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.tsx'
import { ConnectionProvider } from './contexts/ConnectionProvider.tsx'
import { ConnectionsProvider } from './contexts/ConnectionsProvider.tsx'
import { DatabaseProvider } from './contexts/DatabaseProvider.tsx'
import { SchemaProvider } from './contexts/SchemaProvider.tsx'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <ConnectionProvider>
      <ConnectionsProvider>
        <DatabaseProvider>
          <SchemaProvider>
            <App />
          </SchemaProvider>
        </DatabaseProvider>
      </ConnectionsProvider>
    </ConnectionProvider>
  </StrictMode>,
)
