// @ts-check
import { defineConfig, envField } from 'astro/config';

import tailwindcss from '@tailwindcss/vite';
import svgr from 'vite-plugin-svgr';

import react from '@astrojs/react';

// https://astro.build/config
export default defineConfig({
  vite: {
    plugins: [
      tailwindcss(),
      svgr({
        include: '**/*.svg?react',
      }),
    ],
  },
  env: {
    schema: {
      BACKEND_URL: envField.string({
        access: 'public',
        context: 'client',
        url: true,
        default: 'http://localhost:8080',
      }),
    },
  },

  integrations: [react()],
});
