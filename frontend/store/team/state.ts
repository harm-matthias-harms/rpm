import { author } from '../user/state'
import { State, Team } from './type'

export const defaultTeam: Team = {
  id: undefined,
  author,
  editor: author,
  createdAt: undefined,
  editedAt: undefined,
  title: '',
  type: undefined,
  medivac: undefined
}

export const state = (): State => ({
  team: defaultTeam,
  teamList: {
    count: 0,
    teams: []
  },
  teamsLoaded: false
})

export default state
