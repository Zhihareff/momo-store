<template>
  <div class="alert fade show" :class="{'alert-dismissible': dismissible}" role="alert">
    <strong>{{modelValue}}</strong>
    <button v-if="dismissible" type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close" @click="handleClick"></button>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  props: {
    modelValue: {
      type: String,
      default: ''
    },
    dismissible: {
      type: Boolean,
      default: true
    }
  },
  emits: [
    'update:modelValue'
  ],
  methods: {
    handleClick() {
      this.$emit('update:modelValue', '')
    }
  }
})
</script>
