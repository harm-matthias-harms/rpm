import { MutationTree } from 'vuex'
import { State } from './type'
import Cookie from 'js-cookie'

export const mutations: MutationTree<State> = {
  SET_USER_REGISTER(state, user) {
    state.user = {
      id: undefined,
      username: user.username,
      email: user.email,
      password: user.password
    }
  },
  SET_ERROR_REGISTER(state, reason: string) {
    state.registerError = true
    state.registerErrorReason = reason
  },
  UNSET_USER(state) {
    state.user = {
      id: undefined,
      username: '',
      email: '',
      password: ''
    }
  },
  REGISTER_SUCCESS(state) {
    state.registerSuccess = true
  },
  SET_AUTHENTICATE(state, { id, username, expire }) {
    state.user.id = id
    state.user.username = username
    state.isAuthenticated = true
    state.expireSession = expire
  },
  LOGOUT(state) {
    state.user.id = undefined
    state.user.username = ''
    state.isAuthenticated = false
    state.expireSession = undefined
    Cookie.remove('Authorization')
  }
}

export default mutations
