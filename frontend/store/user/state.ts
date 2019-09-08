import { State } from './type'

export const state = (): State => ({
  user: {
    id: undefined,
    username: '',
    email: '',
    password: ''
  },
  registerError: false,
  registerErrorReason: '',
  registerSuccess: false,
  isAuthenticated: false,
  expireSession: undefined
})

export default state
