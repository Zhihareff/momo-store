<template>
  <table class="table table-hover align-middle">
    <thead>
      <tr>
        <th scope="col">Название</th>
        <th scope="col" style="width: 10%">Количество</th>
        <th scope="col">Цена</th>
        <th scope="col">Итого</th>
        <th scope="col"></th>
      </tr>
    </thead>
    <tbody>
      <CartTableRow
        v-for="(cartItem, idx) in $store.getters.cartAsArray"
        :key="idx"
        :cartItem="cartItem"
      />
      <tr v-if="!$store.getters.cartAsArray.length">
        <td colspan="5" class="text-center">Корзина пуста, <router-link to="/catalog">выберите пельмешки в каталоге</router-link></td>
      </tr>
    </tbody>
  </table>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import CartTableRow from '@/components/cart/CartTableRow.vue';

export default defineComponent({
  name: 'Cart',
  components: {
    CartTableRow
  }
})
</script>
