<template>
  <nav class="navbar navbar-expand-sm navbar-dark bg-dark">
    <div class="container-fluid justify-content-end">

      <router-link class="navbar-brand me-auto" to="/"><img src="@/assets/logo.png" style="margin-right: 15px; height: 57px;"/>Пельменная №2</router-link>

      <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNavAltMarkup" aria-controls="navbarNavAltMarkup" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>

      <div class="collapse navbar-collapse" id="navbarNavAltMarkup">

        <div class="navbar-nav me-auto">
        </div>

        <div class="navbar-nav align-items-stretch align-items-sm-center">

          <CartIcon class="d-none d-sm-block" />

        </div>
        
      </div>

      <CartIcon class="d-block d-sm-none ms-3" />

    </div>
  </nav>  
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import CartIcon from './CartIcon.vue';

export default defineComponent({
  components: {
    CartIcon
  },
  methods: {
    handleLogout() {
      this.$store.dispatch('logout').then(() => {
        this.$router.push('/login');
      });
    }
  }
})
</script>
