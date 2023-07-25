<template>
  <div 
    class="spinner-border" 
    role="status"
    :style="{ 
      'width': size, 
      'height': size, 
      'font-size': thickness, 
      'color': color 
    }"
  >
    <span class="visually-hidden">Loading...</span>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'Spinner',
  props: {
    size: {
      type: String,
      default: '4rem'
    },
    thickness: {
      type: String,
      default: '1rem'
    },
    color: {
      type: String,
      default: '#000'
    }
  },
})
</script>
