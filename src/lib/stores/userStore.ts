import { writable } from "svelte/store";


function userStore() {
    const { subscribe, set } = writable(null);

    return {
        subscribe,
        set
    };
}

export const user = userStore();
