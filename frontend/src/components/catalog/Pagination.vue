<template>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import PaginationModel from '@/models/pagination.model';
import ProductModel from '@/models/product.model';

export default defineComponent({
  props: {
    products: {
      type: Object as () => PaginationModel<ProductModel>,
      required: true
    }
  },
  emits: ['pageChange'],
  methods: {
    setPage(page: number) {
      if (page > 0 && page <= this.products.total_pages) {
        this.products.current_page = page;
        this.$emit('pageChange', page);
      }
    },
    handleBur(event: FocusEvent) {
      const target = event.target as HTMLInputElement;
      this.setPage(Number(target.value));
      target.value = this.products.current_page.toString();
    }
  }
})
</script>

<style scoped>

  input {
    width: 100px;
  }

</style>