<template>

  <div class="grid" v-if="products.results.length">
    <Product v-for="(product, idx) in products.results" :key="idx" :product="product" />
  </div>

  <div class="text-center mt-4" v-else>
    <h3 class="text-secondary">Продукты не найдены</h3>
  </div>

</template>

<script lang="ts">
  import { defineComponent } from 'vue';

  import Product from './Product.vue';

  import PaginationModel from '@/models/pagination.model';
  import ProductModel from '@/models/product.model';

  export default defineComponent({
    name: 'Products',
    components: {
      Product
    },
    props: {
      products: {
        type: Object as () => PaginationModel<ProductModel>,
        required: true
      }
    }
  });

</script>

<style scoped>
  .grid {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: 2rem;
  }

  @media screen and (max-width: 768px) {
    .grid {
      grid-template-columns: repeat(2, 1fr);
    }
  }

  @media screen and (max-width: 576px) {
    .grid {
      grid-template-columns: 1fr;
    }
  }
</style>