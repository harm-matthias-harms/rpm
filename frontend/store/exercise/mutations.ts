import { MutationTree } from 'vuex'
import { defaultExercise } from './state'
import { State } from './type'

export const mutations: MutationTree<State> = {
  SET_EXERCISE (state, exercise) {
    state.exercise = exercise
  },
  UNSET_EXERCISE (state) {
    state.exercise = defaultExercise
  }
}

export default mutations
