import { GetterTree, ActionTree, MutationTree } from 'vuex'
import { State as RootState } from '@/store/root'

export const name = 'snackbar'

export interface State {
  message: string
}

export const state = (): State => ({
  message: '',
})

export const getters: GetterTree<State, RootState> = {}

export interface Actions<S, R> extends ActionTree<S, R> {}

export const actions: Actions<State, RootState> = {}

export const mutations: MutationTree<State> = {
  SET (state, message: string) {
    state.message = message
  },
}
