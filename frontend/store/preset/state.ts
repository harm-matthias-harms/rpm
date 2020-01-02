import { State } from './type'

export const state = (): State => ({
  preset: {
    id: undefined,
    author: {
      id: undefined,
      username: undefined,
    },
    editor: {
      id: undefined,
      username: undefined,
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
      height: undefined,
    },
  },
  presetList: {
    count: 0,
    presets: [],
  },
  presetsLoaded: false,
})

export default state
