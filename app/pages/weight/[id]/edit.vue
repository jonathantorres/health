<template>
  <div class="max-w-2xl mx-auto">
    <h1 class="text-2xl font-bold mb-6">Edit Weight Entry</h1>
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
        <button type="submit" class="bg-indigo-600 text-white px-4 py-2 rounded hover:bg-indigo-700">Update</button>
        <NuxtLink to="/" class="bg-gray-300 text-gray-700 px-4 py-2 rounded hover:bg-gray-400">Cancel</NuxtLink>
      </div>
    </form>
  </div>
</template>

<script setup>
definePageMeta({
  middleware: 'auth'
})

const route = useRoute()
const weight = ref('')
const entered_date = ref('')
const error = ref('')

onMounted(async () => {
  try {
    const data = await $fetch('/api/weight?all=true')
    const entry = data.weights.find(w => w.id == route.params.id)
    if (entry) {
      weight.value = entry.weight
      entered_date.value = entry.entered_date?.split('T')[0] || entry.entered_date
    } else {
      await navigateTo('/')
    }
  } catch (e) {
    await navigateTo('/')
  }
})

const handleSubmit = async () => {
  try {
    error.value = ''
    await $fetch(`/api/weight/${route.params.id}`, {
      method: 'PUT',
      body: {
        weight: parseFloat(weight.value),
        entered_date: entered_date.value
      }
    })
    await navigateTo('/')
  } catch (e: any) {
    error.value = e.data?.message || 'Failed to update'
  }
}
</script>
