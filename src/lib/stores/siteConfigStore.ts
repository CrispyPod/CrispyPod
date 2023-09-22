import type { SiteConfig } from '$lib/models/siteConfig';
import { writable } from 'svelte/store';

function siteConfigStore() {
	const { subscribe, set } = writable<SiteConfig>();

	return {
		subscribe,
		set
	};
}

export const siteConfigS = siteConfigStore();
