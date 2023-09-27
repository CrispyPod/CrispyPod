<script lang="ts">
	import { Episode, EpisodeState } from '$lib/models/episode';
	import { token } from '$lib/stores/tokenStore';
	import { get } from 'svelte/store';
	import WaveForm from '../WaveForm.svelte';
	import type { AudioFile } from '$lib/models/audioFile';
	import { onMount } from 'svelte';
	import { graphqlRequest } from '$lib/graphqlRequest';
	import { goto } from '$app/navigation';

	export let handleSubmit: (e: SubmitEvent, episodeData: Episode) => Promise<any>;

	async function onFormSubmit(e: SubmitEvent) {
		handleSubmit(e, episodeData!);
	}

	// export let
	export let episodeData: Episode = new Episode('', '', '', new Date(), EpisodeState.draft, null);
	export let errMessage: string | null = null;
	let fileUploadedAndNotSaved = false;

	let fileList: FileList;
	// Note that `fileList` is of type `FileList`, not an Array:
	// https://developer.mozilla.org/en-US/docs/Web/API/FileList

	onMount(() => {
		(document!.getElementById('afUpload') as HTMLInputElement)!.value = '';
	});

	$: if (episodeData) {
		if (episodeData == null) {
			fileUploadedAndNotSaved = true;
		}
	}

	function fileListChanged() {
		startUpload();
	}

	$: fileList && fileListChanged();

	function formatBytes(bytes: number, decimals = 2) {
		if (!+bytes) return '0 Bytes';

		const k = 1024;
		const dm = decimals < 0 ? 0 : decimals;
		const sizes = ['Bytes', 'KiB', 'MiB', 'GiB', 'TiB', 'PiB', 'EiB', 'ZiB', 'YiB'];

		const i = Math.floor(Math.log(bytes) / Math.log(k));

		return `${parseFloat((bytes / Math.pow(k, i)).toFixed(dm))} ${sizes[i]}`;
	}

	async function startUpload() {
		let file = fileList.item(0);
		const tokenS = get(token);
		// console.log(tokenS);

		let data = new FormData();
		data.append('file', file!);
		data.append('episodeId', episodeData.id);
		let resp = await fetch('/api/audioFile', {
			method: 'POST',
			headers: [['Authorization', 'Bearer ' + tokenS]],
			body: data
		});

		if (resp.status != 200) {
			// TODO: show popup
		}

		let audioFile: AudioFile = await resp.json();

		episodeData!.audioFileName = audioFile.audioFileName;
		episodeData!.audioFileUploadName = file?.name!;
		// episodeData!.audioFileDuration = audioFile.audioFileDuration;
		fileUploadedAndNotSaved = true;
	}

	function handleReupload() {
		episodeData!.audioFileName = null;
		episodeData!.audioFileUploadName = null;
	}
</script>

<form id="newEpisodeForm" on:submit|preventDefault={onFormSubmit}>
	<div class="space-y-12">
		<div class="border-b border-gray-900/10 pb-12">
			<div class="mt-10 grid grid-cols-1 gap-x-6 gap-y-8 sm:grid-cols-6">
				<div class="sm:col-span-4">
					<label for="title" class="block text-sm font-medium leading-6 text-gray-900">Title</label>
					<div class="mt-2">
						<input
							required
							id="title"
							name="title"
							value={episodeData == null ? '' : episodeData.title}
							class="block w-full px-2 rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
						/>
					</div>
				</div>

				<div class="col-span-full">
					<label for="description" class="block text-sm font-medium leading-6 text-gray-900"
						>Description</label
					>
					<div class="mt-2">
						<textarea
							required
							id="description"
							name="description"
							rows="3"
							value={episodeData == null ? '' : episodeData.description}
							class="block w-full rounded-md px-2 border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
						/>
					</div>
				</div>
			</div>
		</div>

		<div class="border-b border-gray-900/10 pb-12">
			{#if episodeData == null || episodeData.audioFileName == null || episodeData.audioFileName.length == 0}
				<div>
					<p>Upload audio file</p>
					<input
						id="afUpload"
						type="file"
						bind:files={fileList}
						class="file-input file-input-bordered w-full max-w-xs"
					/>
				</div>
			{:else}
				{#if fileUploadedAndNotSaved}
					Please hit save so that changes are committed.
				{/if}
				<WaveForm fileUrl="/api/audioFile/{episodeData.audioFileName}" />

				<div class="w-full flex">
					<button
						class="btn btn-outline btn-secondary ml-auto"
						on:click|preventDefault={handleReupload}>Reuplaod</button
					>
				</div>
			{/if}

			<div class="mt-10 space-y-10">
				{#if episodeData != null && episodeData.id != ''}
					<fieldset>
						<legend class="text-sm font-semibold leading-6 text-gray-900">Status</legend>
						<div class="mt-6 space-y-6">
							<div class="flex items-center gap-x-3">
								<input
									required
									id="status-draft"
									name="status"
									type="radio"
									value="0"
									checked={episodeData != null && episodeData.episodeStatus == EpisodeState.draft}
									class="h-4 w-4 border-gray-300 text-indigo-600 focus:ring-indigo-600"
								/>
								<label for="status-draft" class="block text-sm font-medium leading-6 text-gray-900"
									>Draft</label
								>
							</div>
							<div class="flex items-center gap-x-3">
								<input
									required
									id="status-published"
									name="status"
									type="radio"
									value="1"
									checked={episodeData != null &&
										episodeData.episodeStatus == EpisodeState.published}
									class="h-4 w-4 border-gray-300 text-indigo-600 focus:ring-indigo-600"
								/>
								<label
									for="status-published"
									class="block text-sm font-medium leading-6 text-gray-900">Published</label
								>
							</div>
						</div>
					</fieldset>
				{:else}
					Newly created episodes will be saved as draft.
				{/if}
			</div>
		</div>
	</div>

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
		<a href="/admin/episode" type="button" class="btn btn-active">Cancel</a>
		<button type="submit" class="btn btn-active btn-primary">Save</button>
	</div>
</form>
