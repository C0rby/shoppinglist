<script setup lang='ts'>
import { onMounted, ref, watch } from "vue";

interface list {
    id: string
    name: string
}

const url = import.meta.env.VITE_BACKEND_URL + "/api/v1/shoppinglists";
const loading = ref(false)
const shoppinglists = ref<list[]>([])
const selectedList = ref({})

const emit = defineEmits<{
    (e: 'listSelected', id: string): void
}>()

onMounted(() => {
    fetchLists();
});

watch(shoppinglists, (lists) => {
    if (lists.length > 0) {
        selectList(lists[0])
    }
})


async function fetchLists() {
    loading.value = true;
    shoppinglists.value = await (await fetch(url)).json();
    loading.value = false;
}

function selectList(list: list) {
    selectedList.value = list
    emit('listSelected', list.id)
}


</script>
<template>
    <ul class="flex cursor-pointer mx-2 overflow-x-scroll">
        <li v-for="l of shoppinglists"
            :class="{ 'bg-emerald-500 text-white': selectedList == l, 'bg-gray-200': selectedList != l }"
            class="py-2 px-6 rounded-t-lg text-gray-500 min-w-fit" :key="l.id" @click="selectList(l)">
            {{ l.name }}</li>
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



