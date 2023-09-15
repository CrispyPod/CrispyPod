import { writable } from "svelte/store";

function siteConfigStore() {
    const { subscribe, set } = writable('');

    return {
        subscribe, set
    }
}

export const siteConfigS = siteConfigStore();