import { GetterTree } from 'vuex'
import { State } from './type'
import { State as RootState } from '@/store/root'

export const getters: GetterTree<State, RootState> = {
  myOwn: state => username => {
    return state.medicalCasesList.medicalCases.filter(
      medicalCase => medicalCase.author.username === username
    )
  },
  myOwnCount: (_, getters) => username => {
    return getters.myOwn(username).length
  }
}

export default getters
