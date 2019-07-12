import { ActionContext, ActionTree } from 'vuex/types'
import { State } from './type'
import { State as RootState } from '@/store/root'

interface UserActionContext extends ActionContext<State, RootState> {}

export const actions: ActionTree<State, RootState> = {
  register({ state, commit }, user) {
    commit('loader/SET', true, { root: true })
    commit('SET_USER_REGISTER', user)
    this.$axios
      .$post('/auth/register', state.user)
      .then(response => {
        if (response.status === 201) {
          commit('REGISTER_SUCCESS')
        }
        commit('loader/SET', false, { root: true })
        commit('UNSET_USER')
      })
      .catch(error => {
        if (
          error.response.status === 400 &&
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
  }
}

export default actions
