import { ActionTree } from 'vuex/types'
import { State as RootState } from '@/store/root'
import { State } from './type'

export const actions: ActionTree<State, RootState> = {
  create({ commit }, payload) {
    commit('loader/SET', true, { root: true })
    this.$axios
      .$post(`/api/exercises/${payload.exerciseID}/injects`, payload.injects)
      .then((response) => {
        commit('SET_INJECTS_TO_LIST', response.data)
        this.$router.push(`/exercises/${payload.exerciseID}/injects`)
      })
      .catch(() => {
        commit('snackbar/SET', "Couldn't create injects.", { root: true })
      })
      .finally(() => {
        commit('loader/SET', false, { root: true })
      })
  },
  update({ commit }, payload = { inject: null, showInject: false }) {
    commit('loader/SET', true, { root: true })
    return this.$axios
      .$put(
        `/api/exercises/${payload.inject.exerciseID}/injects/${payload.inject.id}`,
        payload.inject
      )
      .then((response) => {
        commit('SET_INJECT', response)
        if (payload.showInject) {
          this.$router.push(
            `/exercises/${response.exerciseID}/injects/${response.id}`
          )
        }
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
      .$get(`/api/exercises/${payload.exerciseID}/injects`)
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
      .$get(`/api/exercises/${payload.exerciseID}/injects/${payload.id}`)
      .then((response) => {
        commit('SET_INJECT', response)
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
      .$delete(`/api/exercises/${payload.exerciseID}/injects/${payload.id}`)
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
