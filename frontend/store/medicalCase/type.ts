interface nestedVitalSign {
  title?: string
  reason?: string
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
    weight?: number
    height?: number
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
    makeup?: string
    otherInformation?: string
    generalInformation: {
      surgical: boolean
      hospilisation: boolean
      usar: boolean
      medivac: boolean
      triage?: string
      shortSummary?: string
      age?: string
      gender?: string
    }
    medicalHistory: {
      problems?: string
      vaccinations?: string
      allergies?: string
      medication?: string
      implantedDevices?: string
    }
    expectations: {
      generalStatus?: string
      onExamination?: string
      expectations?: string
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
      generalInformation: {
        surgical: boolean
        hospilisation: boolean
        usar: boolean
        medivac: boolean
        triage?: string
        age?: string
        gender?: string
      }
    }[]
  }
  medicalCasesLoaded: boolean
}
