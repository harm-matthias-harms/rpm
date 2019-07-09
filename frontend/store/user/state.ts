import { State } from './type'

export const state = (): State => ({
  user: {
    username: '',
    email: '',
    password: ''
  },
  registerError: ''
})

export default state
