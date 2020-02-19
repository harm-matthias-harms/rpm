import Vuex from 'vuex'
import Vue from 'vue'
import * as root from '@/store/root'
import * as loader from '@/store/loader'
import * as snackbar from '@/store/snackbar'
import { user } from '@/store/user'
import { preset } from '@/store/preset'
import { medicalCase } from '@/store/medicalCase'
import { team } from '@/store/team'
Vue.use(Vuex)

// mocking axios methods to prevent axios firing
preset.actions.get_all = () => {}
preset.actions.find = () => {}
preset.actions.delete = () => {}
medicalCase.actions.get_all = () => {}
medicalCase.actions.find = () => {}
medicalCase.actions.delete = () => {}
team.actions.get_all = () => {}
team.actions.find = () => {}
team.actions.delete = () => {}

export const store = new Vuex.Store({
  state: root.state,
  getters: root.getters,
  mutations: root.mutations,
  actions: root.actions,
  modules: {
    loader: Object.assign({ namespaced: true }, loader),
    snackbar: Object.assign({ namespaced: true }, snackbar),
    user,
    preset,
    medicalCase,
    team
  }
})
