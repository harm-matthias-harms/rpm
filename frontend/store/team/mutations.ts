import { MutationTree } from 'vuex'
import { State } from './type'

export const mutations: MutationTree<State> = {
  SET_TEAM(state, team) {
    state.team = team
  },
  SET_TEAM_LIST(state, list) {
    state.teamList = list
    state.teamsLoaded = true
  },
  SET_TEAM_TO_LIST(state, team) {
    state.teamList.count += 1
    state.teamList.teams.unshift(team)
  },
  DELETE_FROM_LIST(state, id) {
    state.teamList.teams = state.teamList.teams.filter((item) => item.id !== id)
    state.teamList.count--
  },
  UNSET_TEAM(state) {
    state.team = {
      id: undefined,
      author: {
        id: undefined,
        username: undefined,
      },
      editor: {
        id: undefined,
        username: undefined,
      },
      createdAt: undefined,
      editedAt: undefined,
      title: '',
      type: undefined,
      medivac: undefined,
    }
  },
}

export default mutations
