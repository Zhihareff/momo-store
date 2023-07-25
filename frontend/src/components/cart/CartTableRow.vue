<template>
  <tr>
    <td>{{ cartItem.product.name }}</td>
    <td>
      <input 
        type="number" 
        class="form-control" 
        v-model="cartItem.quantity"
        min="0"
      />
    </td>
    <td>{{ cartItem.product.price }}</td>
    <td>{{ cartItem.product.price * cartItem.quantity }}</td>
    <td>
      <button class="btn-close" @click="removeFromCart"/>
    </td>
  </tr>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { CartItem } from '@/typings';

export default defineComponent({
  name: "CartCardRow",
  props: {
    cartItem: {
      type: Object as () => CartItem,
      required: true
    }
  },
  methods: {
    removeFromCart() {
      this.$store.commit('removeFromCart', this.$props.cartItem.product);
    }
  }
})
</script>