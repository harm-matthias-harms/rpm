export interface State {
  exercise: {
    id?: string
    author: {
      id?: string
      username?: string
    }
    createdAt?: Date
    updatedAt?: Date
    title: string
    startTime?: Date
    endTime?: Date
    teams: {
      team: {
        id?: string
        author: {
          id?: string
          username?: string
          code?: string
        }
        editor: {
          id?: string
          username?: string
          code?: string
        }
        createdAt?: Date
        updatedAt?: Date
        title: string
        type?: string
        medivac?: boolean
      }
      trainer: {
        id?: string
        username?: string
        code?: string
      }
    }[]
    roleplayManager: {
      id?: string
      username?: string
      code?: string
    }[]
    makeupCenter: {
      title: string
      account: {
        id?: string
        username?: string
        code?: string
      }
    }[]
  }
}
