export class SiteConfig {
    siteName: string;
    siteDescription: string;

    constructor(siteName: string,
        siteDescription: string) {
        this.siteName = siteName;
        this.siteDescription = siteDescription;
    }
}