import { ActionContext, ActionTree } from 'vuex/types'
import { State } from './type'
import { State as RootState } from '@/store/root'
//import axios from 'axios'

interface UserActionContext extends ActionContext<State, RootState> {}

export const actions: ActionTree<State, RootState> = {
  register({ commit }, user) {
    commit('loader/SET', true, {root: true})
    //axios({ url: '/auth/register' }).then(response => {}, error => {})
    setTimeout(() => {
      commit('loader/SET', false, {root: true})
    }, 2000)
    
  }
}

export default actions
