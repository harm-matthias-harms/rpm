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
    makeup: undefined,
    otherInformation: undefined,
    generalInformation: {
      surgical: false,
      hospilisation: false,
      usar: false,
      medivac: false,
      triage: undefined,
      shortSummary: undefined,
      age: undefined,
      gender: undefined
    },
    medicalHistory: {
      problems: undefined,
      vaccinations: undefined,
      allergies: undefined,
      medication: undefined,
      implantedDevices: undefined
    },
    expectations: {
      generalStatus: undefined,
      onExamination: undefined,
      expectations: undefined
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
