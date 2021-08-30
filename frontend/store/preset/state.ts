import { author } from '../user/state'
import { State } from './type'

export const defaultPreset = {
  id: undefined,
  author,
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
      treatmentExpected: undefined
    }
  }
}

export const state = (): State => ({
  preset: defaultPreset,
  presetList: {
    count: 0,
    presets: []
  },
  presetsLoaded: false
})

export default state
