<template>
  <div>
    <div class="flex justify-between items-center mb-6">
      <h1 class="text-2xl font-bold">Weight Entries</h1>
      <NuxtLink to="/weight/new" class="bg-indigo-600 text-white px-4 py-2 rounded hover:bg-indigo-700">Enter new weight entry</NuxtLink>
    </div>
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
          <tr v-for="entry in entries" :key="entry.id">
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ entry.weight }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ formatDate(entry.entered_date) }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium space-x-2">
              <NuxtLink :to="`/weight/${entry.id}/edit`" class="text-indigo-600 hover:text-indigo-900">Edit</NuxtLink>
              <button @click="deleteEntry(entry.id)" class="text-red-600 hover:text-red-900">Delete</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup>
definePageMeta({
  middleware: 'auth'
})

const entries = ref([])

const formatDate = (date) => {
  if (!date) return ''
  return new Date(date).toLocaleDateString()
}

const loadData = async () => {
  try {
    const data = await $fetch('/api/weight?all=true')
    entries.value = data.weights || []
  } catch (e) {
    console.error(e)
  }
}

const deleteEntry = async (id) => {
  if (confirm('Are you sure you want to delete this entry?')) {
    await $fetch(`/api/weight/${id}`, { method: 'DELETE' })
    await loadData()
  }
}

onMounted(() => {
  loadData()
})
</script>
