import { GetterTree, ActionTree, MutationTree } from 'vuex'
import { State as RootState } from '@/store/root'

export const name = 'loader'

export interface State {
  isLoading: boolean
}

export const state = (): State => ({
  isLoading: false
})

export const getters: GetterTree<State, RootState> = {}

export interface Actions<S, R> extends ActionTree<S, R> {}

export const actions: Actions<State, RootState> = {}

export const mutations: MutationTree<State> = {
  SET(state, isLoading: boolean) {
    state.isLoading = isLoading
  }
}
