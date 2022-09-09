<script setup lang="ts">

import { ref, watchEffect, toRaw, computed } from 'vue'

interface entry {
    id: string
    name: string
    amount: string
    buy: boolean
}

const props = defineProps<{
    selectedlistId: string
}>()
const conn = new WebSocket(import.meta.env.VITE_WS_URL + "/api/v1/ws")
const loading = ref(false)
const listEntries = ref<entry[]>([])
const search = ref('')

conn.onclose = () => {
    let item = document.createElement("div");
    item.innerHTML = "<b>Connection closed.</b>";
};

conn.onmessage = async (evt) => {
    let listId = evt.data.trim();
    if (listId === props.selectedlistId) {
        await fetchListEntries(listId);
    }
};

async function fetchListEntries(listId: string) {
    const url = import.meta.env.VITE_BACKEND_URL + "/api/v1/shoppinglists/" + listId + "/entries";
    loading.value = true;
    listEntries.value = await (await fetch(url)).json();
    loading.value = false;
}

watchEffect(async () => {
    await fetchListEntries(props.selectedlistId)
})

const updateEntry = (listId: string, entry: entry) => {
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

const deleteEntry = (listId: string, entry: entry) => {
    fetch(
        import.meta.env.VITE_BACKEND_URL +
        "/api/v1/shoppinglists/" +
        listId +
        "/entries/" +
        entry.id,
        {
            method: "DELETE",
        }
    );
}

const filteredResults = computed(() => {
    if (!search.value) {
        return listEntries.value
            .sort((a: entry, b: entry): number => {
                if (b.buy > a.buy) {
                    return 1
                }
                if (b.buy < a.buy) {
                    return -1
                }
                if (b.name > a.name) {
                    return -1
                }
                if (b.name < a.name) {
                    return 1
                }
                return 0
            });
    }

    return listEntries.value
        .map((elem) => {
            return toRaw(elem);
        })
        .filter((elem) =>
            elem.name.toLowerCase().includes(search.value.toLowerCase())
        )
        .sort((a: entry, b: entry): number => {
            if (b.buy > a.buy) {
                return 1
            }
            if (b.buy < a.buy) {
                return -1
            }
            if (b.name > a.name) {
                return -1
            }
            if (b.name < a.name) {
                return 1
            }
            return 0

        });

})
//const { loading, fetchListEntries, listEntries, addEntry } = useShoppingListEntries();
/*
const search = ref("");
const name = ref("");
const amount = ref("");


watch(() => store.state.selectedList, (val) => {
    store.methods.fetchListEntries(val.id);
});




const clearSearch = () => {
    search.value = "";
}


*/
</script>
<template>
    <div class="flex items-center bg-gray-200 rounded-md mt-4 mx-2 md:mx-1">
        <div class="pl-2">
            <svg class="fill-current text-gray-500 w-6 h-6" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
                <path class="heroicon-ui"
                    d="M16.32 14.9l5.39 5.4a1 1 0 0 1-1.42 1.4l-5.38-5.38a8 8 0 1 1 1.41-1.41zM10 16a6 6 0 1 0 0-12 6 6 0 0 0 0 12z" />
            </svg>
        </div>
        <input
            class="w-full rounded-md bg-gray-200 text-gray-700 leading-tight focus:outline-none py-2 px-2 border-0 focus:ring-0 ring-0"
            id="search" type="text" placeholder="Search or add" v-model="search" />
    </div>
    <div class="mt-4 text-sm flex flex-col flex-auto h-0 overflow-y-scroll px-2 md:px-1">
        <div v-for="entry in filteredResults" :key="entry.id"
            class="flex justify-start cursor-pointer text-gray-700 my-4 first:mt-0 last:mb-0 items-center">
            <label class="flex flex-grow justify-start items-start">
                <input type="checkbox"
                    class="w-8 h-8 rounded-full transition-colors duration-400 focus:outline-0 focus:ring-0 focus:ring-offset-0"
                    :class="{ 'text-emerald-500': !entry.buy }" :checked="!entry.buy"
                    @change="updateEntry(props.selectedlistId, entry)" />
                <div class="flex-grow font-bold text-xl ml-2" :class="{ 'text-gray-300 line-through': !entry.buy }">{{
                        entry.name
                }}</div>
            </label>
            <div class=" font-normal text-gray-500 mr-2">{{ entry.amount }}</div>
            <div class="dropdown">
                <button class="text-gray-500 focus:outline-none h-7 flex items-center" type="button"
                    aria-haspopup="true" aria-expanded="true" aria-controls="headlessui-menu-items-117">
                    <svg xmlns="http://www.w3.org/2000/svg" class="text-gray-500 h-6 w-6" fill="none"
                        viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                        <path stroke-linecap="round" stroke-linejoin="round"
                            d="M12 5v.01M12 12v.01M12 19v.01M12 6a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2z" />
                    </svg>
                </button>
                <div
                    class="opacity-0 invisible dropdown-menu transition-all duration-300 transform origin-top-right -translate-y-2 scale-95">
                    <div class="absolute right-0 mt-2 origin-top-right bg-white border border-gray-200 divide-y divide-gray-100 rounded-md shadow-lg outline-none"
                        aria-labelledby="headlessui-menu-button-1" id="headlessui-menu-items-117" role="menu">
                        <div class="py-1">
                            <a href="javascript:void(0)" @click="deleteEntry(props.selectedlistId, entry)" tabindex="0"
                                class="text-gray-700 flex items-center text-lg px-4 py-2" role="menuitem"><svg
                                    xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none"
                                    viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                                    <path stroke-linecap="round" stroke-linejoin="round"
                                        d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                                </svg> Delete</a>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
<style>
.dropdown:focus-within .dropdown-menu {
    opacity: 1;
    transform: translate(0) scale(1);
    visibility: visible;
}
</style>
