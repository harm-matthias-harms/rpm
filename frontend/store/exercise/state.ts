import { cloneDeep } from 'lodash'

import { author } from '../user/state'
import { Exercise, State } from './type'

export const defaultExercise: Exercise = {
  id: '',
  author: cloneDeep(author),
  createdAt: null,
  editedAt: null,
  title: '',
  startTime: '',
  endTime: '',
  teams: [],
  roleplayManager: [],
  makeupCenter: []
}

export const state = (): State => ({
  exercise: defaultExercise
})

export default state
