import { State } from './type'

export const state = (): State => ({
  user: {
    username: '',
    email: '',
    password: ''
  },
  registerError: false,
  registerErrorReason: '',
  registerSuccess: false
})

export default state
