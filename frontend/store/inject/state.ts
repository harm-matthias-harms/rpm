import { medicalCase } from '../medicalCase/state'
import { team } from '../team/state'
import { author } from '../user/state'
import { State } from './type'

export const state = (): State => ({
  inject: {
    id: undefined,
    author: author,
    editor: author,
    createdAt: undefined,
    editedAt: undefined,
    exerciseID: undefined,
    status: undefined,
    makeupCenterTitle: undefined,
    team: team,
    roleplayer: {},
    medicalCase: medicalCase,
  },
  injectList: {
    count: 0,
    injects: [],
  },
  injectsLoaded: false,
})

export default state
