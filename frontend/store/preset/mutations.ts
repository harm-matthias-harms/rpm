import { MutationTree } from 'vuex'
import { State } from './type'

export const mutations: MutationTree<State> = {
  SET_PRESET (state, preset) {
    state.preset = preset
  },
  SET_PRESET_LIST (state, list) {
    state.presetList = list
    state.presetsLoaded = true
  },
  UNSET_PRESET (state) {
    state.preset = {
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
      vitalSigns: {
        oos: undefined,
        avpu: undefined,
        mobility: undefined,
        respiratoryRate: undefined,
        pulse: undefined,
        temperature: undefined,
        capillaryRefill: undefined,
        bloodPressureSystolic: undefined,
        bloodPressureDiastolic: undefined,
        oxygenSaturation: undefined,
        weight: undefined,
        height: undefined
      }
    }
  }
}

export default mutations
