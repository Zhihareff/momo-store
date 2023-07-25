<template>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

import CategoriesService from '@/services/resources/categories.service';
import CategoryModel from '@/models/category.model';

const categoriesService = new CategoriesService();

export default defineComponent({
  name: 'Filters',
  created() {
    this.fetchCategories();
  },
  data() {
    const categories: CategoryModel[] = [];
    const selected: number[] = [];
    return {
      categories,
      loading: true,
      selected
    }
  },
  emits: [
    'change'
  ],
  methods: {
    async fetchCategories() {
      this.loading = true;
      this.categories = await categoriesService.query();
      this.loading = false;
    },
    handleChange(event: Event) {
      const target = event.target as HTMLInputElement;
      const category_id = parseInt(target.value);
      
      if(target.checked) {
        this.selected.push(category_id);
      } else {
        this.selected = this.selected.filter(id => id !== category_id);
      }

      this.$emit('change', this.selected);
    }
  }
})
</script>
