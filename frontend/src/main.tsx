import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.tsx'
import { ConnectionProvider } from './contexts/ConnectionProvider.tsx'
import { ConnectionsProvider } from './providers/ConnectionsProvider.tsx';
import { DatabaseProvider } from './providers/DatabaseProvider.tsx';
import { TableProvider } from './providers/TableProvider.tsx';

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <ConnectionProvider>
      <ConnectionsProvider>
        <DatabaseProvider>
          <TableProvider>
            <App />
          </TableProvider>
        </DatabaseProvider>
      </ConnectionsProvider>
    </ConnectionProvider>
  </StrictMode>,
)

if (window.matchMedia("(prefers-color-scheme: dark)").matches) {
  document.documentElement.classList.add("dark")
}