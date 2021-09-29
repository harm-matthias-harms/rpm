import { Author } from '../user/type'

export interface State {
  preset: {
    id?: string
    author: Author
    editor: Author
    createdAt?: Date
    editedAt?: Date
    title: string
    vitalSigns: {
      oos?: string
      avpu?: string
      mobility?: string
      respiratoryRate?: number
      pulse?: number
      temperature?: number
      capillaryRefill?: number
      bloodPressureSystolic?: number
      bloodPressureDiastolic?: number
      oxygenSaturation?: number
      expectations: {
        foe?: string
        treatmentExpected?: string
      }
    }
  }
  presetList: {
    count: number
    presets: {
      id: string
      author: Author
      title: string
    }[]
  }
  presetsLoaded: boolean
}
