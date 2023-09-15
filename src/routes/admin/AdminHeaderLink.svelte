<script lang="ts">
	import { page } from '$app/stores';

	export let linkTo: string;
	export let linkName: string;
	export let largeLayout: boolean = false;
</script>

<a
	href={linkTo}
	class=" rounded-md px-3 py-2 font-medium {$page.url.pathname == linkTo
		? 'text-white bg-gray-900'
		: 'hover:bg-gray-700 hover:text-white text-gray-300'} {largeLayout ? 'text-sm' : 'block text-base'}"
	aria-current="page">{linkName}</a
>
