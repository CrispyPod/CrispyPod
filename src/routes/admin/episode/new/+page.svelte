<script lang="ts">
	import { graphqlRequest } from '$lib/graphqlRequest';
	import { get } from 'svelte/store';
	import AdminLayout from '../../AdminLayout.svelte';
	import EpisodeForm from '../../EpisodeForm.svelte';
	import { token } from '$lib/stores/tokenStore';
	import { goto } from '$app/navigation';
	import Popup from '../../../Popup.svelte';
	import type { Episode } from '$lib/models/episode';

	let errMessage: string | null = null;

	let showPopup: boolean = false;
	let popUpTitle: string | null = null;
	let popUpContent: string | null = null;
	let handlePopupConfirm: () => void = () => {};

	async function handleSubmit(e: SubmitEvent, episodeData: Episode) {
		console.log('start processing submit form');

		const form: HTMLFormElement | null = document.querySelector('#newEpisodeForm');
		const formData = new FormData(form!);

		let audioFileField = '';
		if (episodeData.audioFileName != null) {
			// console.log(episodeData);
			audioFileField +=
				',audioFileName:"' +
				episodeData.audioFileName +
				'",audilFileUploadName:"' +
				episodeData.audioFileUploadName;
		}

		const graphqlStr =
			`mutation{createEpisode(input: {title:"` +
			encodeURIComponent(formData.get('title')!.toString()) +
			`",description:"` +
			encodeURIComponent(formData.get('description')!.toString()) +
			`"` +
			audioFileField +
			`}){id,title,description}}`;
		const tokenS = get(token);
		let result = await graphqlRequest(tokenS, graphqlStr);
		var resultJson = await result.json();

		if (resultJson.data != null) {
			goto('/admin/episode');
		} else {
			errMessage = resultJson.errors[0].message;
		}
	}
</script>

<AdminLayout pageTitle="New Episode">
	<EpisodeForm {handleSubmit} {errMessage} />
	{#if showPopup}
		<Popup {popUpTitle} {popUpContent} {handlePopupConfirm} />
	{/if}
</AdminLayout>
