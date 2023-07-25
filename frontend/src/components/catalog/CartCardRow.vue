<template>
  <li 
    class="list-group-item d-flex justify-content-between align-items-center"
    @mouseenter="hovering = true"
    @mouseleave="hovering = false"
  >
    <span>{{ cartItem.product.name }}</span>
    <div class="right-section">
      <button v-if="hovering" @click="removeFromCart" class="btn btn-close"></button>
      <span v-if="!hovering" class="bg-primary quantity">{{ cartItem.quantity }}</span>
    </div>
  </li>
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
  data() {
    return {
      hovering: false
    }
  },
  methods: {
    removeFromCart() {
      this.$store.commit('removeFromCart', this.$props.cartItem.product);
    }
  }
})
</script>

<style scoped>
  .quantity {
    text-align: center;
    color: white;
    display:block;
    border-radius:50%;
    width:24px;
    line-height:24px
  }

  .right-section {
    width: 24px;
    height: 24px;
  }

  .btn-close {
    padding: 0;
  }
</style>
