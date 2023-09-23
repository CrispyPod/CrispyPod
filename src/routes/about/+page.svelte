<script lang="ts">
	import { onMount } from 'svelte';
	import SiteLayout from '../SiteLayout.svelte';
	import type { SiteConfig } from '$lib/models/siteConfig';
	import { get } from 'svelte/store';
	import { siteConfigS } from '$lib/stores/siteConfigStore';

	let fulldesc: string = '';

	onMount(() => {
		fulldesc = get(siteConfigS).siteFullDescription;
	});
</script>

<SiteLayout>
	<div class="flex w-full">
		<div class="grid flex-grow place-items-center">
			<h1>About</h1>
		</div>
		<div>
			{fulldesc}
		</div>
	</div>
</SiteLayout>
