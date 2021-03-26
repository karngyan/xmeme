<template>
  <div class="max-w-7xl h-screen flex items-center justify-center mx-auto px-4 sm:px-6 lg:px-8">
    <div class="max-w-7xl md:w-1/3 sm:w-1/2 mx-auto">
      <div class="mx-auto flex items-center justify-center h-12 w-12 rounded-full bg-indigo-100">
        <svg class="h-6 w-6 text-indigo-600" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
          <path fill-rule="evenodd" d="M2 5a2 2 0 012-2h12a2 2 0 012 2v2a2 2 0 01-2 2H4a2 2 0 01-2-2V5zm14 1a1 1 0 11-2 0 1 1 0 012 0zM2 13a2 2 0 012-2h12a2 2 0 012 2v2a2 2 0 01-2 2H4a2 2 0 01-2-2v-2zm14 1a1 1 0 11-2 0 1 1 0 012 0z" clip-rule="evenodd" />
        </svg>
      </div>
      <div class="mt-3 text-center sm:mt-5">
        <h3 class="text-lg leading-6 font-medium text-gray-900" id="title">
          Add your server's baseURL
        </h3>
        <form @submit.prevent="startXMeme" class="w-full space-y-4">
          <label for="base_url" class="block text-sm font-medium text-gray-700"></label>
          <div class="mt-1 relative rounded-md shadow-sm">
            <input v-model="baseURL" type="text" name="base_url" id="base_url" class="focus:ring-indigo-500 focus:border-indigo-500 block w-full sm:text-sm border-gray-300 rounded-md" placeholder="http://localhost:8081">
          </div>
          <button type="submit" class="w-full inline-flex justify-center rounded-md border border-transparent shadow-sm px-4 py-2 bg-indigo-600 text-base font-medium text-white hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 sm:col-start-2 sm:text-sm">
            <span v-if="!loading">Start X Meme</span>
            <svg v-else class="h-5" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 100 100" preserveAspectRatio="xMidYMid">
              <path d="M10 50A40 40 0 0 0 90 50A40 42 0 0 1 10 50" fill="#ffffff" stroke="none">
                <animateTransform attributeName="transform" type="rotate" dur="1s" repeatCount="indefinite" keyTimes="0;1" values="0 50 51;360 50 51"></animateTransform>
              </path>
            </svg>
          </button>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'PageConfigure',

  data() {
    return {
      baseURL: '',
      loading: false
    }
  },
  methods: {
    startXMeme() {
      this.loading = true
      const xbaseURL = this.baseURL.replace(/\/+$/, "");
      console.debug(xbaseURL)
      this.$store.dispatch('configureServer', { baseURL: xbaseURL })
        .then(() => {
          setTimeout(() => {
            this.loading = false
            this.$router.push({name: 'Home'})
          }, 500)
        })
    }
  },
  created() {
    this.baseURL = this.$store.state.baseURL
  }
}
</script>

<style scoped>

</style>