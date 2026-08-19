import { defineConfig, loadEnv } from 'vite'
import react from '@vitejs/plugin-react'
import tailwindcss from '@tailwindcss/vite'

export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd(), "")

  return {
    plugins: [
      react(),
      tailwindcss(),
      {
        name: "wails-runtime",
        transformIndexHtml(html) {
          if (env.VITE_RUNTIME !== "wails") {
            return html
          }
    
          return {
            html,
            tags: [
              {
                tag: "meta",
                attrs: { name: "wails-options", content: "noautoinject" },
                injectTo: "head",
              },
              {
                tag: "script",
                attrs: { src: "/wails/ipc.js" },
                injectTo: "head",
              },
              {
                tag: "script",
                attrs: { src: "/wails/runtime.js" },
                injectTo: "head",
              }
            ]
          }
        }
      }
    ],
    resolve: {
      tsconfigPaths: true,
    },
    build: {
      outDir: 'web'
    }
  }
})
