import { GetterTree } from 'vuex'
import { State as RootState } from '@/store/root'
import { State } from './type'

export const getters: GetterTree<State, RootState> = {
  requireReview: state => () => {
    return state.medicalCasesList.medicalCases.filter(
      medicalCase => !medicalCase.approved
    )
  }
}

export default getters
