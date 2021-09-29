import { Team } from '../team/type'
import { MedicalCase, MedicalCaseShort } from '../medicalCase/type'
import { Author } from '../user/type'

export interface Inject {
  id?: string
  author: Author | null
  editor: Author | null
  createdAt?: Date
  editedAt?: Date
  exerciseID?: string
  status?: string
  startTime?: Date | null
  makeupCenterTitle?: string | null
  team: Team | null
  roleplayer: Roleplayer
  medicalCase: MedicalCase
}

export interface State {
  inject: Inject
  injectList: {
    count: number
    injects: {
      id?: string
      author: Author
      editor: Author
      createdAt?: Date
      editedAt?: Date
      exerciseID?: string
      status?: string
      startTime?: Date
      makeupCenterTitle?: string
      team: Team
      roleplayer: Roleplayer
      medicalCase: MedicalCaseShort
    }[]
  }
  injectsLoaded: boolean
}

export interface Roleplayer {
  fullName?: string
  gender?: string
  age?: number
  nationality?: string
}
