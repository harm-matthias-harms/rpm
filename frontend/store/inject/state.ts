import { defaultMedicalCase } from '../medicalCase/state'
import { defaultTeam } from '../team/state'
import { author } from '../user/state'
import { State } from './type'

export const defaultInject = {
  id: undefined,
  author,
  editor: author,
  createdAt: undefined,
  editedAt: undefined,
  exerciseID: undefined,
  status: undefined,
  startTime: undefined,
  makeupCenterTitle: undefined,
  team: defaultTeam,
  roleplayer: {},
  medicalCase: defaultMedicalCase
}

export const state = (): State => ({
  inject: defaultInject,
  injectList: {
    count: 0,
    injects: []
  },
  injectsLoaded: false
})

export default state
