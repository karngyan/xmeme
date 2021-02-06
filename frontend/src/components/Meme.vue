<template>
  <li>
    <div class="space-y-4">
      <div class="aspect-w-3 aspect-h-2">
        <img class="object-contain shadow-lg rounded-lg" :src="imgUrl" @error="onImgError" alt="">
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
  </li>
</template>

<script>
import moment from "moment";

export default {
  name: "Meme",
  data() {
    return {
      imgError: false,
    }
  },
  props: {
    name: {
      type: String,
      required: true,
    },
    url: {
      type: String,
      required: true,
    },
    caption: {
      type: String,
      required: true,
    },
    created: {
      type: Number,
      required: true,
    },
    updated: {
      type: Number,
      required: true
    }
  },
  computed: {
    time() {
      return moment.unix(Math.floor(this.created/1e9)).fromNow()
    },
    imgUrl() {
      if (this.imgError) {
        return "https://cdn.sstatic.net/Sites/stackoverflow/img/404.svg"
      }
      return this.url
    }
  },
  methods: {
    onImgError() {
      this.imgError = true
    },
    trimText() {

    }
  }
}
</script>

<style scoped>

</style>