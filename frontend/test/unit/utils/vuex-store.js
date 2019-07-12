import Vuex from 'vuex'
import Vue from 'vue'
import * as root from '@/store/root'
import * as loader from '@/store/loader'
import * as snackbar from '@/store/snackbar'
import { user } from '@/store/user'
Vue.use(Vuex)

export const store = new Vuex.Store({
  state: root.state,
  getters: root.getters,
  mutations: root.mutations,
  actions: root.actions,
  modules: {
    loader: Object.assign({ namespaced: true }, loader),
    snackbar: Object.assign({ namespaced: true }, snackbar),
    user
  }
})
