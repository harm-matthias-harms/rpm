import { State } from './type'

export const state = (): State => ({
  user: {
    id: undefined,
    username: '',
    email: '',
    password: '',
    code: undefined
  },
  registerError: false,
  registerErrorReason: '',
  registerSuccess: false,
  isAuthenticated: false,
  isCodeUser: false,
  expireSession: undefined,
  userList: []
})

export default state
