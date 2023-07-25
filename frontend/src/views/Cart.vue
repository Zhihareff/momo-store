<template>
  <div class="container">
    <h1 class="text-center">Корзина</h1>
    <CartTable />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import CartTable from '@/components/cart/CartTable.vue';

export default defineComponent({
  name: 'Cart',
  components: {
    CartTable
  }
})
</script>
