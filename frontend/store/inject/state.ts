import { defaultMedicalCase } from '../medicalCase/state'
import { defaultTeam } from '../team/state'
import { author } from '../user/state'
import { Roleplayer, State } from './type'

export const defaultRoleplayer: Roleplayer = {
  gender: undefined,
  fullName: undefined,
  age: undefined,
  nationality: undefined,
}

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
  medicalCase: defaultMedicalCase,
}

export const state = (): State => ({
  inject: defaultInject,
  exerciseID: undefined,
  injectList: {
    count: 0,
    injects: [],
  },
  injectsLoaded: false,
})

export default state
