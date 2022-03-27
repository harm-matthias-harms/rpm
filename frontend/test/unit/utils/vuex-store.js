import Vuex from 'vuex'
import Vue from 'vue'
import * as root from '@/store/root'
import * as loader from '@/store/loader'
import * as snackbar from '@/store/snackbar'
import { user } from '@/store/user'
import { medicalCase } from '@/store/medicalCase'
import { team } from '@/store/team'
import { exercise } from '@/store/exercise'
import { inject } from '@/store/inject'
Vue.use(Vuex)

// mocking axios methods to prevent axios firing
medicalCase.actions.get_all = () => {}
medicalCase.actions.find = () => {}
medicalCase.actions.delete = () => {}
team.actions.get_all = () => {}
team.actions.find = () => {}
team.actions.delete = () => {}
exercise.actions.find = () => {}
exercise.actions.delete = () => {}
inject.actions.get_all = () => {}
inject.actions.find = () => {}
inject.actions.delete = () => {}
user.actions.get_all = () => {}

export const store = new Vuex.Store({
  state: root.state,
  getters: root.getters,
  mutations: root.mutations,
  actions: root.actions,
  modules: {
    loader: Object.assign({ namespaced: true }, loader),
    snackbar: Object.assign({ namespaced: true }, snackbar),
    user,
    medicalCase,
    team,
    exercise,
    inject
  }
})
