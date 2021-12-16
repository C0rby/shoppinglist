<template>
  <div class="row">
    <div style="overflow-x: scroll; white-space: nowrap">
      <button
        style="margin-right: 8px"
        class="btn"
        :class="{
          'btn-outline-secondary': current.value !== list,
          'btn-primary': current.value == list,
        }"
        v-for="list of data"
        v-bind:key="list.id" 
        @click="select($event, list)"
      >
        {{ list.name }}
      </button>
    </div>
  </div>
</template>
<script>
import { ref, onMounted, toRaw} from "vue";

export default {
  name: "ShoppingLists",
  props: {},
  setup() {
    const data = ref(null);
    const loading = ref(true);
    const error = ref(null);
    const current = ref({});

    function fetchData() {
      loading.value = true;
      return fetch(import.meta.env.VITE_BACKEND_URL + "/api/v1/shoppinglists", {
        method: "get",
        header: {
          "content-type": "application/json",
        },
      })
        .then((res) => {
          if (!res.ok) {
            const error = new Error(res.statusText);
            error.json = res.json();
            throw error;
          }
          return res.json();
        })
        .then((json) => {
          data.value = json;
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

    onMounted(() => {
      fetchData();
    });

    return {
      data,
      loading,
      error,
      current
    };
  },
  methods: {
    select($event, list) {
      this.current.value = list;
      this.$emit("selected", toRaw(list));
    },
  },
};
</script>