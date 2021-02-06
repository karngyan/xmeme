import { createStore } from 'vuex'
import XMemeApi from 'x_meme_api'

export default createStore({
  state: {
    memes: [],
  },
  mutations: {
  },
  modules: {
  },
  actions: {
    fetchAllMemes() {
      return new Promise((resolve, reject) => {
        const api = new XMemeApi.MemesApi()

        api.memeControllerGetAllMemes(function (error, data) {
          if (error) {
            console.debug("some error occurred")
            reject(error)
          } else {
            resolve(data)
          }
        })
      })
    },

    createMeme({state}, {name, url, caption}) {
      return new Promise((resolve, reject) => {
        const api = new XMemeApi.MemesApi()
        console.debug(state.memes)
        api.memeControllerCreateMeme(name, url, caption, function (error, data) {
          if (error) {
            console.error(error)
            reject(error)
          } else {
            resolve(data)
          }
        })
      })
    }
  }
})
