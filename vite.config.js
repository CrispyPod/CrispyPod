import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vitest/config';

export default defineConfig({
	plugins: [sveltekit()],
	test: {
		include: ['src/**/*.{test,spec}.{js,ts}']
	},
	optimizeDeps: {
		exclude: ['bytemd']
	},
	server: {
		proxy: {
			'/graphql': {
				target: 'http://localhost:8080/graphql',
				changeOrigin: true,
				rewrite: (path) => path.replace(/^\/graphql/, '')
			},
			'/api': {
				target: 'http://localhost:8080/api',
				changeOrigin: true,
				rewrite: (path) => path.replace(/^\/api/, '')
			},
			'/rss': {
				target: 'http://localhost:8080/rss',
				changeOrigin: true,
				rewrite: (path) => path.replace(/^\/rss/, '')
			}
		}
	}
});
