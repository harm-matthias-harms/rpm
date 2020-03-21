import { ActionTree } from 'vuex/types'
import { State as RootState } from '@/store/root'
import { State } from './type'

export const actions: ActionTree<State, RootState> = {
  create ({ commit }, exercise) {
    commit('loader/SET', true, { root: true })
    this.$axios
      .$post('/api/exercises', exercise)
      .then((response) => {
        commit('SET_EXERCISE', response.data)
        this.$router.push('/exercises/' + response.data.id)
      })
      .catch(() => {
        commit('snackbar/SET', "Couldn't create exercise.", { root: true })
      })
      .finally(() => {
        commit('loader/SET', false, { root: true })
      })
  },
  update ({ commit }, exercise) {
    commit('loader/SET', true, { root: true })
    this.$axios
      .$put('/api/exercises/' + exercise.id, exercise)
      .then((response) => {
        commit('SET_EXERCISE', response)
        this.$router.push('/exercises/' + exercise.id)
      })
      .catch(() => {
        commit('snackbar/SET', "Couldn't edit exercise.", { root: true })
      })
      .finally(() => {
        commit('loader/SET', false, { root: true })
      })
  },
  find ({ commit }, payload = { id: null, disableLoader: false }) {
    if (!payload.disableLoader) {
      commit('loader/SET', true, { root: true })
    }
    return this.$axios
      .$get('/api/exercises/' + payload.id)
      .then((response) => {
        commit('SET_EXERCISE', response)
        return response
      })
      .catch(() => {
        commit('snackbar/SET', "Couldn't find exercise.", { root: true })
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
      .$delete('/api/exercises/' + payload.id)
      .then(() => {
        commit('UNSET_EXERCISE', payload.id)
        commit('snackbar/SET', 'Exercise was successfully deleted.', { root: true })
        if (payload.goBack) {
          this.$router.back()
        }
      })
      .catch(() => {
        commit('snackbar/SET', "Couldn't delete exercise.", { root: true })
      })
      .finally(() => {
        commit('loader/SET', false, { root: true })
      })
  }
}

export default actions
