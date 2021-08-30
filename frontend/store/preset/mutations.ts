import { MutationTree } from 'vuex'
import { defaultPreset } from './state'
import { State } from './type'

export const mutations: MutationTree<State> = {
  SET_PRESET (state, preset) {
    state.preset = preset
  },
  SET_PRESET_LIST (state, list) {
    state.presetList = list
    state.presetsLoaded = true
  },
  SET_PRESET_TO_LIST (state, preset) {
    state.presetList.count += 1
    state.presetList.presets.unshift({
      id: preset.id,
      author: preset.author,
      title: preset.title
    })
  },
  DELETE_FROM_LIST (state, id) {
    state.presetList.presets = state.presetList.presets.filter(
      item => item.id !== id
    )
    state.presetList.count--
  },
  UNSET_PRESET (state) {
    state.preset = defaultPreset
  }
}

export default mutations
