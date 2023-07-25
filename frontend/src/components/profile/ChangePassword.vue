<template>

  <Alert v-if="error_message" v-model="error_message" class="alert-danger" />

  <Alert v-if="success_message" v-model="success_message" class="alert-success" :dismissible="false"/>

  <form v-else @submit.prevent="handleSubmit" class="my-3">
    <div class="mb-2">
      <input 
        type="password" 
        v-model="oldPassword" 
        class="form-control" 
        placeholder="Old Password" 
        :disabled="loading"
        required
      />
    </div>
    <div class="mb-2">
      <input 
        type="password" 
        v-model="password" 
        class="form-control" 
        placeholder="New password" 
        ref="password" 
        @input="handlePasswordInput"
        :disabled="loading"
        required 
      />
    </div>
    <div class="mb-2">
      <input 
        type="password" 
        v-model="passwordConfirmation" 
        class="form-control" 
        placeholder="Confirm the new password" 
        ref="passwordConfirmation" 
        @input="handlePasswordInput"
        :disabled="loading"
        required 
      />
    </div>
    <div class="d-flex justify-content-end">
      <Center v-if="loading" class="mx-2">
        <Spinner size="2rem" thickness="0.75rem" />
      </Center>
      <button
        type="button" 
        class="btn btn-secondary me-3" 
        @click="$emit('close')"
        :disabled="loading"
      >Cancel</button>
      <button 
        type="submit" 
        class="btn btn-primary"
        :disabled="loading"
      >Change password</button>
    </div>
  </form>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { AxiosError } from 'axios';

import Spinner from '@/components/misc/Spinner.vue';
import Center from '@/components/misc/Center.vue';
import Alert from '@/components/misc/Alert.vue';

export default defineComponent({
  name: 'ChangePassword',
  components: {
    Spinner,
    Center,
    Alert
  },
  emits: [
    'close'
  ],
  data() {
    return {
      oldPassword: '',
      password: '',
      passwordConfirmation: '',
      error: '',
      loading: false,
      error_message: '',
      success_message: ''
    }
  },
  methods: {
    handleSubmit() {
      this.loading = true;

      this.$store.dispatch('changePassword', {
        old_password: this.oldPassword,
        new_password: this.password
      }).then((data: any) => {

        this.success_message = data.detail;
        setTimeout(() => {
          this.$emit('close');
          this.$router.push('/login');
        }, 3000);

      }).catch((error: AxiosError) => {
        if(!error.isAxiosError || !error.response) {
          this.error_message = 'Connection error';
        } else {
          this.error_message = error.response.data.detail;
        }

      }).finally(() => {
        this.loading = false;
      });
    },
    handlePasswordInput() {
      const confirmationInput = this.$refs.passwordConfirmation as HTMLInputElement;
      confirmationInput.setCustomValidity(this.password === this.passwordConfirmation ? '' : 'Passwords do not match');
    }
  }
})
</script>
