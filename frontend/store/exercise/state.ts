import { State } from './type'

export const state = (): State => ({
  exercise: {
    id: undefined,
    author: {
      id: undefined,
      username: undefined
    },
    createdAt: undefined,
    updatedAt: undefined,
    title: '',
    startTime: undefined,
    endTime: undefined,
    teams: [],
    roleplayManager: [],
    makeupCenter: []
  }
})

export default state
