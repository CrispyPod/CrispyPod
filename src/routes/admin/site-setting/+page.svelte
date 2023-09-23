<script lang="ts">
	import { onMount } from 'svelte';
	import AdminLayout from '../AdminLayout.svelte';
	import type { SiteConfig } from '$lib/models/siteConfig';
	import { siteConfigS } from '$lib/stores/siteConfigStore';
	import { get } from 'svelte/store';
	import { graphqlRequest } from '$lib/graphqlRequest';
	import { token } from '$lib/stores/tokenStore';
	import { goto } from '$app/navigation';

	let siteConfig: SiteConfig;
	let errMessage: string | null = null;

	onMount(async () => {
		siteConfig = get(siteConfigS);
		if (siteConfig == null) {
			const result = await graphqlRequest(
				null,
				`{
  siteConfig{
    siteUrl
    siteName
    siteDescription
    siteFullDescription
  }
}`
			);
			const jsonResp = await result.json();
			if (jsonResp.data != null) {
				siteConfigS.set(jsonResp.data.siteConfig);
				siteConfig = jsonResp.data.siteConfig;
			}
		}
	});

	async function handleFormSubmit(e: SubmitEvent) {
		const form = document.querySelector('#siteConfigForm');
		const formData = new FormData(form as HTMLFormElement);
		const tokenS = get(token);
		const result = await graphqlRequest(
			tokenS,
			`mutation{
  modifySiteConfig(input:{siteName:"` +
				formData.get('SiteName') +
				`",siteDescription:"` +
				formData.get('SiteDescription') +
				`",siteFullDescription:"` +
				formData.get('SiteFullDescription') +
				`",siteUrl:"` +
				formData.get('SiteUrl') +
				`"}){
    siteUrl
    siteName
    siteDescription
    siteFullDescription
  }
}`
		);

		var resultJson = await result.json();
		if (resultJson.data != null) {
			siteConfigS.set(resultJson.data.modifySiteConfig);
			goto('/admin');
		} else {
			errMessage = resultJson.errors[0].message;
		}
	}
</script>

<AdminLayout pageTitle="Site setting">
	<form id="siteConfigForm" on:submit|preventDefault={handleFormSubmit}>
		<div class="form-control w-full max-w-xs">
			<label class="label" for="SiteName">
				<span class="label-text text-sm font-medium leading-6 text-gray-900">Podcast name</span>
			</label>
			<input
				id="SiteName"
				name="SiteName"
				type="text"
				placeholder="Type here"
				value={siteConfig == null ? '' : siteConfig.siteName}
				class="input input-bordered w-full max-w-xs"
			/>

			<label class="label" for="SiteDescription">
				<span class="label-text text-sm font-medium leading-6 text-gray-900"
					>Podcast short description</span
				>
			</label>
			<input
				id="SiteDescription"
				name="SiteDescription"
				type="text"
				placeholder="Type here"
				value={siteConfig == null ? '' : siteConfig.siteDescription}
				class="input input-bordered w-full max-w-xs"
			/>

			<label class="label" for="SiteUrl">
				<span class="label-text text-sm font-medium leading-6 text-gray-900"
					>Podcast short description</span
				>
			</label>
			<input
				id="SiteUrl"
				name="SiteUrl"
				type="url"
				placeholder="Type here"
				value={siteConfig == null ? '' : siteConfig.siteUrl}
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
			name="SiteFullDescription"
			value={siteConfig == null ? '' : siteConfig.siteFullDescription}
			class="textarea textarea-bordered w-full"
			placeholder="Type here"
		/>

		<div class="mt-6 flex items-center justify-end gap-x-6">
			{#if errMessage != null}
				<div class="mr-auto">
					<div
						class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded relative"
						role="alert"
					>
						<strong class="font-bold">{errMessage}</strong>
					</div>
				</div>
			{/if}
			<a href="/admin" type="button" class="text-sm font-semibold leading-6 text-gray-900">Cancel</a
			>
			<button
				type="submit"
				class="rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
				>Save</button
			>
		</div>
	</form>
</AdminLayout>
