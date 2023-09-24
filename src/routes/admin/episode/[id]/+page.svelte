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
				episodeData.audioFileUploadName+'",';
		}

		let stat = parseInt(formData.get('status')!.toString());
		const result = await graphqlRequest(
			toeknS,
			`mutation{  modifyEpisode(id:"` +
				episodeData.id +
				`",input: {title:"` +
				formData.get('title') +
				`",description:"` +
				formData.get('description') +
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
</script>

<AdminLayout pageTitle="Edit episode">
	<EpisodeForm {handleSubmit} episodeData={fetchedEpisode} {errMessage} />
</AdminLayout>
