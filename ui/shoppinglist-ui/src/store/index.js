import { reactive, readonly } from "@vue/reactivity";

const url = import.meta.env.VITE_BACKEND_URL + "/api/v1/shoppinglists";

const state = reactive({
    shoppinglists: [],
    loading: true,
    selectedList: null,
    name: '',
    amount: '',
    entries: [],
    listEntries: [],
    loadingEntries: true,
});

const methods = {
    async fetchLists() {
        state.loading = true;
        state.shoppinglists = await (await fetch(url)).json();
        state.loading = false;
    },
    async fetchListEntries(listId) {
        const url = import.meta.env.VITE_BACKEND_URL + "/api/v1/shoppinglists/" + listId + "/entries";
        state.loadingEntries = true;
        state.listEntries = await (await fetch(url)).json();
        state.loadingEntries = false;
    }
}

export default {
    state: state,
    methods,
};


