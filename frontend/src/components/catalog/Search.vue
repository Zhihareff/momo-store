<template>

</template>

<script lang="ts">
import { defineComponent } from "vue";
import Debouncer from 'svelte-debouncer';

const DEBOUNCE_MS = 1000;

export default defineComponent({
  name: "Search",
  data() {
    return {
      search_debouncer: new Debouncer((q: string) => this.$emit("search", q.trim()), DEBOUNCE_MS, true),
    }
  },
  props: {
    query: {
      type: String,
      default: ""
    }
  },
  emits: [ "search" ]
});
</script>
