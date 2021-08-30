import { Author, State } from './type'

export const author: Author = {
  id: undefined,
  username: undefined
}

export const defaultUser = {
  id: undefined,
  username: '',
  email: '',
  password: '',
  code: undefined,
  roles: []
}

export const state = (): State => ({
  user: defaultUser,
  registerError: false,
  registerErrorReason: '',
  registerSuccess: false,
  isLoaded: false,
  isAuthenticated: false,
  isCodeUser: false,
  expireSession: undefined,
  userList: []
})

export default state
