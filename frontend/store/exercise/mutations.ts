import { MutationTree } from 'vuex'
import { State } from './type'

export const mutations: MutationTree<State> = {
  SET_EXERCISE(state, exercise) {
    state.exercise = exercise
  },
  UNSET_EXERCISE(state) {
    state.exercise = {
      id: undefined,
      author: {
        id: undefined,
        username: undefined
      },
      createdAt: undefined,
      updatedAt: undefined,
      title: '',
      startTime: undefined,
      endTime: undefined,
      teams: [],
      roleplayManager: [],
      makeupCenter: []
    }
  }
}

export default mutations
