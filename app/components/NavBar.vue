<template>
  <nav class="bg-white shadow">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex justify-between h-16">
        <div class="flex">
          <div class="flex-shrink-0 flex items-center">
            <NuxtLink to="/" class="text-xl font-bold text-gray-800">Health</NuxtLink>
          </div>
          <div class="hidden sm:ml-6 sm:flex sm:space-x-8">
            <!-- Blood Pressure Dropdown -->
            <div class="relative" @mouseenter="bpOpen = true" @mouseleave="bpOpen = false">
              <button class="inline-flex items-center px-1 pt-1 text-sm font-medium text-gray-900">
                Blood Pressure
                <svg class="ml-2 h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                  <path fill-rule="evenodd" d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z" clip-rule="evenodd" />
                </svg>
              </button>
              <div v-show="bpOpen" class="absolute z-10 mt-2 w-48 rounded-md shadow-lg bg-white ring-1 ring-black ring-opacity-5">
                <div class="py-1">
                  <NuxtLink to="/blood-pressure" class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100">See all readings</NuxtLink>
                  <NuxtLink to="/blood-pressure/new" class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100">Enter new reading</NuxtLink>
                </div>
              </div>
            </div>
            <!-- Weight Dropdown -->
            <div class="relative" @mouseenter="weightOpen = true" @mouseleave="weightOpen = false">
              <button class="inline-flex items-center px-1 pt-1 text-sm font-medium text-gray-900">
                Weight
                <svg class="ml-2 h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                  <path fill-rule="evenodd" d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z" clip-rule="evenodd" />
                </svg>
              </button>
              <div v-show="weightOpen" class="absolute z-10 mt-2 w-48 rounded-md shadow-lg bg-white ring-1 ring-black ring-opacity-5">
                <div class="py-1">
                  <NuxtLink to="/weight" class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100">See all entries</NuxtLink>
                  <NuxtLink to="/weight/new" class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100">Enter new weight entry</NuxtLink>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div class="hidden sm:ml-6 sm:flex sm:items-center">
          <div class="ml-3 relative" @mouseenter="userOpen = true" @mouseleave="userOpen = false">
            <button class="flex items-center text-sm font-medium text-gray-900">
              {{ user?.name }} {{ user?.last_name }}
              <svg class="ml-2 h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z" clip-rule="evenodd" />
              </svg>
            </button>
            <div v-show="userOpen" class="absolute right-0 z-10 mt-2 w-48 rounded-md shadow-lg bg-white ring-1 ring-black ring-opacity-5">
              <div class="py-1">
                <button @click="handleLogout" class="block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100">Logout</button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </nav>
</template>

<script setup>
const { user, logout } = useAuth()
const bpOpen = ref(false)
const weightOpen = ref(false)
const userOpen = ref(false)

const handleLogout = () => {
  logout()
}
</script>
