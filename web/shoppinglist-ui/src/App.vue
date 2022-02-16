<script>
// https://dev.to/burakgur/create-todo-app-with-vue-3-composition-api-1ok7

import { provide } from "vue";
import store from "./store";
import ShoppingLists from "./components/ShoppingLists.vue";
import ShoppingListEntries from "./components/ListEntries.vue";

export default {
  setup() {
    provide('store', store)

    return {};
  },
  components: {
    ShoppingLists,
    ShoppingListEntries,
  },
};
</script>

<template>
  <div class="container-lg vh-100">
    <header
      class="d-flex flex-wrap justify-content-center py-3 mb-4 border-bottom"
    >
      <a
        href="/shoppinglist"
        class="
          d-flex
          align-items-center
          mb-3 mb-md-0
          me-md-auto
          text-dark text-decoration-none
        "
      >
        <svg class="bi me-2" width="40" height="32">
          <use xlink:href="#bootstrap"></use>
        </svg>
        <span class="fs-4">ShoppingList</span>
      </a>

      <div class="text-end">
        <button type="button" class="btn btn-light me-2">Logout</button>
      </div>
    </header>
    <div style="max-width: 600px; margin: auto">
      <shopping-lists />
      <shopping-list-entries />
    </div>
  </div>
</template>
