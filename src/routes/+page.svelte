<script lang="ts">
	import PagedEpisodes from './PagedEpisodes.svelte';
	import SiteLayout from './SiteLayout.svelte';
	import { graphqlRequest } from '$lib/graphqlRequest';
	import { onMount } from 'svelte';
	import type { Episode } from '$lib/models/episode';

	let episodes: Array<Episode> | null = null;
	let siteName: string = '';
	let siteDescription: string = '';

	let sum = 0;
	let hasNextPage = false;
	let hasPreviousPage = false;

	onMount(async () => {
		getEpisodes();
	});

	async function getEpisodes() {
		let result = await graphqlRequest(
			null,
			`query{
				siteConfig{
    siteName
    siteDescription
    siteUrl
  },
  episodes(skip: 0,take: 10,order: {createTime:DESC},where: {episodeState:{eq:PUBLISHED}}){
    items{
      id,
      title,
	  description,
      createTime,
      thumbnailFileName,
    },
    totalCount,
    pageInfo{
      hasNextPage,
      hasPreviousPage
    },
  } 
}`
		);

		let json_resp = await result.json();
		episodes = json_resp.data.episodes.items;
		siteName = json_resp.data.siteConfig.siteName;
		siteDescription = json_resp.data.siteConfig.siteDescription;

		hasPreviousPage = json_resp.data.episodes.pageInfo.hasPreviousPage ?? false;
		hasNextPage = json_resp.data.episodes.pageInfo.hasNextPage ?? false;
		sum = json_resp.data.episodes.totalCount ?? 0;
	}
</script>

<SiteLayout>
	<div class="hero w-full h-96 bg-base-200">
		<div class="hero-content text-center">
			<div class="max-w-md">
				<h1 class="text-5xl font-bold">{siteName}</h1>
				<p class="py-6">{siteDescription}</p>
				<!-- <button class="btn btn-primary">Get Started</button> -->
			</div>
		</div>
	</div>
	{#if episodes != null}
		<div class="container mx-auto mt-5">
			Episodes
			<PagedEpisodes {episodes} {hasNextPage} {hasPreviousPage} {sum} />
		</div>
	{/if}
</SiteLayout>
