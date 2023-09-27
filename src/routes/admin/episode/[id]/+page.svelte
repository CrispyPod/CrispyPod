<script lang="ts">
	import { onMount } from 'svelte';
	import AdminLayout from '../../AdminLayout.svelte';
	import EpisodeForm from '../../EpisodeForm.svelte';
	import { graphqlRequest } from '$lib/graphqlRequest';
	import { get } from 'svelte/store';
	import { token } from '$lib/stores/tokenStore';
	import type { Episode } from '$lib/models/episode';
	import { goto } from '$app/navigation';
	import type { AudioFile } from '$lib/models/audioFile';

	let fetchedEpisode: Episode;
	let errMessage: string | null = null;

	export let data;
	onMount(async () => {
		const tokenS = get(token);
		const result = await graphqlRequest(
			tokenS,
			`{episode(id:"` +
				data.id +
				`"){id,title,description,episodeStatus,audioFileName,audioFileUploadName,audioFileDuration}}`
		);
		const jsonResp = await result.json();
		if (jsonResp.data != null) {
			fetchedEpisode = jsonResp.data.episode;
		}
	});

	let showPopup: boolean = false;
	let popUpTitle: string | null = null;
	let popUpContent: string | null = null;
	let handlePopupConfirm: () => void = () => {};

	async function handleSubmit(e: SubmitEvent, episodeData: Episode) {
		const form: HTMLFormElement | null = document.querySelector('#newEpisodeForm');
		const formData = new FormData(form!);
		const toeknS = get(token);

		let audioFileField = '';
		if (episodeData.audioFileName!.length > 0) {
			// console.log(episodeData);
			audioFileField +=
				',audioFileName:"' +
				episodeData.audioFileName +
				'",audioFileUploadName:"' +
				episodeData.audioFileUploadName +
				'",';
		}

		let stat = parseInt(formData.get('status')!.toString());
		const result = await graphqlRequest(
			toeknS,
			`mutation{  modifyEpisode(id:"` +
				episodeData.id +
				`",input: {title:"` +
				encodeURIComponent(formData.get('title')!.toString()) +
				`",description:"` +
				encodeURIComponent(formData.get('description')!.toString()) +
				`",episodeStatus:` +
				stat +
				audioFileField +
				`}){title,description,episodeStatus}}`
		);
		var resultJson = await result.json();

		if (resultJson.data != null) {
			goto('/admin/episode');
		} else {
			errMessage = resultJson.errors[0].message;
		}
	}

	function handleDelete() {
		if (fetchedEpisode.id.length == 0) {
			return;
		}
		let a: any = document.getElementById('confirmDeleteModal');
		a.showModal();
	}

	async function confirmDelete() {
		if (fetchedEpisode.id.length == 0) {
			return;
		}

		const tokenS = get(token);
		let result = await graphqlRequest(
			tokenS,
			`mutation{
  deleteEpisode(id:"` +
				fetchedEpisode.id +
				`"){
    result
  }
}`
		);

		const resultJson = await result.json();
		if (resultJson.data != null && resultJson.data.deleteEpisode.result) {
			goto('/admin/episode');
		} else {
			try {
				errMessage = resultJson.errors[0].message;
			} catch (e) {
				errMessage = 'failed to delete episode';
			}
		}
	}
</script>

<AdminLayout pageTitle="Edit episode">
	<span slot="actions">
		<button class="btn btn-error" on:click|preventDefault={handleDelete}>Delete</button>
	</span>
	<EpisodeForm {handleSubmit} episodeData={fetchedEpisode} {errMessage} />
</AdminLayout>

{#if fetchedEpisode != null}
	<dialog id="confirmDeleteModal" class="modal modal-bottom sm:modal-middle">
		<div class="modal-box">
			<h3 class="font-bold text-lg alert alert-error my-4">Delete confirm</h3>
			<p>Are you sure you want to delete podcast titled</p>
			<br />
			<h2 class="font-bold text-lg">{fetchedEpisode.title}</h2>
			<br />
			<p>Deletion can not be undone.</p>
			<div class="modal-action">
				<form method="dialog">
					<button class="btn btn-error" on:click={confirmDelete}>Confirm</button>
					<!-- if there is a button in form, it will close the modal -->
					<button class="btn">Close</button>
				</form>
			</div>
		</div>
	</dialog>
{/if}
