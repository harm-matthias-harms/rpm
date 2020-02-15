import { Module } from 'vuex'
import { getters } from './getters'
import { actions } from './actions'
import { mutations } from './mutations'
import { State } from './type'
import { state } from './state'
import { RootState } from '@/store/root'

const namespaced: boolean = true

export const team: Module<State, RootState> = {
  namespaced,
  state,
  getters,
  actions,
  mutations
}
