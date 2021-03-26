<template>

<div>
  <div class="mx-auto py-6 px-4 max-w-7xl sm:px-6 lg:px-8 lg:py-6">
    <div class="space-y-12">
      <div class="flex flex-col sm:flex-row justify-between">

        <div class="space-y-5 sm:space-y-4 md:max-w-xl lg:max-w-3xl xl:max-w-none">
          <h2 class="text-3xl font-extrabold tracking-tight sm:text-4xl">X Meme</h2>
          <p class="text-xl text-gray-500">{{ serverConfigured ? 'baseURL: ' + baseURL : 'please configure the baseUrl to create memes'}}</p>
        </div>
        <div class="justify-end flex items-end sm:justify-start sm:items-start space-x-2">
          <router-link :to="{name: 'Configure'}" class="inline-flex items-center px-4 py-2 border border-transparent text-base font-medium rounded-md text-indigo-700 bg-indigo-100 hover:bg-indigo-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
            Configure
          </router-link>
          <button @click="sideBarOpen = true" type="button" class="cursor-pointer inline-flex items-center px-4 py-2 border border-transparent text-base font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
            Create Meme
          </button>
        </div>
      </div>

      <ul class="space-y-12 sm:grid sm:grid-cols-2 sm:gap-x-6 sm:gap-y-12 sm:space-y-0 lg:grid-cols-3 lg:gap-x-8">
        <Meme v-for="meme in memes" :key="meme.id" :name="meme.name" :caption="meme.caption" :url="meme.url" :created="meme.created" :updated="meme.updated" />
      </ul>
    </div>
  </div>
</div>

<div v-show="sideBarOpen" class="fixed inset-0 overflow-hidden">
  <div class="absolute inset-0 overflow-hidden">
    <section class="absolute inset-y-0 right-0 pl-10 max-w-full flex sm:pl-16" aria-labelledby="slide-over-heading">
    <transition
      enter-active-class="transform transition ease-in-out duration-500 sm:duration-500"
      enter-from-class="translate-x-full"
      enter-to-class="translate-x-0"
      leave-active-class="transform transition ease-in-out duration-500 sm:duration-500"
      leave-from-class="translate-x-0"
      leave-to-class="translate-x-full"
    >
      <div v-show="sideBarOpen" class="w-screen max-w-2xl">
        <form class="h-full flex flex-col bg-white shadow-xl overflow-y-scroll">
          <div class="flex-1">
            <div class="px-4 py-6 bg-gray-50 sm:px-6">
              <div class="flex items-start justify-between space-x-3">
                <div class="space-y-1">
                  <h2 id="slide-over-heading" class="text-lg font-medium text-gray-900">
                    New X Meme
                  </h2>
                  <p class="text-sm text-gray-500">
                    Get started by filling in the information below to create your new xmeme.
                  </p>
                </div>
                <div class="h-7 flex items-center">
                  <div @click="sideBarOpen = false" class="cursor-pointer bg-white rounded-md text-gray-400 hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-indigo-500">
                    <span class="sr-only">Close panel</span>
                    <!-- Heroicon name: outline/x -->
                    <svg class="h-6 w-6" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                    </svg>
                  </div>
                </div>
              </div>
            </div>

            <div class="py-6 space-y-6 sm:py-0 sm:space-y-0 sm:divide-y sm:divide-gray-200">
              <div class="space-y-1 px-4 sm:space-y-0 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6 sm:py-5">
                <div>
                  <label for="project_name" class="block text-sm font-medium text-gray-900 sm:mt-px sm:pt-2">
                    Meme Owner*
                  </label>
                </div>
                <div class="sm:col-span-2">
                  <input placeholder="Enter your full name" type="text" name="project_name" v-model="meme.name" class="block w-full shadow-sm sm:text-sm focus:ring-indigo-500 focus:border-indigo-500 border-gray-300 rounded-md">
                </div>
              </div>

              <div class="space-y-1 px-4 sm:space-y-0 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6 sm:py-5">
                <div>
                  <label for="project_description" class="block text-sm font-medium text-gray-900 sm:mt-px sm:pt-2">
                    Caption*
                  </label>
                </div>
                <div class="sm:col-span-2">
                  <textarea v-model="meme.caption" placeholder="Be creative with the caption" id="project_description" name="project_description" rows="3" class="block w-full shadow-sm sm:text-sm focus:ring-indigo-500 focus:border-indigo-500 border-gray-300 rounded-md"></textarea>
                </div>
              </div>

              <div class="space-y-1 px-4 sm:space-y-0 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6 sm:py-5">
                <div>
                  <label for="project_name" class="block text-sm font-medium text-gray-900 sm:mt-px sm:pt-2">
                    Meme URL*
                  </label>
                </div>
                <div class="sm:col-span-2">
                  <input v-model="meme.url" placeholder="Enter URL of your meme here" type="text" name="project_name" id="project_name" class="block w-full shadow-sm sm:text-sm focus:ring-indigo-500 focus:border-indigo-500 border-gray-300 rounded-md">
                </div>
              </div>

              <div class="space-y-1 px-4 sm:space-y-0 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6 sm:py-5">
                <div>
                  <label for="project_name" class="block text-sm font-medium text-gray-900 sm:mt-px sm:pt-2">
                    Meme Preview
                  </label>
                </div>
                <div class="sm:col-span-2">
                  <div class="space-y-4">
                    <div class="aspect-w-3 aspect-h-2">
                      <img class="object-contain shadow-lg rounded-lg" :src="url" @error="onImgError" alt="">
                    </div>

                    <div class="space-y-2">
                      <div class="text-lg leading-6 font-medium space-y-1">
                        <h3 class="truncate">{{ caption }}</h3>
                      </div>
                      <div class="flex flex-row truncate justify-between items-center">
                        <p class="truncate text-gray-500">{{ name }}</p>
                        <p class="text-gray-500">{{ time }}</p>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <p class="space-y-1 px-4 sm:space-y-0 sm:px-6 sm:py-5  text-red-700" v-show="showError">{{ errorString }}</p>
          <p class="space-y-1 px-4 sm:space-y-0 sm:px-6 sm:py-5  text-red-700" v-show="showCustomError">{{ customErrorString }}</p>

          <!-- Action buttons -->
          <div class="flex-shrink-0 px-4 border-t border-gray-200 py-5 sm:px-6">
            <div class="space-x-3 flex justify-end">
              <button @click="sideBarOpen=false" type="button" class="bg-white py-2 px-4 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
                Cancel
              </button>
              <button v-show="!showSpinner" @click="submitMeme" type="submit" class="inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
                Create
              </button>
              <div v-show="showSpinner" class="inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-indigo-600 opacity-75">
                Create
              </div>
            </div>
          </div>
        </form>
      </div>
    </transition>
    </section>
  </div>
</div>


</template>

<script>
import Meme from '@/components/Meme';

export default {
  name: "Home",
  components: {Meme},
  data() {
    return {
      memes: [],
      meme: {
        name: '',
        url: '',
        caption: ''
      },

      sideBarOpen: false,
      showSpinner: false,
      errorString: '',
      customErrorString: '',
      showCustomError: false,
    }
  },
  computed: {
    serverConfigured() {
      return this.$store.state.serverConfigured
    },
    baseURL() {
      return this.$store.state.baseURL
    },
    name() {
      return this.meme.name.trim()
    },
    url() {
      return this.meme.url.trim()
    },
    caption() {
      return this.meme.caption.trim()
    },
    showError() {
      return this.name === '' || this.url === '' || this.caption === '';
    }
  },
  methods: {
    onImgError(e) {
      e.target.src = "https://cdn.sstatic.net/Sites/stackoverflow/img/404.svg"
    },
    submitMeme(e) {
      e.preventDefault()
      this.showSpinner = true
      if (this.showError) {
        this.errorString = 'Make sure all fields are filled up.'
        return
      }
      this.errorString = ''

      const store = this.$store
      store.dispatch("createMeme", {name: this.name, caption: this.caption, url: this.url})
        .then(() => {

          store.dispatch("fetchAllMemes")
            .then(memes => {
              this.memes = memes
              this.showSpinner = false
              this.sideBarOpen = false
            })
        })
        .catch(e => {
          this.customErrorString = e.toString()
          this.showCustomError = true
          this.showSpinner = false
        })

    }
  },
  created() {
    console.debug('all set')
    const store = this.$store

    if (!store.state.serverConfigured) {
      this.$router.push('/configure')
      return
    }

    store.dispatch("fetchAllMemes")
      .then(memes => {
        this.memes = memes
      })
  }
}
</script>

<style scoped>

</style>