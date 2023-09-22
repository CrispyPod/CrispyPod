<script lang="ts">
	import { onMount } from 'svelte';
	import AdminLayout from '../AdminLayout.svelte';
	import type { SiteConfig } from '$lib/models/siteConfig';

	let siteConfig: SiteConfig;

	onMount(() => {});

	async function handleFormSubmit() {}
</script>

<AdminLayout pageTitle="Site config">
	<form id="siteConfigForm" on:submit|preventDefault={handleFormSubmit}>
		<div class="form-control w-full max-w-xs">
			<label class="label" for="SiteName">
				<span class="label-text text-sm font-medium leading-6 text-gray-900">Podcast name</span>
			</label>
			<input
				id="SiteName"
				type="text"
				placeholder="Type here"
				class="input input-bordered w-full max-w-xs"
			/>

			<label class="label" for="SiteDescription">
				<span class="label-text text-sm font-medium leading-6 text-gray-900"
					>Podcast short description</span
				>
			</label>
			<input
				id="SiteDescription"
				type="text"
				placeholder="Type here"
				class="input input-bordered w-full max-w-xs"
			/>
		</div>

		<label class="label" for="SiteFullDescription">
			<span class="label-text text-sm font-medium leading-6 text-gray-900"
				>Podcast full description</span
			>
		</label>
		<textarea
			id="SiteFullDescription"
			class="textarea textarea-bordered w-full"
			placeholder="Type here"
		/>
	</form>
</AdminLayout>
