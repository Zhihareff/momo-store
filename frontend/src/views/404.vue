<template>
  <div class="d-flex flex-column justify-content-center align-items-center h-100">
    <div class="card">
      <div class="card-header d-flex justify-content-center align-items-center p-2">
        <h2>404</h2>
        <h3 class="text-secondary ms-3">Страница не найдена</h3>
      </div>
      <div class="card-body p-3">
        <p>Нажмите <router-link to="/">сюда</router-link> чтобы вернуться на главную</p>
        <p>Или ждите <strong>{{ remainingTime }}</strong> секунд и вы перенесётесь туда автоматически</p>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  data() {
    return {
      remainingTime: 5
    }
  },
  mounted() {
    const interval = setInterval(() => {
      this.remainingTime--
      if (this.remainingTime === 0) {
        clearInterval(interval)
        this.$router.push('/')
      }
    }, 1000);
  }
})
</script>
