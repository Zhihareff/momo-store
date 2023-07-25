<template>
  <router-link class="mx-0 mx-sm-2 my-2 my-sm-0" to="/cart">
    <i class="fas fa-shopping-cart">
      <span class="badge badge-pill" v-if="$store.getters.numberOfItemsInCart">{{ $store.getters.numberOfItemsInCart }}</span>
    </i>
  </router-link>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  data() {
    const numberOfItems = 10;
    return {
      numberOfItems
    }
  },
})
</script>

<style scoped>
  .fas {
    font-size: 1.5rem;
    color: white;
    position: relative;
  }

  .badge {
    background-color: white;

    position: absolute;
    bottom: -12px;
    left: -12px;

    text-align: center;
    color: white;
    display:block;
    border-radius:50%;

    width:24px;
    height:24px;
    line-height: 24px;

    font-size: 12px;
    font-weight: bold;
    padding: 1px;

    color: black;
    
  }
</style>
