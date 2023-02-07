<script setup lang='ts'>

import { ref } from "vue";
import ShoppingLists from "./components/ShoppingLists.vue";
import ShoppingListEntries from "./components/ShoppingListEntries.vue";
import NavBar from './components/nav/NavBar.vue';

const selectedList = ref("")

function selected(id: string) {
    selectedList.value = id
}
</script>

<template>
    <NavBar></NavBar>
    <div class="w-screen md:w-7/12 xl:w-2/5 mx-auto flex flex-col flex-1">
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
        <shopping-lists @list-selected="selected" />
        <shopping-list-entries :selectedlist-id="selectedList" />
    </div>
    <!-- </div> -->
</template>
