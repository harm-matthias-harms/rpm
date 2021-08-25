import { author } from '../user/state'
import { MedicalCase, State } from './type'

export const medicalCase: MedicalCase = {
  id: undefined,
  author: author,
  editor: author,
  createdAt: undefined,
  editedAt: undefined,
  title: '',
  approved: false,
  general: {
    discipline: [],
    context: [],
    scenario: [],
  },
  patient: {
    type: undefined,
    triage: undefined,
    gender: [],
    age: undefined,
  },
  medical: {
    signs: undefined,
    allergies: undefined,
    medication: undefined,
    past: undefined,
    loi: undefined,
    events: undefined,
  },
  makeup: {
    makeup: undefined,
    acting: undefined,
  },
  vitalSigns: [],
  files: [],
}

export const state = (): State => ({
  medicalCase: medicalCase,
  medicalCasesList: {
    count: 0,
    medicalCases: [],
  },
  medicalCasesLoaded: false,
})

export default state
