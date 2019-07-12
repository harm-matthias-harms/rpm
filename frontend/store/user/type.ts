export interface State {
  user: {
    username: string
    email: string
    password: string
  }
  registerError: boolean,
  registerErrorReason: string
  registerSuccess: boolean
}
