import { State } from './type'

export const state = (): State => ({
  medicalCase: {
    id: undefined,
    author: {
      id: undefined,
      username: undefined
    },
    editor: {
      id: undefined,
      username: undefined
    },
    createdAt: undefined,
    updatedAt: undefined,
    title: '',
    approved: false,
    general: {
      discipline: [],
      context: [],
      scenario: []
    },
    patient: {
      type: undefined,
      triage: undefined,
      gender: [],
      age: undefined
    },
    medical: {
      signs: undefined,
      allergies: undefined,
      medication: undefined,
      past: undefined,
      loi: undefined,
      events: undefined
    },
    makeup: {
      makeup: undefined,
      acting: undefined
    },
    vitalSigns: [],
    files: []
  },
  medicalCasesList: {
    count: 0,
    medicalCases: []
  },
  medicalCasesLoaded: false
})

export default state
