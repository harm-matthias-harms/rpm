<template>
  <div>
    <v-autocomplete
      v-model="startTime"
      :items="options"
      item-text="text"
      item-value="value"
    />
  </div>
</template>

<script lang="ts">
import { Prop, Component, Vue } from 'vue-property-decorator'
import { mapActions } from 'vuex'

interface DateOption {
  value: Date
  text: string
}

@Component({
  methods: {
    ...mapActions('exercise', {
      findExercise: 'find'
    })
  }
})
export default class StartTime extends Vue {
  @Prop({ type: Date }) readonly value!: Date | undefined

  findExercise!: (id) => Promise<void>

  get startTime (): Date | undefined {
    return this.value
  }

  set startTime (time: Date | undefined) {
    this.$emit('input', time)
  }

  options: DateOption[] = []

  mounted () {
    const id = this.$route.params.id
    if (this.$store.state.exercise.exercise.id === id) {
      this.options = this.generateTimes()
    } else {
      this.findExercise({ id }).then(() => {
        this.options = this.generateTimes()
      })
    }
  }

  generateTimes (): DateOption[] {
    const startTime = new Date(
      Math.max(
        ...[
          new Date().getTime(),
          new Date(this.$store.state.exercise.exercise.startTime).getTime()
        ]
      )
    ).toISOString()
    return this.makeTimeArray(
      startTime,
      this.$store.state.exercise.exercise.endTime
    )
  }

  makeTimeArray (startTime, endTime): DateOption[] {
    const times: DateOption[] = []
    const hours = [
      '00:00 - 03:00',
      '03:00 - 06:00',
      '06:00 - 09:00',
      '09:00 - 12:00',
      '12:00 - 15:00',
      '15:00 - 18:00',
      '18:00 - 21:00',
      '21:00 - 00:00'
    ]
    const hoursLength = hours.length
    const j = this.getCurrentTimeIndex(startTime, hours)
    for (
      let d = new Date(startTime);
      d <= new Date(endTime);
      d.setDate(d.getDate() + 1)
    ) {
      if (d === new Date(startTime)) {
        for (let i = j; i < hoursLength; i++) {
          const date = new Date(d)
          times.push({
            value: date,
            text: date.toISOString().slice(0, 10) + ' ' + hours[i]
          })
        }
      } else {
        for (let i = 0; i < hoursLength; i++) {
          const date = new Date(d)
          times.push({
            value: date,
            text: date.toISOString().slice(0, 10) + ' ' + hours[i]
          })
        }
      }
    }
    return times
  }

  getCurrentTimeIndex (startTime, hours): number {
    const startTimeHour = parseInt(startTime.slice(11, 13))
    for (let i = 0; i < hours.length; i++) {
      const arr = hours[i].split(' - ')
      const min = parseInt(arr[0].slice(0, 2))
      const max = parseInt(arr[1].slice(0, 2))
      if (min <= startTimeHour && startTimeHour <= max) {
        return i
      }
    }
    return 0
  }
}
</script>
