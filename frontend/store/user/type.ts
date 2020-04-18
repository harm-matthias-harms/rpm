export interface State {
  user: {
    id?: string
    username: string
    email: string
    password: string
    code?: string
  }
  registerError: boolean
  registerErrorReason: string
  registerSuccess: boolean
  isAuthenticated: boolean
  isCodeUser: boolean
  expireSession?: Date
  userList: {
    is?: string
    username: string
    email: string
    code?: string
  }[]
}
