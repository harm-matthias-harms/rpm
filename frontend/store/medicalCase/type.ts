interface nestedVitalSign {
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
  childs?: nestedVitalSignArray
}

interface nestedVitalSignArray {
  [index: number]: nestedVitalSign
}

interface file {
  id?: string
  name?: string
  size?: number
}

interface fileArray {
    [index: number]: file
}

export interface State {
  medicalCase: {
    id?: string
    author: {
      id?: string
      username?: string
    }
    editor: {
      id?: string
      username?: string
    }
    createdAt?: Date
    updatedAt?: Date
    title: string
    approved: boolean
    general: {
      discipline: string[]
      context: string[]
      scenario: string[]
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
    vitalSigns: nestedVitalSignArray
    files: fileArray
  }
  medicalCasesList: {
    count: number
    medicalCases: {
      id: string
      author: {
        id: string
        username: string
      }
      title: string
      approved: boolean
      general: {
        discipline?: string
        context: string[]
        scenario: string[]
      }
      patient: {
        triage?: string
      }
    }[]
  }
  medicalCasesLoaded: boolean
}
