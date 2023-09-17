import type { AudioFile } from './audioFile';

export enum EpisodeState {
	draft = 0,
	published = 1
}

export class Episode {
	id: string;
	title: string;
	description: string;
	createTime: Date;
	publishTime: Date | null;
	episodeStatus: number;
	// audioFiles: AudioFile[] = [];
	thumbnailFileName: string | null = null;

	audioFileName: string | null = null;
	audioFileUploadName: string | null = null;
	audioFileDuration: number | null = null;

	constructor(
		id: string,
		title: string,
		description: string,
		createTime: Date,
		episodeState: number,
		publishTime: Date | null
	) {
		this.id = id;
		this.title = title;
		this.episodeStatus = episodeState;
		this.description = description;
		this.createTime = createTime;
		this.publishTime = publishTime;
	}
}
