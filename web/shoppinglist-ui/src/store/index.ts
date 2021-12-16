import { reactive, readonly } from "@vue/reactivity";

const state = reactive({
    lists: [],
    currentList: {},
});

export default { state: readonly(state) };