import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import { resolve, extname } from 'node:path';

export default defineConfig({
  plugins: [react()],
  base: './',
  build: {
    outDir: '../src/main/resources/com/dremio/support/diagnostics/server',
    emptyOutDir: true,
    assetsDir: 'assets',
    sourcemap: false,
    rollupOptions: {
      input: resolve(__dirname, 'index.html'),
      output: {
        entryFileNames: 'assets/app.js',
        chunkFileNames: 'assets/[name].js',
        assetFileNames: (assetInfo) => {
          const extension = extname(assetInfo.name ?? '');
          if (extension === '.css') {
            return 'assets/app.css';
          }
          return 'assets/[name][extname]';
        }
      }
    }
  }
});
