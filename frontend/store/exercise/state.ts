import { author } from '../user/state'
import { State } from './type'

export const defaultExercise = {
  id: undefined,
  author,
  createdAt: undefined,
  editedAt: undefined,
  title: '',
  startTime: undefined,
  endTime: undefined,
  teams: [],
  roleplayManager: [],
  makeupCenter: []
}

export const state = (): State => ({
  exercise: defaultExercise
})

export default state
