<template>
  <div class="card">
    <ul class="list-group list-group-flush">

      <li class="list-group-item">
        <router-link class="cart-link" to="/cart">
          <h4 class="card-title mt-1 mb-3">Корзина</h4>
        </router-link>
      </li>

      <CartCardRow v-for="(cartItem, idx) in $store.getters.cartAsArray" :key="idx" :cartItem="cartItem" />

      <li class="list-group-item list-bottom" v-if="!$store.getters.cartAsArray.length">
        <p class="text-center m-0">Корзина пуста</p>
      </li>

      <li class="list-group-item list-bottom">
        <span class="col">Итого: <strong>${{$store.getters.cartTotal}}</strong></span>
        <router-link class="cart-link" to="/cart">
            <button class="btn btn-primary" :disabled="!$store.getters.numberOfItemsInCart">Оформить</button>
        </router-link>
      </li>

    </ul>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

import CartCardRow from './CartCardRow.vue';

export default defineComponent({
  name: 'CartCard',
  components: {
    CartCardRow
  }
})
</script>

<style scoped>
  .cart-link {
    text-decoration: none;
    color: black;
  }

  .cart-link:hover {
    color: #285388;
  }

  .list-bottom {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

</style>