export interface State {
  team: {
    id?: string
    author: {
      id?: string
      username?: string
    }
    editor: {
      id?: string
      username?: string
    }
    createdAt?: Date
    updatedAt?: Date
    title: string
    type?: string
    medivac?: boolean
  }
  teamList: {
    count: number
    teams: {
      id?: string
      author: {
        id?: string
        username?: string
      }
      editor: {
        id?: string
        username?: string
      }
      createdAt?: Date
      updatedAt?: Date
      title: string
      type?: string
      medivac?: boolean
    }[]
  }
  teamsLoaded: boolean
}
