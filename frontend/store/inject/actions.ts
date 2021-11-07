import { ActionTree } from 'vuex/types'
import { State as RootState } from '@/store/root'
import { State } from './type'

export const actions: ActionTree<State, RootState> = {
  create({ commit }, payload) {
    commit('loader/SET', true, { root: true })
    return this.$axios
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
  },
  async get_all({ commit, dispatch, rootState }, payload = { exerciseID: null }) {
    if (payload.exerciseID !== rootState.exercise.exercise.id) {
      await dispatch('exercise/find', {id: payload.exerciseID, disableLoader: true}, { root: true })
    }
    const params = filterOptions(rootState, payload.exerciseID)
    return this.$axios
      .$get(`/api/exercises/${payload.exerciseID}/injects`, { params })
      .then((response) => {
        commit('SET_INJECT_LIST', {
          list: response,
          exerciseID: payload.exerciseID,
        })
      })
      .catch(() => {
        commit('snackbar/SET', "Couldn't load injects.", { root: true })
      })
  },
  find(
    { commit },
    payload = { id: null, exerciseID: null }
  ) {
    return this.$axios
      .$get(`/api/exercises/${payload.exerciseID}/injects/${payload.id}`)
      .then((response) => {
        commit('SET_INJECT', response)
      })
      .catch(() => {
        commit('snackbar/SET', "Couldn't find inject.", { root: true })
      })
  },
  delete({ commit }, payload = { id: null, exerciseID: null, goBack: false }) {
    return this.$axios
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
  },
}

function filterOptions(rootState: RootState, exerciseID: string): object {
  const userID = rootState.user.user.id
  const roles = rootState.user.user.roles.filter((role) => role.exercise.id === exerciseID)
  if (roles.some((role) => role.role === 'admin' || role.role === 'role play manager')) return {}
  if (roles.some((role) => role.role === 'make-up center')) {
    const makeupcenter = rootState.exercise.exercise.makeupCenter.filter((center) => center.account.id === userID)[0]
    return { makeupCenterTitle: makeupcenter.title }
  }
  if (roles.some((role) => role.role === 'trainer')) {
    console.log(rootState.exercise.exercise.teams)
    const team = rootState.exercise.exercise.teams.filter((team) => team.trainer.id === userID)[0]
    return { teamId: team.team.id }
  }
  return {}
}

export default actions
