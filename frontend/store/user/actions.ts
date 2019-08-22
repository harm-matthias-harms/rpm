import { ActionContext, ActionTree } from 'vuex/types'
import Cookie from 'js-cookie'
import jwtDecode from 'jwt-decode'
import { State } from './type'
import { State as RootState } from '@/store/root'

interface UserActionContext extends ActionContext<State, RootState> {}

export const actions: ActionTree<State, RootState> = {
  register({ state, commit }, user) {
    commit('loader/SET', true, { root: true })
    commit('SET_USER_REGISTER', user)
    this.$axios
      .$post('/auth/register', state.user)
      .then(() => {
        commit('REGISTER_SUCCESS')
        commit('loader/SET', false, { root: true })
        commit('UNSET_USER')
      })
      .catch(error => {
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
        commit('loader/SET', false, { root: true })
      })
  },
  signin({ commit }, user) {
    commit('loader/SET', true, { root: true })
    this.$axios
      .$post('/auth/authenticate', user)
      .then(response => {
        console.log(response)
        const cookie = Cookie.get('Authorization')
        console.log(cookie)
        if (cookie) {
          const decoded = jwtDecode(cookie)
          commit('SET_AUTHENTICATE', {
            id: decoded.id,
            username: decoded.username,
            expire: decoded.exp
          })
        } else {
          commit('snackbar/SET', 'Wrong username or password', { root: true })
        }
        commit('loader/SET', false, { root: true })
      })
      .catch(error => {
        if (error.response && error.response.status == 404) {
          commit('snackbar/SET', "Couldn't connect to network", { root: true })
        } else if (error.response && error.response.status == 401) {
          commit('snackbar/SET', 'Wrong username or password', { root: true })
        }
        commit('loader/SET', false, { root: true })
      })
  }
}

export default actions
