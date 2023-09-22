export class SiteConfig {
	siteName: string;
	siteDescription: string;
	siteFullDescription: string;
	siteUrl: string;

	constructor(
		siteName: string,
		siteDescription: string,
		siteFullDescription: string,
		siteUrl: string
	) {
		this.siteName = siteName;
		this.siteDescription = siteDescription;
		this.siteFullDescription = siteFullDescription;
		this.siteUrl = siteUrl;
	}
}
