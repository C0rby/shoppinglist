import { reactive, toRefs } from "vue";

const url = import.meta.env.VITE_BACKEND_URL + "/api/v1/shoppinglists";

const state = reactive({
    shoppinglists: [],
    loading: true,
    selectedList: null
});

export default function useShoppingLists() {
    const fetchLists = async () => {
        state.loading = true;
        state.shoppinglists = await (await fetch(url)).json();
        state.loading = false;
    }

    return {
        ...toRefs(state),
        fetchLists
    }
}