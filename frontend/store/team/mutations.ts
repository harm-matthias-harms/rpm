import { MutationTree } from 'vuex'
import { State } from './type'

export const mutations: MutationTree<State> = {
  SET_TEAM (state, team) {
    state.team = team
  },
  SET_TEAM_LIST (state, list) {
    state.teamList = list
    state.teamsLoaded = true
  },
  SET_TEAM_TO_LIST (state, team) {
    state.teamList.count += 1
    state.teamList.teams.unshift(team)
  },
  UNSET_TEAM (state) {
    state.team = {
      id: undefined,
      author: {
        id: undefined,
        username: undefined
      },
      editor: {
        id: undefined,
        username: undefined
      },
      createdAt: undefined,
      updatedAt: undefined,
      title: '',
      type: undefined,
      medivac: undefined
    }
  }
}

export default mutations
