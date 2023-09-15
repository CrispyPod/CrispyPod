import type { AudioFile } from "./audioFile";

export enum EpisodeState {
    draft = "DRAFT",
    published = "PUBLISHED",
}

export class Episode {
    id: string;
    title: string;
    description: string;
    createTime: Date;
    publishTime: Date | null;
    episodeState: EpisodeState;
    // audioFiles: AudioFile[] = [];
    thumbnailFileName: string | null = null;

    audioFileName: string | null = null;
    audioFileUploadFileName: string | null = null;
    audioFileDuration: number | null = null;

    constructor(id: string,
        title: string,
        description: string,
        createTime: Date,
        episodeState: EpisodeState,
        publishTime: Date | null) {
        this.id = id;
        this.title = title;
        this.episodeState = episodeState;
        this.description = description;
        this.createTime = createTime;
        this.publishTime = publishTime;
    }
}