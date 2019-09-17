import { ActionTree } from 'vuex/types'
import { State } from './type'
import { State as RootState } from '@/store/root'

export const actions: ActionTree<State, RootState> = {
  create ({ commit }, preset) {
    commit('loader/SET', true, { root: true })
    this.$axios
      .$post('/api/presets', preset)
      .then((response) => {
        commit('SET_PRESET', response.data)
      })
      .catch(() => {
        commit('snackbar/SET', "Couldn't create preset.", { root: true })
      })
      .finally(() => {
        commit('loader/SET', false, { root: true })
      })
  },
  get_all ({ commit }) {
    commit('loader/SET', true, { root: true })
    this.$axios
      .$get('/api/presets')
      .then((response) => {
        commit('SET_PRESET_LIST', response)
      })
      .catch(() => {
        commit('snackbar/SET', "Couldn't load presets.", { root: true })
      })
      .finally(() => {
        commit('loader/SET', false, { root: true })
      })
  },
  find ({ commit }, id) {
    commit('loader/SET', true, { root: true })
    this.$axios
      .$get('/api/presets/' + id)
      .then((response) => {
        commit('SET_PRESET', response)
      })
      .catch(() => {
        commit('snackbar/SET', "Couldn't find preset.", { root: true })
      })
      .finally(() => {
        commit('loader/SET', false, { root: true })
      })
  }
}

export default actions
