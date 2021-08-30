import { ActionTree } from 'vuex/types'
import { State as RootState } from '@/store/root'
import { State } from './type'

export const actions: ActionTree<State, RootState> = {
  create({ commit }, payload) {
    commit('loader/SET', true, { root: true })
    this.$axios
      .$post(`/exercises/${payload.exerciseID}/injects`, payload.injects)
      .then((response) => {
        commit('SET_INJECTS_TO_LIST', response.data)
        this.$router.push(`/exercise/${payload.exerciseID}/injects`)
      })
      .catch(() => {
        commit('snackbar/SET', "Couldn't create injects.", { root: true })
      })
      .finally(() => {
        commit('loader/SET', false, { root: true })
      })
  },
  update({ commit }, inject) {
    commit('loader/SET', true, { root: true })
    this.$axios
      .$put(`/exercises/${inject.exerciseID}/injects/${inject.id}`, inject)
      .then((response) => {
        commit('SET_PRESET', response)
        this.$router.push(`/exercise/${inject.exerciseID}/injects/${inject.id}`)
      })
      .catch(() => {
        commit('snackbar/SET', "Couldn't edit preset.", { root: true })
      })
      .finally(() => {
        commit('loader/SET', false, { root: true })
      })
  },
  get_all({ commit }, payload = { exerciseID: null, disableLoader: false }) {
    if (!payload.disableLoader) {
      commit('loader/SET', true, { root: true })
    }
    this.$axios
      .$get(`/exercises/${payload.exerciseID}/injects`)
      .then((response) => {
        commit('SET_INJECT_LIST', response)
      })
      .catch(() => {
        commit('snackbar/SET', "Couldn't load injects.", { root: true })
      })
      .finally(() => {
        if (!payload.disableLoader) {
          commit('loader/SET', false, { root: true })
        }
      })
  },
  find(
    { commit },
    payload = { id: null, exerciseID: null, disableLoader: false }
  ) {
    if (!payload.disableLoader) {
      commit('loader/SET', true, { root: true })
    }
    return this.$axios
      .$get(`/exercises/${payload.exerciseID}/injects/${payload.id}`)
      .then((response) => {
        commit('SET_INJECT', response)
        return response
      })
      .catch(() => {
        commit('snackbar/SET', "Couldn't find inject.", { root: true })
      })
      .finally(() => {
        if (!payload.disableLoader) {
          commit('loader/SET', false, { root: true })
        }
      })
  },
  delete({ commit }, payload = { id: null, exerciseID: null, goBack: false }) {
    commit('loader/SET', true, { root: true })
    this.$axios
      .$delete(`/exercises/${payload.exerciseID}/injects/${payload.id}`)
      .then(() => {
        commit('DELETE_FROM_LIST', payload.id)
        commit('snackbar/SET', 'Inject  was successfully deleted.', {
          root: true,
        })
        if (payload.goBack) {
          this.$router.back()
        }
      })
      .catch(() => {
        commit('snackbar/SET', "Couldn't delete inject.", { root: true })
      })
      .finally(() => {
        commit('loader/SET', false, { root: true })
      })
  },
}

export default actions
