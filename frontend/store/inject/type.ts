import { Team } from '../team/type'
import { MedicalCase, MedicalCaseShort } from '../medicalCase/type'
import { Author } from '../user/type'

export interface State {
  inject: {
    id?: string
    author: Author
    editor: Author
    createdAt?: Date
    editedAt?: Date
    exerciseID?: string
    status?: string
    makeupCenterTitle?: string
    team: Team
    roleplayer: Roleplayer
    medicalCase: MedicalCase
  }
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

interface Roleplayer {
  fullName?: string
  gender?: string
  age?: number
  nationality?: string
}
