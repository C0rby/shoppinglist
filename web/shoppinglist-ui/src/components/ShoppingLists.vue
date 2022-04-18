<template>
  <ul class="flex cursor-pointer">
    <li
      v-for="l of store.state.shoppinglists"
      :class="{'bg-emerald-500 text-white': store.state.selectedList == l}"
      class="py-2 px-6 rounded-t-lg bg-gray-200 text-gray-500"
        :key="l.id"
        @click="store.state.selectedList = l"
    >{{ l.name }}</li>
  </ul>
  <!-- <div class="row">
    <h1 v-if="store.state.loading">Loading lists...</h1>
    <div v-else style="overflow-x: scroll; white-space: nowrap">
      <button
        style="margin-right: 8px"
        class="btn"
        :class="{
          'btn-outline-secondary': store.state.selectedList !== l,
          'btn-primary': store.state.selectedList == l,
        }"
        v-for="l of store.state.shoppinglists"
        :key="l.id"
        @click="store.state.selectedList = l"
      >
        {{ l.name }}
      </button>
    </div>
  </div>-->
</template>
<script>
import { onMounted, inject } from "vue";

export default {
  name: "ShoppingLists",
  setup() {
    const store = inject('store')

    onMounted(() => {
      store.methods.fetchLists();
    });

    return {
      store,
    };
  },
};
</script>