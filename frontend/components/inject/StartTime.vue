<template>
  <div>
    <v-autocomplete
      v-model="startTime"
      :items="options"
      item-text="text"
      item-value="value"
    >
      <template v-slot:item="{ item }">
        <div :class="item.isFuture ? null : 'text-decoration-line-through'">
          {{ item.text }}
        </div>
      </template>
    </v-autocomplete>
  </div>
</template>

<script lang="ts">
import { Prop, Component, Vue } from 'vue-property-decorator'
import { mapActions } from 'vuex'

interface DateOption {
  value: Date
  text: string
  isFuture: boolean
}

@Component({
  methods: {
    ...mapActions('exercise', {
      findExercise: 'find',
    }),
  },
})
export default class StartTime extends Vue {
  @Prop({ type: Date }) readonly value!: Date | undefined

  findExercise!: (id) => Promise<void>

  get startTime(): Date | undefined {
    return this.value
  }

  set startTime(time: Date | undefined) {
    this.$emit('input', time)
  }

  options: DateOption[] = []

  mounted() {
    const id = this.$route.params.id
    if (this.$store.state.exercise.exercise.id === id) {
      this.options = this.generateTimes()
    } else {
      this.findExercise({ id }).then(() => {
        this.options = this.generateTimes()
      })
    }
  }

  generateTimes(): DateOption[] {
    return this.makeTimeArray(
      this.$store.state.exercise.exercise.startTime,
      this.$store.state.exercise.exercise.endTime
    )
  }

  makeTimeArray(startTime, endTime): DateOption[] {
    const times: DateOption[] = []

    const hours = [
      '00:00 - 03:00',
      '03:00 - 06:00',
      '06:00 - 09:00',
      '09:00 - 12:00',
      '12:00 - 15:00',
      '15:00 - 18:00',
      '18:00 - 21:00',
      '21:00 - 00:00',
    ]
    const hoursLength = hours.length
    const endDate = new Date(endTime)
    endDate.setDate(endDate.getDate() + 1)

    for (
      let d = new Date(startTime);
      d <= endDate;
      d.setDate(d.getDate() + 1)
    ) {
      for (let i = 0; i < hoursLength; i++) {
        const date = new Date(d)
        const value = new Date(
          date.toISOString().slice(0, 10) + ' ' + hours[i].slice(0, 4)
        )

        times.push({
          value: value,
          text: date.toISOString().slice(0, 10) + ' ' + hours[i],
          isFuture: value.getTime() >= Date.now(),
        })
      }
    }

    if (
      this.value &&
      !times.some((a: DateOption) => {
        return a.value == this.value!
      })
    ) {
      times.push({
        value: this.value!,
        text: this.value.toLocaleString('sv-SE', { timeZone: 'UTC' }),
        isFuture: this.value.getTime() >= Date.now(),
      })
    }

    times.sort(this.sortByIsFutureAndTime)

    return times
  }

  sortByIsFutureAndTime(a: DateOption, b: DateOption) {
    if (a.isFuture == b.isFuture) {
      return a.value.getTime() - b.value.getTime()
    }
    return a.isFuture > b.isFuture ? -1 : 1
  }
}
</script>
