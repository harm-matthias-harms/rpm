import { GetterTree } from 'vuex'
import { State } from './type'
import { State as RootState } from '@/store/root'

export const getters: GetterTree<State, RootState> = {
  requireReview: state => () => {
    return state.medicalCasesList.medicalCases.filter(
      medicalCase => !medicalCase.approved
    )
  }
}

export default getters
