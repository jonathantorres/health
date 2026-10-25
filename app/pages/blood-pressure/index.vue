<template>
  <div>
    <div class="flex justify-between items-center mb-6">
      <h1 class="text-2xl font-bold">Blood Pressure Readings</h1>
      <NuxtLink to="/blood-pressure/new" class="bg-indigo-600 text-white px-4 py-2 rounded hover:bg-indigo-700">Enter new reading</NuxtLink>
    </div>
    <div class="overflow-x-auto bg-white shadow rounded-lg">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">SYS</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">DIA</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Pulse</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Date</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Severity</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Actions</th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-for="reading in readings" :key="reading.id">
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ reading.systolic }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ reading.diastolic }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ reading.pulse }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ formatDate(reading.reading_date) }}</td>
            <td class="px-6 py-4 whitespace-nowrap">
              <span :class="getSeverity(reading.systolic, reading.diastolic).class" class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full">
                {{ getSeverity(reading.systolic, reading.diastolic).text }}
              </span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium space-x-2">
              <NuxtLink :to="`/blood-pressure/${reading.id}`" class="text-blue-600 hover:text-blue-900">View</NuxtLink>
              <NuxtLink :to="`/blood-pressure/${reading.id}/edit`" class="text-indigo-600 hover:text-indigo-900">Edit</NuxtLink>
              <button @click="deleteReading(reading.id)" class="text-red-600 hover:text-red-900">Delete</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup>
import { getBloodPressureSeverity as getSeverity } from '~/utils/bpSeverity'

definePageMeta({
  middleware: 'auth'
})

const readings = ref([])

const formatDate = (date) => {
  if (!date) return ''
  return new Date(date).toLocaleDateString()
}

const loadData = async () => {
  try {
    const data = await $fetch('/api/blood-pressure?all=true')
    readings.value = data.readings || []
  } catch (e) {
    console.error(e)
  }
}

const deleteReading = async (id) => {
  if (confirm('Are you sure you want to delete this reading?')) {
    await $fetch(`/api/blood-pressure/${id}`, { method: 'DELETE' })
    await loadData()
  }
}

onMounted(() => {
  loadData()
})
</script>
