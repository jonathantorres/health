<template>
  <div class="max-w-2xl mx-auto">
    <h1 class="text-2xl font-bold mb-6">Add Weight Entry</h1>
    <form @submit.prevent="handleSubmit" class="bg-white shadow rounded-lg p-6 space-y-4">
      <div>
        <label class="block text-sm font-medium text-gray-700">Weight</label>
        <input v-model="weight" type="number" step="0.1" required class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm">
      </div>
      <div>
        <label class="block text-sm font-medium text-gray-700">Entry Date</label>
        <input v-model="entered_date" type="date" required class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm">
      </div>
      <div v-if="error" class="text-red-600 text-sm">{{ error }}</div>
      <div class="flex space-x-4">
        <button type="submit" class="bg-indigo-600 text-white px-4 py-2 rounded hover:bg-indigo-700">Save</button>
        <NuxtLink to="/" class="bg-gray-300 text-gray-700 px-4 py-2 rounded hover:bg-gray-400">Cancel</NuxtLink>
      </div>
    </form>
  </div>
</template>

<script setup>
definePageMeta({
  middleware: 'auth'
})

const weight = ref('')
const entered_date = ref(new Date().toISOString().split('T')[0])
const error = ref('')

const handleSubmit = async () => {
  try {
    error.value = ''
    await $fetch('/api/weight', {
      method: 'POST',
      body: {
        weight: parseFloat(weight.value),
        entered_date: entered_date.value
      }
    })
    await navigateTo('/')
  } catch (e) {
    error.value = e.data?.message || 'Failed to save'
  }
}
</script>
