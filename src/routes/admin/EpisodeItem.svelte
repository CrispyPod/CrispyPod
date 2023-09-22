<script lang="ts">
	import { goto } from '$app/navigation';
	import type { Episode } from '$lib/models/episode';
	import { Icon, Pencil, Radio } from 'svelte-hero-icons';
	export let episode: Episode;
</script>

<li class="flex justify-between gap-x-6 py-5">
	<div class="flex gap-x-4">
		<Icon src={Radio} size="48" />
		<div class="min-w-0 flex-auto">
			<p class="text-sm font-semibold leading-6 text-gray-900">{episode.title}</p>
			<p class="mt-1 truncate text-xs leading-5 text-gray-500">{episode.createTime?.toString()}</p>
		</div>
	</div>
	<div class="hidden sm:flex sm:flex-col sm:items-end">
		<a class="sm:ml-3" href="/admin/episode/{episode.id}">
			<button
				type="button"
				class="inline-flex items-center rounded-md px-3 py-2 text-sm font-semibold text-black shadow-sm hover:bg-slate-300 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
			>
				<div class="mr-1.5">
					<Icon src={Pencil} size="16" />
				</div>
				Edit
			</button>
		</a>
		<!-- <p class="text-sm leading-6 text-gray-900">{episode.episodeState}</p> -->
		<!-- <p class="mt-1 text-xs leading-5 text-gray-500">
			{episode.description}
		</p> -->
	</div>
</li>
