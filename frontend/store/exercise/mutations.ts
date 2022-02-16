import { MutationTree } from 'vuex'
import { defaultExercise } from './state'
import { Exercise, State } from './type'

export const mutations: MutationTree<State> = {
  SET_EXERCISE(state, exercise: Exercise) {
    state.exercise = exercise
  },
  UNSET_EXERCISE(state) {
    state.exercise = defaultExercise
  }
}

export default mutations
