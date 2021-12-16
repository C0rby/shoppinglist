<template>
  <div v-if="list && list.name">
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
          <span class="badge bg-info square-pill">
            {{ entry.amount }}
          </span>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import { ref, toRaw, reactive } from "vue";

export default {
  name: "ShoppingListEntries",
  props: {
    list: Object,
  },
  setup() {
    const data = reactive({ list: [] });
    const loading = ref(true);
    const error = reactive({ message: "" });
    const search = ref("");
    const name = ref("");
    const amount = ref("");

    function fetchData(id) {
      loading.value = true;
      return fetch(
        import.meta.env.VITE_BACKEND_URL + "/api/v1/shoppinglists/" +
          this.list.id +
          "/entries",
        {
          method: "get",
          headers: {
            "content-type": "application/json",
          },
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
          data.list = json;
        })
        .catch((err) => {
          error.value = err;
          if (err.json) {
            return err.json.then((json) => {
              error.value.message = json.message;
            });
          }
        })
        .then(() => {
          loading.value = false;
        });
    }

    function addEntry(id) {
      fetch(
        import.meta.env.VITE_BACKEND_URL + "/api/v1/shoppinglists/" + id + "/entries",
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
          data.list.push(json);
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
        import.meta.env.VITE_BACKEND_URL + "/api/v1/shoppinglists/" +
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

    return {
      data,
      loading,
      error,
      fetchData,
      search,
      name,
      amount,
      addEntry,
      updateEntry,
    };
  },
  watch: {
    list: function () {
      this.fetchData();
    },
  },
  computed: {
    filteredResults() {
      if (!this.search)
        return this.data.list.sort((a, b) => {
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

      return this.data.list
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