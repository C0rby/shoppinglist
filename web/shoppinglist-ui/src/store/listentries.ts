import { reactive, toRefs } from "vue";

interface State {
    listEntries: Object[]
    loading: boolean
}

const state: State = reactive({
    listEntries: [],
    loading: true,
});

function addEntry(listId: string, name: string, amount: string) {
      fetch(
        import.meta.env.VITE_BACKEND_URL +
          "/api/v1/shoppinglists/" +
          listId +
          "/entries",
        {
          method: "post",
          headers: {
            "content-type": "application/json",
          },
          body: JSON.stringify({ name: name, amount: amount}),
        }
      )
        .then((res) => {
          if (!res.ok) {
            const error: any = new Error(res.statusText);
            error.json = res.json();
            throw error;
          }
          return res.json();
        })
        .then((json: Object) => {
          state.listEntries.push(json);
        })
        .catch((err: any) => {
            // TODO properly handle errors
        });
    }

export default function useShoppingListEntries() {
    const fetchListEntries = async (listId: string) => {
    const url = import.meta.env.VITE_BACKEND_URL + "/api/v1/shoppinglists/" + listId + "/entries";
        state.loading = true;
        state.listEntries = await (await fetch(url)).json();
        state.loading = false;
    }

    return {
        ...toRefs(state),
        fetchListEntries,
        addEntry
    }
}
