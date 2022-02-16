import { ActionTree } from 'vuex/types'
import { State as RootState } from '@/store/root'
import { Exercise, State } from './type'

export const actions: ActionTree<State, RootState> = {
  create({ commit, dispatch }, exercise: Exercise) {
    commit('loader/SET', true, { root: true })
    exercise.startTime = exercise.startTime.slice(0, 10) + 'T00:00:00.000Z'
    exercise.endTime = exercise.endTime.slice(0, 10) + 'T23:59:59.000Z'
    this.$axios
      .$post('/api/exercises', exercise)
      .then((response) => {
        commit('SET_EXERCISE', response.data)
        dispatch('user/load', undefined, { root: true })
        this.$router.push('/exercises/' + response.data.id)
      })
      .catch(() => {
        commit('snackbar/SET', "Couldn't create exercise.", { root: true })
      })
      .finally(() => {
        commit('loader/SET', false, { root: true })
      })
  },
  update({ commit, dispatch }, exercise: Exercise) {
    commit('loader/SET', true, { root: true })
    exercise.startTime = exercise.startTime.slice(0, 10) + 'T00:00:00.000Z'
    exercise.endTime = exercise.endTime.slice(0, 10) + 'T23:59:59.000Z'
    this.$axios
      .$put('/api/exercises/' + exercise.id, exercise)
      .then((response) => {
        commit('SET_EXERCISE', response)
        dispatch('user/load', undefined, { root: true })
        this.$router.push('/exercises/' + exercise.id)
      })
      .catch(() => {
        commit('snackbar/SET', "Couldn't edit exercise.", { root: true })
      })
      .finally(() => {
        commit('loader/SET', false, { root: true })
      })
  },
  find({ commit }, payload = { id: null, disableLoader: false }) {
    if (!payload.disableLoader) {
      commit('loader/SET', true, { root: true })
    }
    return this.$axios
      .$get('/api/exercises/' + payload.id)
      .then((response) => {
        const exercise = response
        exercise.startTime = exercise.startTime.slice(0, 10)
        exercise.endTime = exercise.endTime.slice(0, 10)
        commit('SET_EXERCISE', exercise)
        return exercise
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
  delete({ commit, dispatch }, payload = { id: null }) {
    commit('loader/SET', true, { root: true })
    this.$axios
      .$delete('/api/exercises/' + payload.id)
      .then(() => {
        commit('UNSET_EXERCISE')
        commit('snackbar/SET', 'Exercise was successfully deleted.', {
          root: true
        })
        dispatch('user/load', undefined, { root: true })
        this.$router.push('/')
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
