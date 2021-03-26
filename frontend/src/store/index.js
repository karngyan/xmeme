import { createStore } from 'vuex'
import axios from "axios";

export default createStore({
  state: {
    baseURL: '',
    serverConfigured: null,
  },

  mutations: {
    setServer(state, baseURL) {
      if (baseURL) {
        state.baseURL = baseURL
        state.serverConfigured = true
      }
    }
  },

  getters: {
    baseURL(state) {
      return state.baseURL
    },
    serverConfigured(state) {
      return state.serverConfigured
    }
  },

  actions: {
    fetchAllMemes({state}) {
      return new Promise((resolve, reject) => {

        const endpoint = state.baseURL + '/memes'
        axios.get(endpoint)
            .then((resp) => {
              resolve(resp.data)
            })
            .catch((error) => {
              reject(error)
            })
      })
    },

    createMeme({state}, {name, url, caption}) {
      return new Promise((resolve, reject) => {
        const endpoint = state.baseURL + '/memes'
        axios.post(endpoint, {
          name: name,
          url: url,
          caption
        }).then((resp) => {
          resolve(resp.data)
        }).catch((error) => {
          reject(error)
        })
      })
    },

    configureServer({commit}, {baseURL}) {
      return new Promise((resolve) => {
        commit('setServer', baseURL)
        resolve('configured')
      })
    },
  },





})
