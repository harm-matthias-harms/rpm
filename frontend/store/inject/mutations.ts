import { MutationTree } from 'vuex'
import { defaultInject } from './state'
import { State } from './type'

export const mutations: MutationTree<State> = {
  SET_INJECT(state, inject) {
    state.inject = inject
  },
  SET_INJECT_LIST(state, list) {
    state.injectList = list
    state.injectsLoaded = true
  },
  SET_INJECTS_TO_LIST(state, injects) {
    state.injectList.count += injects.length
    injects.forEach((inject) => {
      state.injectList.injects.unshift(inject)
    })
  },
  DELETE_FROM_LIST(state, id) {
    state.injectList.injects = state.injectList.injects.filter(
      (item) => item.id !== id
    )
    state.injectList.count--
  },
  UNSET_INJECT(state) {
    state.inject = defaultInject
  },
}

export default mutations
