import { State } from './type'

export const state = (): State => ({
  team: {
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
  },
  teamList: {
    count: 0,
    teams: []
  },
  teamsLoaded: false
})

export default state
