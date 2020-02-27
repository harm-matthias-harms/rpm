import { Module } from 'vuex'
import { RootState } from '@/store/root'
import { getters } from './getters'
import { actions } from './actions'
import { mutations } from './mutations'
import { State } from './type'
import { state } from './state'

const namespaced: boolean = true

export const preset: Module<State, RootState> = {
  namespaced,
  state,
  getters,
  actions,
  mutations
}
