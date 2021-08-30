import { medicalCase } from '../medicalCase/state'
import { team } from '../team/state'
import { author } from '../user/state'
import { State } from './type'

export const defaultInject = {
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
}

export const state = (): State => ({
  inject: defaultInject,
  injectList: {
    count: 0,
    injects: [],
  },
  injectsLoaded: false,
})

export default state
