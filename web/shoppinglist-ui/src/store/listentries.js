import { reactive, toRefs } from "vue";

const state = reactive({
    listEntries: [],
    loading: true,
});

export default function useShoppingListEntries() {
    const fetchListEntries = async (listId) => {
    const url = import.meta.env.VITE_BACKEND_URL + "/api/v1/shoppinglists/" + listId + "/entries";
        state.loading = true;
        state.listEntries = await (await fetch(url)).json();
        state.loading = false;
    }

    return {
        ...toRefs(state),
        fetchListEntries
    }
}