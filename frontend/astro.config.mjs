// @ts-check
import { defineConfig } from 'astro/config';

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

  integrations: [react()],
});
