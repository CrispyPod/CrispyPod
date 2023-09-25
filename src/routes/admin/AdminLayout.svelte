<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import { token } from '$lib/stores/tokenStore';
	import { onMount } from 'svelte';
	import { get } from 'svelte/store';
	import AdminHeaderLink from './AdminHeaderLink.svelte';
	import { BreadCombItem } from '$lib/models/breadCombItems';
	// import { Icon, UserCircle } from 'svelte-hero-icons';
	import { graphqlRequest } from '$lib/graphqlRequest';
	// import type { SiteConfig } from '$lib/models/siteConfig';
	// import { siteConfigS } from '$lib/stores/siteConfigStore';

	// let siteConfig: SiteConfig;

	let showAvatarOption = false;
	function handleSignOut() {
		token.set(null);
	}

	let breadCombItems: Array<BreadCombItem> = [];

	onMount(async () => {
		// await siteConfigS.init();
		// siteConfig = get(siteConfigS);
		const tokenS = get(token);

		if ($page.url.pathname != '/admin/signin' && tokenS != null) {
			let result = await graphqlRequest(tokenS, `query{me{id,userName}}`);
			let jsonResp = await result.json();
			if (jsonResp.data == null || jsonResp.data.me == null) {
				token.set(null);
			}
		}

		if ($page.url.pathname == '/admin/signin') {
			if (tokenS != null) {
				goto('/admin');
			}
		} else {
			if (tokenS == null) {
				goto('/admin/signin');
			}
		}

		if ($page.url.pathname.replace('/admin', '').length > 0) {
			let builtLink = '';
			breadCombItems = $page.url.pathname
				.split('/')
				.filter((x) => x.length > 0)
				.map((x) => {
					builtLink += '/' + x;
					const comb = new BreadCombItem(builtLink, x.slice(0, 1).toUpperCase() + x.slice(1));
					return comb;
				});
		}
	});

	export function loginCheck() {
		const tokenStore = get(token);
		if (tokenStore == null) {
			goto('/admin/signin');
		}
	}

	export let pageTitle: String;
</script>

<svelte:head>
	<title>
		{pageTitle} - CrispyPod
	</title>
</svelte:head>

<div class="min-h-full">
	<nav class="bg-gray-800">
		<div class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
			<div class="flex h-16 items-center justify-between">
				<div class="flex items-center">
					<div class="flex-shrink-0 text-white font-bold">
						CrispyPod
						<!-- <img
							class="h-8 w-8"
							src="https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=500"
							alt="Your Company"
						/> -->
					</div>
					<div class="hidden md:block">
						<div class="ml-10 flex items-baseline space-x-4">
							<!-- Current: "bg-gray-900 text-white", Default: "text-gray-300 hover:bg-gray-700 hover:text-white" -->
							<!-- <AdminHeaderLink linkName="Site Setting" linkTo="/admin/site-config" largeLayout={true} /> -->
							<AdminHeaderLink linkName="Dashboard" linkTo="/admin" largeLayout={true} />
							<AdminHeaderLink linkName="Episodes" linkTo="/admin/episode" largeLayout={true} />
							<!-- <a
								href="/admin"
								class=" rounded-md px-3 py-2 text-sm font-medium {$page.url.pathname == '/admin'
									? 'text-white bg-gray-900'
									: 'hover:bg-gray-700 hover:text-white text-gray-300'}"
								aria-current="page">Dashboard</a
							> -->
							<!-- <a
								href="/admin/episode"
								class=" rounded-md px-3 py-2 text-sm font-medium {$page.url.pathname == '/admin/episode'
									? 'text-white'
									: 'hover:bg-gray-700 hover:text-white text-gray-300'}">Episodes</a
							> -->
							<!-- <a
								href="#"
								class="text-gray-300 hover:bg-gray-700 hover:text-white rounded-md px-3 py-2  font-medium"
								>Projects</a
							>
							<a
								href="#"
								class="text-gray-300 hover:bg-gray-700 hover:text-white rounded-md px-3 py-2 text-sm font-medium"
								>Calendar</a
							>
							<a
								href="#"
								class="text-gray-300 hover:bg-gray-700 hover:text-white rounded-md px-3 py-2 text-sm font-medium"
								>Reports</a
							> -->
						</div>
					</div>
				</div>
				<div class="hidden md:block">
					<div class="ml-4 flex items-center md:ml-6">
						<button
							type="button"
							class="rounded-full bg-gray-800 p-1 text-gray-400 hover:text-white focus:outline-none focus:ring-2 focus:ring-white focus:ring-offset-2 focus:ring-offset-gray-800"
						>
							<span class="sr-only">View notifications</span>
							<svg
								class="h-6 w-6"
								fill="none"
								viewBox="0 0 24 24"
								stroke-width="1.5"
								stroke="currentColor"
								aria-hidden="true"
							>
								<path
									stroke-linecap="round"
									stroke-linejoin="round"
									d="M14.857 17.082a23.848 23.848 0 005.454-1.31A8.967 8.967 0 0118 9.75v-.7V9A6 6 0 006 9v.75a8.967 8.967 0 01-2.312 6.022c1.733.64 3.56 1.085 5.455 1.31m5.714 0a24.255 24.255 0 01-5.714 0m5.714 0a3 3 0 11-5.714 0"
								/>
							</svg>
						</button>

						<!-- Profile dropdown -->
						<div class="relative ml-3">
							<div>
								<button
									type="button"
									class="flex max-w-xs items-center rounded-full bg-gray-800 text-sm focus:outline-none focus:ring-2 focus:ring-white focus:ring-offset-2 focus:ring-offset-gray-800"
									id="user-menu-button"
									aria-expanded="false"
									aria-haspopup="true"
									on:click={() => (showAvatarOption = !showAvatarOption)}
								>
									<span class="sr-only">Open user menu</span>

									<!-- <Icon src={UserCircle} size="32"/> -->
									<svg
										xmlns="http://www.w3.org/2000/svg"
										fill="none"
										viewBox="0 0 24 24"
										stroke-width="1.5"
										stroke="white"
										class="w-8 h-8"
									>
										<path
											stroke-linecap="round"
											stroke-linejoin="round"
											d="M17.982 18.725A7.488 7.488 0 0012 15.75a7.488 7.488 0 00-5.982 2.975m11.963 0a9 9 0 10-11.963 0m11.963 0A8.966 8.966 0 0112 21a8.966 8.966 0 01-5.982-2.275M15 9.75a3 3 0 11-6 0 3 3 0 016 0z"
										/>
									</svg>

									<!-- <img
										class="h-8 w-8 rounded-full"
										src="https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80"
										alt=""
									/> -->
								</button>
							</div>

							<!--
                  Dropdown menu, show/hide based on menu state.
  
                  Entering: "transition ease-out duration-100"
                    From: "transform opacity-0 scale-95"
                    To: "transform opacity-100 scale-100"
                  Leaving: "transition ease-in duration-75"
                    From: "transform opacity-100 scale-100"
                    To: "transform opacity-0 scale-95"
                -->
							{#if showAvatarOption}
								<div
									class="absolute right-0 z-10 mt-2 w-48 origin-top-right rounded-md bg-white py-1 shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none"
									role="menu"
									aria-orientation="vertical"
									aria-labelledby="user-menu-button"
									tabindex="-1"
								>
									<!-- Active: "bg-gray-100", Not Active: "" -->
									<a
										href="/admin/profile"
										class="block px-4 py-2 text-sm text-gray-700"
										role="menuitem"
										tabindex="-1"
										id="user-menu-item-0">Your Profile</a
									>
									<a
										href="/admin/site-setting"
										class="block px-4 py-2 text-sm text-gray-700"
										role="menuitem"
										tabindex="-1"
										id="user-menu-item-1">Site Setting</a
									>
									<a
										href="#"
										class="block px-4 py-2 text-sm text-gray-700"
										role="menuitem"
										tabindex="-1"
										id="user-menu-item-2"
										on:click={handleSignOut}>Sign out</a
									>
								</div>
							{/if}
						</div>
					</div>
				</div>
				<div class="-mr-2 flex md:hidden">
					<!-- Mobile menu button -->
					<button
						type="button"
						class="inline-flex items-center justify-center rounded-md bg-gray-800 p-2 text-gray-400 hover:bg-gray-700 hover:text-white focus:outline-none focus:ring-2 focus:ring-white focus:ring-offset-2 focus:ring-offset-gray-800"
						aria-controls="mobile-menu"
						aria-expanded="false"
						on:click={() => (showAvatarOption = !showAvatarOption)}
					>
						<span class="sr-only">Open main menu</span>
						<!-- Menu open: "hidden", Menu closed: "block" -->
						<svg
							class="block h-6 w-6"
							fill="none"
							viewBox="0 0 24 24"
							stroke-width="1.5"
							stroke="currentColor"
							aria-hidden="true"
						>
							<path
								stroke-linecap="round"
								stroke-linejoin="round"
								d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5"
							/>
						</svg>
						<!-- Menu open: "block", Menu closed: "hidden" -->
						<svg
							class="hidden h-6 w-6"
							fill="none"
							viewBox="0 0 24 24"
							stroke-width="1.5"
							stroke="currentColor"
							aria-hidden="true"
						>
							<path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
						</svg>
					</button>
				</div>
			</div>
		</div>

		<!-- Mobile menu, show/hide based on menu state. -->
		{#if showAvatarOption}
			<div class="md:hidden" id="mobile-menu">
				<div class="space-y-1 px-2 pb-3 pt-2 sm:px-3">
					<!-- Current: "bg-gray-900 text-white", Default: "text-gray-300 hover:bg-gray-700 hover:text-white" -->

					<AdminHeaderLink linkName="Dashboard" linkTo="/admin" />
					<AdminHeaderLink linkName="Episodes" linkTo="/admin/episode" />
					<!-- <a
						href="/admin"
						class="bg-gray-900 text-white block rounded-md px-3 py-2 text-base font-medium"
						aria-current="page">Dashboard</a
					>
					<a
						href="/admin/episode"
						class="text-gray-300 hover:bg-gray-700 hover:text-white block rounded-md px-3 py-2 text-base font-medium"
						>Episodes</a
					> -->

					<!-- <a
					href="#"
					class="text-gray-300 hover:bg-gray-700 hover:text-white block rounded-md px-3 py-2 text-base font-medium"
					>Projects</a
				>
				<a
					href="#"
					class="text-gray-300 hover:bg-gray-700 hover:text-white block rounded-md px-3 py-2 text-base font-medium"
					>Calendar</a
				>
				<a
					href="#"
					class="text-gray-300 hover:bg-gray-700 hover:text-white block rounded-md px-3 py-2 text-base font-medium"
					>Reports</a
				> -->
				</div>
				<div class="border-t border-gray-700 pb-3 pt-4">
					<div class="flex items-center px-5">
						<div class="flex-shrink-0">
							<svg
								class="block h-6 w-6"
								fill="none"
								viewBox="0 0 24 24"
								stroke-width="1.5"
								stroke="currentColor"
								aria-hidden="true"
							>
								<path
									stroke-linecap="round"
									stroke-linejoin="round"
									d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5"
								/>
							</svg>
							<!-- <Icon src={UserCircle} /> -->
							<!-- <img
								class="h-10 w-10 rounded-full"
								src="https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80"
								alt=""
							/> -->
						</div>
						<div class="ml-3">
							<div class="text-base font-medium leading-none text-white">Tom Cook</div>
							<div class="text-sm font-medium leading-none text-gray-400">tom@example.com</div>
						</div>
						<button
							type="button"
							class="ml-auto flex-shrink-0 rounded-full bg-gray-800 p-1 text-gray-400 hover:text-white focus:outline-none focus:ring-2 focus:ring-white focus:ring-offset-2 focus:ring-offset-gray-800"
						>
							<span class="sr-only">View notifications</span>
							<svg
								class="h-6 w-6"
								fill="none"
								viewBox="0 0 24 24"
								stroke-width="1.5"
								stroke="currentColor"
								aria-hidden="true"
							>
								<path
									stroke-linecap="round"
									stroke-linejoin="round"
									d="M14.857 17.082a23.848 23.848 0 005.454-1.31A8.967 8.967 0 0118 9.75v-.7V9A6 6 0 006 9v.75a8.967 8.967 0 01-2.312 6.022c1.733.64 3.56 1.085 5.455 1.31m5.714 0a24.255 24.255 0 01-5.714 0m5.714 0a3 3 0 11-5.714 0"
								/>
							</svg>
						</button>
					</div>

					<div class="mt-3 space-y-1 px-2">
						<a
							href="/admin/profile"
							class="block rounded-md px-3 py-2 text-base font-medium text-gray-400 hover:bg-gray-700 hover:text-white"
							>Your Profile</a
						>
						<a
							href="/admin/site-setting"
							class="block rounded-md px-3 py-2 text-base font-medium text-gray-400 hover:bg-gray-700 hover:text-white"
							>Site-Setting</a
						>
						<a
							href="#"
							class="block rounded-md px-3 py-2 text-base font-medium text-gray-400 hover:bg-gray-700 hover:text-white"
							on:click={handleSignOut}>Sign out</a
						>
					</div>
				</div>
			</div>
		{/if}
	</nav>

	<header class="bg-white shadow">
		<div class="mx-auto max-w-7xl px-4 py-6 sm:px-6 lg:px-8">
			<!-- <h1 class="text-3xl font-bold tracking-tight text-gray-900">{pageTitle}</h1> -->

			<nav class="w-full rounded-md">
				<ol class="list-reset flex">
					{#each breadCombItems as b, i}
						{#if i != breadCombItems.length - 1}
							<li>
								<a
									href={b.link}
									class="text-primary transition duration-150 ease-in-out hover:text-primary-600 focus:text-primary-600 active:text-primary-700 dark:text-primary-400 dark:hover:text-primary-500 dark:focus:text-primary-500 dark:active:text-primary-600"
									>{b.name}</a
								>
							</li>

							<li>
								<span class="mx-2 text-neutral-500 dark:text-neutral-400">/</span>
							</li>
						{:else}
							<li class="text-neutral-500 dark:text-neutral-400">{b.name}</li>{/if}
					{/each}
					<!-- <li>
						<a
							href="#"
							class="text-primary transition duration-150 ease-in-out hover:text-primary-600 focus:text-primary-600 active:text-primary-700 dark:text-primary-400 dark:hover:text-primary-500 dark:focus:text-primary-500 dark:active:text-primary-600"
							>Home</a
						>
					</li>
					<li>
						<span class="mx-2 text-neutral-500 dark:text-neutral-400">/</span>
					</li>
					<li class="text-neutral-500 dark:text-neutral-400">Library</li> -->
				</ol>
			</nav>

			<div class="lg:flex lg:items-center lg:justify-between">
				<div class="min-w-0 flex-1">
					<h2
						class="text-2xl font-bold leading-7 text-gray-900 sm:truncate sm:text-3xl sm:tracking-tight"
					>
						{pageTitle}
					</h2>
					<!-- <div class="mt-1 flex flex-col sm:mt-0 sm:flex-row sm:flex-wrap sm:space-x-6">
						<div class="mt-2 flex items-center text-sm text-gray-500">
							<svg
								class="mr-1.5 h-5 w-5 flex-shrink-0 text-gray-400"
								viewBox="0 0 20 20"
								fill="currentColor"
								aria-hidden="true"
							>
								<path
									fill-rule="evenodd"
									d="M6 3.75A2.75 2.75 0 018.75 1h2.5A2.75 2.75 0 0114 3.75v.443c.572.055 1.14.122 1.706.2C17.053 4.582 18 5.75 18 7.07v3.469c0 1.126-.694 2.191-1.83 2.54-1.952.599-4.024.921-6.17.921s-4.219-.322-6.17-.921C2.694 12.73 2 11.665 2 10.539V7.07c0-1.321.947-2.489 2.294-2.676A41.047 41.047 0 016 4.193V3.75zm6.5 0v.325a41.622 41.622 0 00-5 0V3.75c0-.69.56-1.25 1.25-1.25h2.5c.69 0 1.25.56 1.25 1.25zM10 10a1 1 0 00-1 1v.01a1 1 0 001 1h.01a1 1 0 001-1V11a1 1 0 00-1-1H10z"
									clip-rule="evenodd"
								/>
								<path
									d="M3 15.055v-.684c.126.053.255.1.39.142 2.092.642 4.313.987 6.61.987 2.297 0 4.518-.345 6.61-.987.135-.041.264-.089.39-.142v.684c0 1.347-.985 2.53-2.363 2.686a41.454 41.454 0 01-9.274 0C3.985 17.585 3 16.402 3 15.055z"
								/>
							</svg>
							Full-time
						</div>
						<div class="mt-2 flex items-center text-sm text-gray-500">
							<svg
								class="mr-1.5 h-5 w-5 flex-shrink-0 text-gray-400"
								viewBox="0 0 20 20"
								fill="currentColor"
								aria-hidden="true"
							>
								<path
									fill-rule="evenodd"
									d="M9.69 18.933l.003.001C9.89 19.02 10 19 10 19s.11.02.308-.066l.002-.001.006-.003.018-.008a5.741 5.741 0 00.281-.14c.186-.096.446-.24.757-.433.62-.384 1.445-.966 2.274-1.765C15.302 14.988 17 12.493 17 9A7 7 0 103 9c0 3.492 1.698 5.988 3.355 7.584a13.731 13.731 0 002.273 1.765 11.842 11.842 0 00.976.544l.062.029.018.008.006.003zM10 11.25a2.25 2.25 0 100-4.5 2.25 2.25 0 000 4.5z"
									clip-rule="evenodd"
								/>
							</svg>
							Remote
						</div>
						<div class="mt-2 flex items-center text-sm text-gray-500">
							<svg
								class="mr-1.5 h-5 w-5 flex-shrink-0 text-gray-400"
								viewBox="0 0 20 20"
								fill="currentColor"
								aria-hidden="true"
							>
								<path
									d="M10.75 10.818v2.614A3.13 3.13 0 0011.888 13c.482-.315.612-.648.612-.875 0-.227-.13-.56-.612-.875a3.13 3.13 0 00-1.138-.432zM8.33 8.62c.053.055.115.11.184.164.208.16.46.284.736.363V6.603a2.45 2.45 0 00-.35.13c-.14.065-.27.143-.386.233-.377.292-.514.627-.514.909 0 .184.058.39.202.592.037.051.08.102.128.152z"
								/>
								<path
									fill-rule="evenodd"
									d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-8-6a.75.75 0 01.75.75v.316a3.78 3.78 0 011.653.713c.426.33.744.74.925 1.2a.75.75 0 01-1.395.55 1.35 1.35 0 00-.447-.563 2.187 2.187 0 00-.736-.363V9.3c.698.093 1.383.32 1.959.696.787.514 1.29 1.27 1.29 2.13 0 .86-.504 1.616-1.29 2.13-.576.377-1.261.603-1.96.696v.299a.75.75 0 11-1.5 0v-.3c-.697-.092-1.382-.318-1.958-.695-.482-.315-.857-.717-1.078-1.188a.75.75 0 111.359-.636c.08.173.245.376.54.569.313.205.706.353 1.138.432v-2.748a3.782 3.782 0 01-1.653-.713C6.9 9.433 6.5 8.681 6.5 7.875c0-.805.4-1.558 1.097-2.096a3.78 3.78 0 011.653-.713V4.75A.75.75 0 0110 4z"
									clip-rule="evenodd"
								/>
							</svg>
							$120k &ndash; $140k
						</div>
						<div class="mt-2 flex items-center text-sm text-gray-500">
							<svg
								class="mr-1.5 h-5 w-5 flex-shrink-0 text-gray-400"
								viewBox="0 0 20 20"
								fill="currentColor"
								aria-hidden="true"
							>
								<path
									fill-rule="evenodd"
									d="M5.75 2a.75.75 0 01.75.75V4h7V2.75a.75.75 0 011.5 0V4h.25A2.75 2.75 0 0118 6.75v8.5A2.75 2.75 0 0115.25 18H4.75A2.75 2.75 0 012 15.25v-8.5A2.75 2.75 0 014.75 4H5V2.75A.75.75 0 015.75 2zm-1 5.5c-.69 0-1.25.56-1.25 1.25v6.5c0 .69.56 1.25 1.25 1.25h10.5c.69 0 1.25-.56 1.25-1.25v-6.5c0-.69-.56-1.25-1.25-1.25H4.75z"
									clip-rule="evenodd"
								/>
							</svg>
							Closing on January 9, 2020
						</div>
					</div> -->
				</div>
				<slot name="actions" />
			</div>
		</div>
	</header>
	<main>
		<div class="mx-auto max-w-7xl py-6 lg:px-8 px-6">
			<!-- Your content -->
			<slot />
		</div>
	</main>
</div>
