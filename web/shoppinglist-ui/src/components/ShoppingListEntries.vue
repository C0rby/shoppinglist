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



const deleteEntry = (listId, entry) => {
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

const clearSearch = () => {
    search.value = "";
}


*/
</script>
<template>
    <div class="flex items-center bg-gray-200 rounded-md mt-4">
        <div class="pl-2">
            <svg class="fill-current text-gray-500 w-6 h-6" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
                <path class="heroicon-ui"
                    d="M16.32 14.9l5.39 5.4a1 1 0 0 1-1.42 1.4l-5.38-5.38a8 8 0 1 1 1.41-1.41zM10 16a6 6 0 1 0 0-12 6 6 0 0 0 0 12z" />
            </svg>
        </div>
        <input class="w-full rounded-md bg-gray-200 text-gray-700 leading-tight focus:outline-none py-2 px-2"
            id="search" type="text" placeholder="Search or add" />
    </div>
    <div class="mt-4 text-sm flex flex-col flex-auto h-0 overflow-y-scroll">
        <div v-for="entry in filteredResults" :key="entry.id"
            class="flex justify-start cursor-pointer text-gray-700 hover:text-blue-400 hover:bg-blue-100 rounded-md px-2 py-2 my-2">
            <label class="flex flex-grow justify-start items-start">
                <div
                    class="bg-white border-2 rounded border-gray-400 w-6 h-6 flex flex-shrink-0 justify-center items-center mr-2 focus-within:border-emerald-500">
                    <input type="checkbox" class="opacity-0 absolute" :checked="!entry.buy"
                        @change="updateEntry(props.selectedlistId, entry)" />
                    <svg class="fill-current hidden w-4 h-4 text-emerald-500 pointer-events-none" viewBox="0 0 20 20">
                        <path d="M0 11l2-2 5 5L18 3l2 2L7 18z" />
                    </svg>
                </div>
                <div class="flex-grow font-medium px-2">{{ entry.name }}</div>
            </label>
            <div class="text-sm font-normal text-gray-500 tracking-wide">{{ entry.amount }}</div>
        </div>
    </div>
</template>
<style>
input:checked+svg {
    display: block;
}
</style>
