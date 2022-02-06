import { Author } from '../user/type'

export interface nestedVitalSign {
  title?: string
  data?: {
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
  childs?: nestedVitalSign[]
}

interface file {
  id?: string
  name?: string
  size?: number
}

interface fileArray {
  [index: number]: file
}

export interface MedicalCase {
  id?: string
  author: Author
  editor: Author
  createdAt?: Date
  editedAt?: Date
  title: string
  approved: boolean
  general: {
    discipline: string[]
    context: string[]
    scenario: string[]
    preHospital?: boolean
    medevac?: boolean
  }
  patient: {
    type?: string
    triage?: string
    gender: string[]
    age?: string
  }
  medical: {
    signs?: string
    allergies?: string
    medication?: string
    past?: string
    loi?: string
    events?: string
  }
  makeup: {
    makeup?: string
    acting?: string
  }
  vitalSigns: nestedVitalSign[]
  files: fileArray
}

export interface MedicalCaseShort {
  id: string
  author: Author
  title: string
  approved: boolean
  general: {
    discipline?: string
    context: string[]
    scenario: string[]
    preHospital?: boolean
    medevac?: boolean
  }
  patient: {
    triage?: string
    gender: string[]
    age?: string
  }
}

export interface State {
  medicalCase: MedicalCase
  medicalCasesList: {
    count: number
    medicalCases: MedicalCaseShort[]
  }
  medicalCasesLoaded: boolean
}
