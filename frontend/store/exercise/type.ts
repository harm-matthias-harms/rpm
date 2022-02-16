import { Team } from '../team/type'
import { Author, ShortUser } from '../user/type'

export interface MakeupCenter {
  title: string
  account: ShortUser
}

export interface Exercise {
  id: string
  author: Author
  createdAt: Date | null
  editedAt: Date | null
  title: string
  startTime: string
  endTime: string
  teams: {
    team: Team
    trainer: ShortUser
  }[]
  roleplayManager: ShortUser[]
  makeupCenter: MakeupCenter[]
}

export interface State {
  exercise: Exercise
}
