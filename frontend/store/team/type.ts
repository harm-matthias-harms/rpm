export interface State {
    preset: {
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
            weight?: number
            height?: number
        }
    }
    presetList: {
        count: number
        presets: {
            id: string
            author: {
                id: string
                username: string
            }
            title: string
        }[]
    }
    presetsLoaded: boolean
}
