<script setup>

import { provide } from "vue";
import store from "./store";
import ShoppingLists from "./components/ShoppingLists.vue";
import ListEntries from "./components/ListEntries.vue";
import ShoppingListEntries from "./components/ShoppingListEntries.vue";
import NavBar from './components/nav/NavBar.vue';


provide('store', store)
// export default {
//   setup() {
//     provide('store', store)

//     return {};
//   },
</script>

<template>
  <NavBar></NavBar>
  <div class="p-2 w-screen md:w-7/12 xl:w-2/5 mx-auto flex flex-col flex-1 bg-white">
    <!-- <div class="container-lg vh-100"> -->
    <!-- <header
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
    </header>-->
    <!-- <div style="max-width: 600px; margin: auto"> -->
    <shopping-lists />
    <!-- <list-entries /> -->
    <shopping-list-entries />
  </div>
  <!-- </div> -->
</template>
