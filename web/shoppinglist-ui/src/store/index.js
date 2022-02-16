import { reactive, readonly } from "@vue/reactivity";

const url = import.meta.env.VITE_BACKEND_URL + "/api/v1/shoppinglists";

const state = reactive({
    shoppinglists: [],
    loading: true,
    selectedList: null,
    name: '',
    amount: '',
    entries: [],
});

const methods = {
    async fetchLists() {
        state.loading = true;
        state.shoppinglists = await (await fetch(url)).json();
        state.loading = false;
    }
}

export default {
    state: state,
    methods,
};


