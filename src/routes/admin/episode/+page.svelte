<script lang="ts">
	import { onMount } from 'svelte';
	import AdminLayout from '../AdminLayout.svelte';
	import type { Episode } from '$lib/models/episode';
	import EpisodeItem from '../EpisodeItem.svelte';
	import Pager from '../../Pager.svelte';
	import { Icon, Plus } from 'svelte-hero-icons';
	import { graphqlRequest } from '$lib/graphqlRequest';
	import { get } from 'svelte/store';
	import { token } from '$lib/stores/tokenStore';
	let sum = 0;
	let perPage = 10;
	let hasNextPage = false;
	let hasPreviousPage = false;

	let episodes: Array<Episode> = [];

	let mobileMoreExpanded = false;

	async function getPagedData(newPage: number) {
		// calculate skip
		let skip = newPage * perPage;

		const tokenS = get(token);
		const result = await graphqlRequest(
			tokenS,
			`{
				episodes(pagination: {pageIndex:` +
				newPage +
				`, perPage: ` +
				perPage +
				`}){
    items{
      id
      title
	  description
      createTime
	  episodeStatus
      thumbnailFileName
    }
    totalCount
    pageInfo{
      hasNextPage
      hasPreviousPage
    }
  } 
}`
		);

		const resultJson = await result.json();

		episodes = resultJson.data.episodes.items ?? [];
		hasPreviousPage = resultJson.data.episodes.pageInfo.hasPreviousPage ?? false;
		hasNextPage = resultJson.data.episodes.pageInfo.hasNextPage ?? false;
		sum = resultJson.data.episodes.totalCount ?? 0;
	}

	onMount(async () => {
		await getPagedData(0);
		// console.log(episodeArr);
	});

	async function handlePageChange(pageIndex: number) {
		// curPage = pageIndex;
		// console.log(pageIndex);
		await getPagedData(pageIndex);
	}
</script>

<AdminLayout pageTitle="Episodes">
	<span slot="actions">
		<div class="mt-5 flex lg:ml-4 lg:mt-0">
			<a class="sm:ml-3" href="/admin/episode/new">
				<button class="btn btn-active btn-primary"> <Icon src={Plus} size="24" />New</button>
			</a>

			<!-- Dropdown -->
			<div class="relative ml-3 sm:hidden">
				<button
					type="button"
					class="inline-flex items-center rounded-md bg-white px-3 py-2 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:ring-gray-400"
					id="mobile-menu-button"
					aria-expanded="false"
					aria-haspopup="true"
					on:click={() => (mobileMoreExpanded = !mobileMoreExpanded)}
				>
					More
					<svg
						class="-mr-1 ml-1.5 h-5 w-5 text-gray-400"
						viewBox="0 0 20 20"
						fill="currentColor"
						aria-hidden="true"
					>
						<path
							fill-rule="evenodd"
							d="M5.23 7.21a.75.75 0 011.06.02L10 11.168l3.71-3.938a.75.75 0 111.08 1.04l-4.25 4.5a.75.75 0 01-1.08 0l-4.25-4.5a.75.75 0 01.02-1.06z"
							clip-rule="evenodd"
						/>
					</svg>
				</button>

				<!--
			  Dropdown menu, show/hide based on menu state.
	  
			  Entering: "transition ease-out duration-200"
				From: "transform opacity-0 scale-95"
				To: "transform opacity-100 scale-100"
			  Leaving: "transition ease-in duration-75"
				From: "transform opacity-100 scale-100"
				To: "transform opacity-0 scale-95"
			-->
				{#if mobileMoreExpanded}
					<div
						class="absolute right-0 z-10 -mr-1 mt-2 w-48 origin-top-right rounded-md bg-white py-1 shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none"
						role="menu"
						aria-orientation="vertical"
						aria-labelledby="mobile-menu-button"
						tabindex="-1"
					>
						<!-- Active: "bg-gray-100", Not Active: "" -->
						<a
							href="#"
							class="block px-4 py-2 text-sm text-gray-700"
							role="menuitem"
							tabindex="-1"
							id="mobile-menu-item-0">Edit</a
						>
						<a
							href="#"
							class="block px-4 py-2 text-sm text-gray-700"
							role="menuitem"
							tabindex="-1"
							id="mobile-menu-item-1">View</a
						>
					</div>
				{/if}
			</div>
		</div>
	</span>

	<ul role="list" class="divide-y divide-gray-100">
		{#each episodes as p, i}
			<EpisodeItem episode={p} />
		{/each}
	</ul>
	<Pager {handlePageChange} {sum} {perPage} {hasNextPage} {hasPreviousPage} />
</AdminLayout>
