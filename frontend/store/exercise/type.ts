import { Team } from '../team/type'
import { Author, ShortUser } from '../user/type'

export interface State {
  exercise: {
    id?: string
    author: Author
    createdAt?: Date
    editedAt?: Date
    title: string
    startTime?: Date
    endTime?: Date
    teams: {
      team: Team
      trainer: ShortUser
    }[]
    roleplayManager: ShortUser[]
    makeupCenter: {
      title: string
      account: ShortUser
    }[]
  }
}
