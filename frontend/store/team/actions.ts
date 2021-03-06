import { ActionTree } from 'vuex/types'
import { State as RootState } from '@/store/root'
import { State } from './type'

export const actions: ActionTree<State, RootState> = {
  create ({ commit }, team) {
    commit('loader/SET', true, { root: true })
    this.$axios
      .$post('/api/teams', team)
      .then((response) => {
        commit('SET_TEAM_TO_LIST', response.data)
        this.$router.push('/teams/' + response.data.id)
      })
      .catch(() => {
        commit('snackbar/SET', "Couldn't create team.", { root: true })
      })
      .finally(() => {
        commit('loader/SET', false, { root: true })
      })
  },
  update ({ commit }, team) {
    commit('loader/SET', true, { root: true })
    this.$axios
      .$put('/api/teams/' + team.id, team)
      .then((response) => {
        commit('SET_TEAM', response)
        this.$router.push('/teams/' + team.id)
      })
      .catch(() => {
        commit('snackbar/SET', "Couldn't edit team.", { root: true })
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
      .$get('/api/teams')
      .then((response) => {
        commit('SET_TEAM_LIST', response)
      })
      .catch(() => {
        commit('snackbar/SET', "Couldn't load teams.", { root: true })
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
      .$get('/api/teams/' + payload.id)
      .then((response) => {
        commit('SET_TEAM', response)
        return response
      })
      .catch(() => {
        commit('snackbar/SET', "Couldn't find team.", { root: true })
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
      .$delete('/api/teams/' + payload.id)
      .then(() => {
        commit('DELETE_FROM_LIST', payload.id)
        commit('snackbar/SET', 'Team  was successfully deleted.', { root: true })
        if (payload.goBack) {
          this.$router.back()
        }
      })
      .catch(() => {
        commit('snackbar/SET', "Couldn't delete team.", { root: true })
      })
      .finally(() => {
        commit('loader/SET', false, { root: true })
      })
  }
}

export default actions
