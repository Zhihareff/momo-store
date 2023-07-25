<template>
  <div class="card">
    <img :src="product.image" class="card-img-top" :alt='`Изображение "${product.name}"`'>
    <div class="card-body">
      <h5 class="card-title">{{product.name}}</h5>
      <h6 class="card-subtitle fw-bold">${{product.price}}</h6>
      <p class="card-text">{{product.description}}</p>
      <button @click="addToCart" class="btn btn-primary">Добавить в корзину</button>
    </div>
  </div>
</template>

<script lang="ts">
  import { defineComponent } from 'vue';
  import ProductModel from '@/models/product.model';

  export default defineComponent({
    name: 'Product',
    props: {
      product: {
        type: ProductModel,
        required: true
      }
    },
    methods: {
      addToCart() {
        this.$store.commit('addToCart', this.product);
      }  
    }
  });

</script>

<style scoped>

  img {
    aspect-ratio: 3/2;
    object-fit: cover;
  }

</style>