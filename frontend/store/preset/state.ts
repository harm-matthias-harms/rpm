import { author } from '../user/state'
import { State } from './type'

export const state = (): State => ({
  preset: {
    id: undefined,
    author: author,
    editor: author,
    createdAt: undefined,
    editedAt: undefined,
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
      expectations: {
        foe: undefined,
        treatmentExpected: undefined,
      },
    },
  },
  presetList: {
    count: 0,
    presets: [],
  },
  presetsLoaded: false,
})

export default state
