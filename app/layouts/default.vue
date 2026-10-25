<template>
  <div class="min-h-screen bg-gray-50">
    <NavBar v-if="user" />
    <main class="py-6">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <slot />
      </div>
    </main>
  </div>
</template>

<script setup>
const { user } = useAuth()
</script>
