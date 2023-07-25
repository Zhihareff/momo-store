<template>
  <div class="container d-flex flex-column h-100">
    <div class="row">
      <div class="col-12 col-lg-9">
        <h1 class="text-center mb-3">Каталог</h1>
      </div>
    </div>
    <div class="row flex-fill">

      <div class="col-12 col-lg-9">
        <div class="d-flex flex-column h-100">

          <div class="flex-fill">

            <div v-if="loading" class="d-flex align-items-center justify-content-center h-100" >
              <div class="spinner-wrapper">
                <div class="spinner-border" role="status">
                  <span class="visually-hidden">Загрузка...</span>
                </div>
              </div>
            </div>

            <Products v-else :products="products"  />

          </div>

          <div>
            <Pagination :products="products" @pageChange="fetchProducts" />
          </div>
        </div>
      </div>

      <div class="col-12 col-lg-3 my-3 my-lg-0">
        <Search @search="handleSearch" :query="search_query"/>
        <br />
        <CartCard />
        <br />
        <Filters @change="handleFiltersChange"/>
      </div>

    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

import Products from "@/components/catalog/Products.vue";
import Pagination from "@/components/catalog/Pagination.vue";
import Search from "@/components/catalog/Search.vue";
import CartCard from "@/components/catalog/CartCard.vue";
import Filters from "@/components/catalog/Filters.vue";

import ProductsService from "@/services/resources/products.service";
import ProductModel from "@/models/product.model";
import PaginationModel from "@/models/pagination.model";

const productsService = new ProductsService();

const PAGE_SIZE = 12;

export default defineComponent({
  name: "Catalog",
  components: {
    Products,
    Pagination,
    Search,
    CartCard,
    Filters,
  },
  data() {
    const products = PaginationModel.empty(ProductModel);
    const categories_filter: number[] = [];
    return {
      products,
      categories_filter,
      search_query: "",
      loading: true,
    };
  },
  created() {
    this.fetchProducts();
  },
  methods: {
    fetchProducts(page?: number) {
      page = page || this.products.current_page;
      this.loading = true;
      return productsService
        .query({
          page,
          size: PAGE_SIZE,
          categories: this.categories_filter,
          search: this.search_query || undefined,
        })
        .then((products) => {
          this.products = products;
          this.loading = false;
        });
    },
    handleSearch(query: string) {
      this.search_query = query;
      this.products.current_page = 1;
      this.fetchProducts();
    },
    handleFiltersChange(categories: number[]) {
      this.categories_filter = categories;
      this.products.current_page = 1;
      this.fetchProducts();
    },
  },
});
</script>

<style scoped>
  .spinner-wrapper {
    width: 100%;
    text-align: center;
  }

  .spinner-border {
    width: 8rem;
    height: 8rem;
    font-size: 1.5rem;
  }
</style>