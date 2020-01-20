import { ActionTree } from 'vuex/types'
import { State } from './type'
import { State as RootState } from '@/store/root'

export const actions: ActionTree<State, RootState> = {
  create ({ commit }, preset) {
    commit('loader/SET', true, { root: true })
    this.$axios
      .$post('/api/presets', preset)
      .then((response) => {
        commit('SET_PRESET_TO_LIST', response.data)
        this.$router.push('/presets')
      })
      .catch(() => {
        commit('snackbar/SET', "Couldn't create preset.", { root: true })
      })
      .finally(() => {
        commit('loader/SET', false, { root: true })
      })
  },
  update ({ commit }, preset) {
    commit('loader/SET', true, { root: true })
    this.$axios
      .$put('/api/presets/' + preset.id, preset)
      .then((response) => {
        commit('SET_PRESET', response)
        this.$router.push('/presets/' + preset.id)
      })
      .catch(() => {
        commit('snackbar/SET', "Couldn't edit preset.", { root: true })
      })
      .finally(() => {
        commit('loader/SET', false, { root: true })
      })
  },
  get_all ({ commit }, payload = { disableLoader: false }) {
    if (!payload.disableLoader) {
      commit('loader/SET', true, { root: true })
    }
    this.$axios
      .$get('/api/presets')
      .then((response) => {
        commit('SET_PRESET_LIST', response)
      })
      .catch(() => {
        commit('snackbar/SET', "Couldn't load presets.", { root: true })
      })
      .finally(() => {
        if (!payload.disableLoader) {
          commit('loader/SET', false, { root: true })
        }
      })
  },
  find ({ commit }, payload = { id: null, disableLoader: false }) {
    if (!payload.disableLoader) {
      commit('loader/SET', true, { root: true })
    }
    return this.$axios
      .$get('/api/presets/' + payload.id)
      .then((response) => {
        commit('SET_PRESET', response)
        return response
      })
      .catch(() => {
        commit('snackbar/SET', "Couldn't find preset.", { root: true })
      })
      .finally(() => {
        if (!payload.disableLoader) {
          commit('loader/SET', false, { root: true })
        }
      })
  },
  delete ({ commit }, payload = { id: null, goBack: false }) {
    commit('loader/SET', true, { root: true })
    this.$axios
      .$delete('/api/presets/' + payload.id)
      .then(() => {
        commit('DELETE_FROM_LIST', payload.id)
        commit('snackbar/SET', 'Preset  was successfully deleted.', { root: true })
        if (payload.goBack) {
          this.$router.back()
        }
      })
      .catch(() => {
        commit('snackbar/SET', "Couldn't delete preset.", { root: true })
      })
      .finally(() => {
        commit('loader/SET', false, { root: true })
      })
  },
}

export default actions
