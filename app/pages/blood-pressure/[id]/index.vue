<template>
  <div class="max-w-2xl mx-auto">
    <h1 class="text-2xl font-bold mb-6">Blood Pressure Reading Details</h1>
    <div class="bg-white shadow rounded-lg p-6 space-y-4">
      <div>
        <span class="font-medium">SYS:</span> {{ reading.systolic }}
      </div>
      <div>
        <span class="font-medium">DIA:</span> {{ reading.diastolic }}
      </div>
      <div>
        <span class="font-medium">Pulse:</span> {{ reading.pulse }}
      </div>
      <div>
        <span class="font-medium">Reading Date:</span> {{ formatDate(reading.reading_date) }}
      </div>
      <div>
        <span class="font-medium">Severity:</span>
        <span :class="getSeverity(reading.systolic, reading.diastolic).class" class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full ml-2">
          {{ getSeverity(reading.systolic, reading.diastolic).text }}
        </span>
      </div>
      <div class="flex space-x-4 pt-4">
        <NuxtLink :to="`/blood-pressure/${reading.id}/edit`" class="bg-indigo-600 text-white px-4 py-2 rounded hover:bg-indigo-700">Edit</NuxtLink>
        <NuxtLink to="/" class="bg-gray-300 text-gray-700 px-4 py-2 rounded hover:bg-gray-400">Back to Dashboard</NuxtLink>
      </div>
    </div>
  </div>
</template>

<script setup>
import { getBloodPressureSeverity as getSeverity } from '~/utils/bpSeverity'

definePageMeta({
  middleware: 'auth'
})

const route = useRoute()
const reading = ref({})

const formatDate = (date) => {
  if (!date) return ''
  return new Date(date).toLocaleDateString()
}

onMounted(async () => {
  try {
    const data = await $fetch(`/api/blood-pressure/${route.params.id}`)
    reading.value = data.reading
  } catch (e) {
    await navigateTo('/')
  }
})
</script>
