export interface State {
  user: {
    id?: string
    username: string
    email: string
    password: string
  }
  registerError: boolean
  registerErrorReason: string
  registerSuccess: boolean
  isAuthenticated: boolean
  expireSession?: Date
}
