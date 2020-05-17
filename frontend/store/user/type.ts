export interface State {
  user: {
    id?: string
    username: string
    email: string
    password: string
    code?: string
    roles?: {
      role: string
      exercise: {
        id: string
        title: string
        startTime: Date
        endTime: Date
      }
    }[]
  }
  registerError: boolean
  registerErrorReason: string
  registerSuccess: boolean
  isAuthenticated: boolean
  isLoaded: boolean
  isCodeUser: boolean
  expireSession?: Date
  userList: {
    is?: string
    username: string
    email: string
    code?: string
  }[]
}
