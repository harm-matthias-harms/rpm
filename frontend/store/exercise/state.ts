import { author } from '../user/state'
import { State } from './type'

export const state = (): State => ({
  exercise: {
    id: undefined,
    author: author,
    createdAt: undefined,
    editedAt: undefined,
    title: '',
    startTime: undefined,
    endTime: undefined,
    teams: [],
    roleplayManager: [],
    makeupCenter: [],
  },
})

export default state
