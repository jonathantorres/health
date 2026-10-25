<template>
  <div>
    <h1 class="text-2xl font-bold mb-6">Dashboard</h1>
    <div class="mb-8">
      <h2 class="text-xl font-semibold mb-4">Latest Blood Pressure Readings</h2>
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
            <tr v-for="reading in bpReadings" :key="reading.id">
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
                <button @click="deleteBP(reading.id)" class="text-red-600 hover:text-red-900">Delete</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
    <div>
      <h2 class="text-xl font-semibold mb-4">Latest Weight Entries</h2>
      <div class="overflow-x-auto bg-white shadow rounded-lg">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Weight</th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Date</th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Actions</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="entry in weightEntries" :key="entry.id">
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ entry.weight }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ formatDate(entry.entered_date) }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium space-x-2">
                <NuxtLink :to="`/weight/${entry.id}/edit`" class="text-indigo-600 hover:text-indigo-900">Edit</NuxtLink>
                <button @click="deleteWeight(entry.id)" class="text-red-600 hover:text-red-900">Delete</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { getBloodPressureSeverity as getSeverity } from '~/utils/bpSeverity'

definePageMeta({
  middleware: 'auth'
})

const bpReadings = ref<any[]>([])
const weightEntries = ref<any[]>([])

const formatDate = (date: any) => {
  if (!date) return ''
  return new Date(date).toLocaleDateString()
}

const loadData = async () => {
  try {
    const bp: any = await $fetch('/api/blood-pressure?limit=10')
    const weight: any = await $fetch('/api/weight?limit=10')
    bpReadings.value = bp.readings || []
    weightEntries.value = weight.weights || []
  } catch (e) {
    console.error(e)
  }
}

const deleteBP = async (id: number) => {
  if (confirm('Are you sure you want to delete this reading?')) {
    await $fetch(`/api/blood-pressure/${id}`, { method: 'DELETE' })
    await loadData()
  }
}

const deleteWeight = async (id: number) => {
  if (confirm('Are you sure you want to delete this entry?')) {
    await $fetch(`/api/weight/${id}`, { method: 'DELETE' })
    await loadData()
  }
}

onMounted(() => {
  loadData()
})
</script>
