<script lang="ts">
	import { onMount, tick } from 'svelte';
	import SiteLayout from '../../SiteLayout.svelte';
	import { graphqlRequest } from '$lib/graphqlRequest';
	import type { Episode } from '$lib/models/episode';
	import WaveForm from '../../WaveForm.svelte';

	let episodeData: Episode;
	export let data;

	onMount(async () => {
		const result = await graphqlRequest(
			null,
			`{episode(id:"` +
				data.id +
				`"){id,title,createTime,description,audioFileName,audioFileDuration}}`
		);
		const jsonResp = await result.json();
		if (jsonResp.data != null) {
			episodeData = jsonResp.data.episode;
			await tick();
			// audioFile = episodeData.audioFiles[0];
		}
	});
</script>

<SiteLayout>
	{#if episodeData != null}
		<h1 class="text-2xl m-10">
			{episodeData.title}
		</h1>
		{#if episodeData.audioFileName != null}
			<div class="card lg:card-side bg-base-100 shadow-xl m-10">
				<img
					class="w-80 h-80"
					src={episodeData.thumbnailFileName ?? '/EpisodeDefaultThumbnailSquare.png'}
					alt={episodeData.title}
				/>

				<div class="card-body">
					<WaveForm fileUrl="/api/AudioFile/{episodeData.audioFileName}" />
				</div>
			</div>
		{/if}
		<p class="m-10">
			{episodeData.description}
		</p>
	{/if}
</SiteLayout>
