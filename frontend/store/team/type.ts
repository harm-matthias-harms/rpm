import { Author } from '../user/type'

export interface State {
  team: Team
  teamList: {
    count: number
    teams: Team[]
  }
  teamsLoaded: boolean
}

export interface Team {
  id?: string
  author: Author
  editor: Author
  createdAt?: Date
  editedAt?: Date
  title: string
  type?: string
  medivac?: boolean
}
