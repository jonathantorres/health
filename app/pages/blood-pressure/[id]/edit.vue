<template>
  <div class="max-w-2xl mx-auto">
    <h1 class="text-2xl font-bold mb-6">Edit Blood Pressure Reading</h1>
    <form @submit.prevent="handleSubmit" class="bg-white shadow rounded-lg p-6 space-y-4">
      <div>
        <label class="block text-sm font-medium text-gray-700">SYS</label>
        <input v-model="systolic" type="number" required class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm">
      </div>
      <div>
        <label class="block text-sm font-medium text-gray-700">DIA</label>
        <input v-model="diastolic" type="number" required class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm">
      </div>
      <div>
        <label class="block text-sm font-medium text-gray-700">Pulse</label>
        <input v-model="pulse" type="number" required class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm">
      </div>
      <div>
        <label class="block text-sm font-medium text-gray-700">Reading Date</label>
        <input v-model="reading_date" type="date" required class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm">
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
const systolic = ref('')
const diastolic = ref('')
const pulse = ref('')
const reading_date = ref('')
const error = ref('')

onMounted(async () => {
  try {
    const data = await $fetch(`/api/blood-pressure/${route.params.id}`)
    const reading = data.reading
    systolic.value = reading.systolic
    diastolic.value = reading.diastolic
    pulse.value = reading.pulse
    reading_date.value = reading.reading_date?.split('T')[0] || reading.reading_date
  } catch (e) {
    await navigateTo('/')
  }
})

const handleSubmit = async () => {
  try {
    error.value = ''
    await $fetch(`/api/blood-pressure/${route.params.id}`, {
      method: 'PUT',
      body: {
        systolic: parseInt(systolic.value),
        diastolic: parseInt(diastolic.value),
        pulse: parseInt(pulse.value),
        reading_date: reading_date.value
      }
    })
    await navigateTo('/')
  } catch (e) {
    error.value = e.data?.message || 'Failed to update'
  }
}
</script>
