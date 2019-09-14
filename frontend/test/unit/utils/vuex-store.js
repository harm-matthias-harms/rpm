import Vuex from 'vuex'
import Vue from 'vue'
import * as root from '@/store/root'
import * as loader from '@/store/loader'
import * as snackbar from '@/store/snackbar'
import { user } from '@/store/user'
import { preset } from '@/store/preset'
Vue.use(Vuex)

// mocking axios methods to prevent axios firing
preset.actions.get_all = () => {}

export const store = new Vuex.Store({
  state: root.state,
  getters: root.getters,
  mutations: root.mutations,
  actions: root.actions,
  modules: {
    loader: Object.assign({ namespaced: true }, loader),
    snackbar: Object.assign({ namespaced: true }, snackbar),
    user,
    preset
  }
})
