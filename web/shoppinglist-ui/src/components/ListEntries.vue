<template>
  <div>
    <div class="row mb-2">
      <h1>{{ list.name }}</h1>
      <div>
        <input
          v-model="search"
          @input="filter"
          class="form-control"
          type="text"
          placeholder="Find"
          aria-label="Find"
          style="background-color: #f3f3f3"
        />
      </div>
    </div>
    <div class="row mb-1">
      <div class="col">
        <input
          v-model="name"
          type="text"
          placeholder="Name"
          aria-label="Name"
          class="form-control"
        />
      </div>
      <div class="col">
        <input
          v-model="amount"
          type="text"
          placeholder="Amount"
          aria-label="Amount"
          class="form-control"
        />
      </div>
      <div class="col-3 col-md-2">
        <button
          class="btn btn-outline-secondary"
          type="button"
          id="btn-add-entry"
          @click="addEntry(list.id)"
        >
          Add
        </button>
      </div>
    </div>
    <div class="row">
      <div class="scrollable list-group list-group-flush">
        <div
          v-for="entry in filteredResults"
          v-bind:key="entry.id"
          class="list-group-item d-flex justify-content-between"
        >
          <label>
            <input
              class="form-check-input me-1"
              type="checkbox"
              :checked="!entry.buy"
              :value="!entry.buy"
              aria-label="..."
              @change="updateEntry(list.id, entry)"
            />
            {{ entry.name }}
          </label>
          <div>
            <span class="badge bg-info me-4">
              {{ entry.amount }}
            </span>
            <button
              class="btn btn-outline-dark"
              type="button"
              id="dropdownMenuButton1"
              data-bs-toggle="dropdown"
              aria-expanded="false"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="16"
                height="16"
                fill="currentColor"
                class="bi bi-three-dots-vertical"
                viewBox="0 0 16 16"
              >
                <path
                  d="M9.5 13a1.5 1.5 0 1 1-3 0 1.5 1.5 0 0 1 3 0zm0-5a1.5 1.5 0 1 1-3 0 1.5 1.5 0 0 1 3 0zm0-5a1.5 1.5 0 1 1-3 0 1.5 1.5 0 0 1 3 0z"
                />
              </svg>
            </button>
            <ul class="dropdown-menu" aria-labelledby="dropdownMenuButton1">
              <li>
                <a class="dropdown-item" href="#" @click="deleteEntry(list.id, entry)"
                  ><svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="16"
                    height="16"
                    fill="currentColor"
                    class="bi bi-trash"
                    viewBox="0 0 16 16"
                  >
                    <path
                      d="M5.5 5.5A.5.5 0 0 1 6 6v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm2.5 0a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm3 .5a.5.5 0 0 0-1 0v6a.5.5 0 0 0 1 0V6z"
                    />
                    <path
                      fill-rule="evenodd"
                      d="M14.5 3a1 1 0 0 1-1 1H13v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V4h-.5a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1H6a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1h3.5a1 1 0 0 1 1 1v1zM4.118 4 4 4.059V13a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4.059L11.882 4H4.118zM2.5 3V2h11v1h-11z"
                    />
                  </svg>
                  Delete</a
                >
              </li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import { ref, toRaw, reactive, onMounted } from "vue";
import useShoppingListEntries from "../store/listentries";

export default {
  name: "ShoppingListEntries",
  props: {
    list: Object,
  },
  setup(props) {
    const { loading, fetchListEntries, listEntries } = useShoppingListEntries();
    const error = reactive({ message: "" });
    const search = ref("");
    const name = ref("");
    const amount = ref("");

    onMounted(() => {
      fetchListEntries(props.list.id);
    });

    function addEntry(id) {
      fetch(
        import.meta.env.VITE_BACKEND_URL +
          "/api/v1/shoppinglists/" +
          id +
          "/entries",
        {
          method: "post",
          headers: {
            "content-type": "application/json",
          },
          body: JSON.stringify({ name: name.value, amount: amount.value }),
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
          name.value = "";
          amount.value = "";
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

    function updateEntry(listId, entry) {
      entry.buy = !entry.buy;
      fetch(
        import.meta.env.VITE_BACKEND_URL +
          "/api/v1/shoppinglists/" +
          listId +
          "/entries/" +
          entry.id,
        {
          method: "put",
          headers: {
            "content-type": "application/json",
          },
          body: JSON.stringify(entry),
        }
      );
    }

    function deleteEntry(listId, entry) {
      fetch(
        import.meta.env.VITE_BACKEND_URL +
          "/api/v1/shoppinglists/" +
          listId +
          "/entries/" +
          entry.id,
        {
          method: "DELETE",
        }
      ).then(() => {
      this.fetchListEntries(props.list.id);
      }
      );

    }

    return {
      loading,
      listEntries,
      search,
      name,
      amount,
      addEntry,
      updateEntry,
      deleteEntry,
      fetchListEntries,
    };
  },
  watch: {
    list: function (nval) {
      this.fetchListEntries(toRaw(nval).id);
    },
  },
  computed: {
    filteredResults() {
      if (!this.search)
        return this.listEntries.sort((a, b) => {
          if (b.buy > a.buy) {
            return 1;
          }
          if (b.buy < a.buy) {
            return -1;
          }
          if (b.name > a.name) {
            return -1;
          }
          if (b.name < a.name) {
            return 1;
          }
        });

      return this.listEntries
        .map((elem) => {
          return toRaw(elem);
        })
        .filter((elem) =>
          elem.name.toLowerCase().includes(this.search.toLowerCase())
        )
        .sort((a, b) => a.buy - b.buy);
    },
  },
};
</script>
<style>
#newItemRow {
  width: 600px;
  margin: auto;
}
#btn-add-entry {
  width: 100%;
}
</style>