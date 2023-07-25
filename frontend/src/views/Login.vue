<template>
  <div class="container d-flex flex-column justify-content-between align-items-stretch h-100">

    <div>
      <Alert v-if="error" v-model="error" class="alert-danger"/>
    </div>

    <Center>

      <Spinner 
        v-if="loading"
        size="8rem"
        thickness="1.25rem"
      />

      <div v-else class="card d-inline-block text-start">
        <form @submit.prevent="login">

          <div class="card-header">
            <h3 class="card-title">Login</h3>
          </div>

          <div class="card-body">
            <div class="form-group mb-3">
              <label for="email">Email</label>
              <input type="email" class="form-control" id="email" v-model="email" />
            </div>
            <div class="form-group mb-3">
              <label for="password">Password</label>
              <input type="password" class="form-control" id="password" v-model="password" />
            </div>
          </div>

          <div class="card-footer text-end">
            <button type="submit" class="btn btn-primary">Login</button>
          </div>

        </form>
      </div>

    </Center>

    <div></div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { AxiosError } from 'axios';

import Spinner from '@/components/misc/Spinner.vue'
import Center from '@/components/misc/Center.vue'
import Alert from '@/components/misc/Alert.vue'

export default defineComponent({
  name: 'Login',
  components: {
    Spinner,
    Center,
    Alert
  },
  data() {
    return {
      email: '',
      password: '',
      loading: false,
      error: '',
      goto: ''
    }
  },
  created() {
    
    if(this.$store.state.auth.isLoggedIn) {
      return this.$router.push('/profile');
    }

    const url = new URL(window.location.href);
    this.goto = url.searchParams.get('goto') as string;
  },
  methods: {
    login() {
      this.loading = true;
      this.$store.dispatch('login', {
        email: this.email,
        password: this.password
      }).then((result: any) => {
        this.$router.push(this.goto || '/');
      }).catch((error: AxiosError) => {

        if(!error.isAxiosError || !error.response) {
          this.error = 'Connection error';
        } else {
          this.error = error.response.data.detail;
        }

        this.loading = false;

      });
    }
  }
})
</script>

<style scoped>
  .card {
    width: 400px;
  }
</style>