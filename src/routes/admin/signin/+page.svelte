<script lang="ts">
	import { goto } from '$app/navigation';
	import SiteLayout from '../../SiteLayout.svelte';
	import { token } from '$lib/stores/tokenStore';
	import { onMount } from 'svelte';
	import { get } from 'svelte/store';

	let errMessage: string | null = null;

	async function handleSubmit() {
		const email = document.getElementById('email') as HTMLInputElement;
		const password = document.getElementById('password') as HTMLInputElement;
		const result = await fetch('/api/login', {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({
				EmailOrUserName: email.value,
				Password: password.value
			})
		});

		if (result.status == 200) {
			const jsonResult = await result.json();
			token.set(jsonResult.token);
			goto('/admin');
		} else {
			const jsonResult = await result.json();
			if (jsonResult.message == null) {
				errMessage = 'Sign in failed';
			} else {
				errMessage = jsonResult.message;
			}
		}
	}

	onMount(() => {
		const curToken = get(token);
		if (curToken != null) {
			goto('/admin');
		}
	});

	function clearErrMessage() {
		errMessage = null;
	}
</script>

<SiteLayout>
	<div class="flex min-h-full flex-col justify-center px-6 py-12 lg:px-8">
		<div class="sm:mx-auto sm:w-full sm:max-w-sm">
			<!-- <img
			class="mx-auto h-10 w-auto"
			src="https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=600"
			alt="Your Company"
		/> -->
			<h2 class="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-gray-900">
				Sign in to your account
			</h2>
		</div>

		<div class="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
			<form class="space-y-6" on:submit|preventDefault={handleSubmit}>
				<div>
					<label for="email" class="block text-sm font-medium leading-6 text-gray-900"
						>Email address or user name</label
					>
					<div class="mt-2">
						<input
							on:focus={clearErrMessage}
							id="email"
							name="email"
							autocomplete="email"
							required
							class="block w-full rounded-md border-0 p-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
						/>
					</div>
				</div>

				<div>
					<div class="flex items-center justify-between">
						<label for="password" class="block text-sm font-medium leading-6 text-gray-900"
							>Password</label
						>
						<div class="text-sm">
							<a href="#" class="font-semibold text-indigo-600 hover:text-indigo-500"
								>Forgot password?</a
							>
						</div>
					</div>
					<div class="mt-2">
						<input
							on:focus={clearErrMessage}
							id="password"
							name="password"
							type="password"
							autocomplete="current-password"
							required
							class="block w-full rounded-md border-0 p-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
						/>
					</div>
				</div>
				{#if errMessage != null}
					<div class=" text-red-500">
						<p>{errMessage}</p>
					</div>
				{/if}
				<div>
					<button
						type="submit"
						class="flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
						>Sign in</button
					>
				</div>
			</form>
		</div>
	</div>
</SiteLayout>
