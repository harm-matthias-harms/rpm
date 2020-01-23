import { MutationTree } from 'vuex'
import { State } from './type'

export const mutations: MutationTree<State> = {
  SET_MEDICAL_CASE (state, medicalCase) {
    state.medicalCase = medicalCase
  },
  SET_MEDICAL_CASE_LIST (state, list) {
    state.medicalCasesList = list
    state.medicalCasesLoaded = true
  },
  SET_MEDICAL_CASE_TO_LIST (state, medicalCase) {
    state.medicalCasesList.count += 1
    state.medicalCasesList.medicalCases.unshift({
      id: medicalCase.id,
      author: {
        id: medicalCase.author.id,
        username: medicalCase.author.username
      },
      title: medicalCase.title
    })
  },
  DELETE_FROM_LIST (state, id) {
    state.medicalCasesList.medicalCases = state.medicalCasesList.medicalCases.filter(item => item.id !== id)
  },
  UNSET_MEDICAL_CASE (state) {
    state.medicalCase = {
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
    }
  }
}

export default mutations
