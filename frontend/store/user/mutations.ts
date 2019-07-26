import { MutationTree } from 'vuex'
import { State } from './type'

export const mutations: MutationTree<State> = {
  SET_USER_REGISTER(state, user) {
    state.user = {
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
      username: '',
      email: '',
      password: ''
    }
  },
  REGISTER_SUCCESS(state) {
    state.registerSuccess = true
  }
}

export default mutations
