<template>
  <div class="row">
    <h1 v-if="loading">Loading lists...</h1>
    <div v-else style="overflow-x: scroll; white-space: nowrap">
      <button
        style="margin-right: 8px"
        class="btn"
        :class="{
          'btn-outline-secondary': selectedList !== l,
          'btn-primary': selectedList == l,
        }"
        v-for="l of shoppinglists"
        :key="l.id"
        @click="selectedList = l"
      >
        {{ l.name }}
      </button>
    </div>
  </div>
</template>
<script>
import { onMounted } from "vue";
import useShoppingLists from "../store/shoppinglists";

export default {
  name: "ShoppingLists",
  setup() {
    const { shoppinglists, fetchLists, selectedList, loading } = useShoppingLists();

    onMounted(() => {
      fetchLists();
    });

    return {
      shoppinglists,
      selectedList,
      loading,
    };
  },
};
</script>