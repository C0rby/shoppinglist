import { reactive, toRefs } from "vue";

const state = reactive({
    listEntries: [],
    loading: true,
});

function addEntry(listId, name, amount) {
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
            const error = new Error(res.statusText);
            error.json = res.json();
            throw error;
          }
          return res.json();
        })
        .then((json) => {
          listEntries.value.push(json);
        })
        .catch((err) => {
          error.value = err;
          if (err.json) {
            return err.json.then((json) => {
              error.value.message = json.message;
            });
          }
        });
    }

export default function useShoppingListEntries() {
    const fetchListEntries = async (listId) => {
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