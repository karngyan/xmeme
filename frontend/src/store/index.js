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
        const body = new XMemeApi.ModelsMeme()
        body.name = name
        body.url = url
        body.caption = caption
        api.memeControllerCreateMeme(body, function (error, data) {
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
