import { MutationTree } from 'vuex'
import Cookie from 'js-cookie'
import { State } from './type'

export const mutations: MutationTree<State> = {
  SET_USER_REGISTER (state, user) {
    state.user = {
      id: undefined,
      username: user.username,
      email: user.email,
      password: user.password,
      code: undefined
    }
  },
  SET_ERROR_REGISTER (state, reason: string) {
    state.registerError = true
    state.registerErrorReason = reason
  },
  UNSET_USER (state) {
    state.user = {
      id: undefined,
      username: '',
      email: '',
      password: '',
      code: undefined
    }
  },
  REGISTER_SUCCESS (state) {
    state.registerSuccess = true
  },
  SET_AUTHENTICATE (state, { id, username, expire, code }) {
    state.user.id = id
    state.user.username = username
    state.isAuthenticated = true
    if (code) {
      state.isCodeUser = true
    } else {
      state.isCodeUser = false
    }
    state.expireSession = expire
  },
  LOGOUT (state) {
    state.user.id = undefined
    state.user.username = ''
    state.isAuthenticated = false
    state.isCodeUser = false
    state.expireSession = undefined
    Cookie.remove('Authorization')
  },
  SET_USER_LIST( state, userList) {
    state.userList = userList
  }
}

export default mutations
