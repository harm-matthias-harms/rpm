import { ActionTree } from 'vuex/types'
import Cookie from 'js-cookie'
import jwtDecode from 'jwt-decode'
import { State as RootState } from '@/store/root'
import { State } from './type'

export const actions: ActionTree<State, RootState> = {
  register ({ state, commit }, user) {
    commit('loader/SET', true, { root: true })
    commit('SET_USER_REGISTER', user)
    this.$axios
      .$post('/auth/register', state.user)
      .then(() => {
        commit('REGISTER_SUCCESS')
        commit('UNSET_USER')
      })
      .catch((error) => {
        if (
          error.response &&
          error.response.data.message === 'user already exists'
        ) {
          commit('SET_ERROR_REGISTER', 'already exists')
          commit('snackbar/SET', 'Account already exists please login.', {
            root: true
          })
        } else {
          commit('SET_ERROR_REGISTER', 'connection error')
          commit('snackbar/SET', "Couldn't create new account.", { root: true })
        }
      })
      .finally(() => {
        commit('loader/SET', false, { root: true })
      })
  },
  signin ({ commit }, user) {
    commit('loader/SET', true, { root: true })
    this.$axios
      .$post('/auth/authenticate', user)
      .then(() => {
        const cookie = Cookie.get('Authorization')
        if (cookie) {
          const decoded = jwtDecode(cookie)
          commit('SET_AUTHENTICATE', {
            id: decoded.id,
            username: decoded.username,
            expire: decoded.exp,
            code: decoded.code
          })
        } else {
          commit('snackbar/SET', 'Wrong username, password or code', { root: true })
        }
      })
      .catch((error) => {
        if (error.response && error.response.status === 404) {
          commit('snackbar/SET', "Couldn't connect to network", { root: true })
        } else if (error.response && error.response.status === 401) {
          commit('snackbar/SET', 'Wrong username, password or code', { root: true })
        }
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
      .$get('/api/user')
      .then((response) => {
        commit('SET_USER_LIST', response)
      })
      .catch(() => {
        commit('snackbar/SET', "Couldn't load users.", { root: true })
      })
      .finally(() => {
        if (!payload.disableLoader) {
          commit('loader/SET', false, { root: true })
        }
      })
  }
}

export default actions
