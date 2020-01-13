<template>
  <v-row justify="center">
    <v-col
      lg="6"
      md="10"
      sm="12"
    >
      <v-card>
        <v-card-title primary-title>
          Edit Preset
        </v-card-title>
        <v-card-text>
          <PresetForm
            :preset="preset"
            :at-submit="update"
            :is-new="false"
          />
        </v-card-text>
      </v-card>
    </v-col>
  </v-row>
</template>

<script lang="ts">
  import { Component, Vue } from 'vue-property-decorator'
  import { mapActions } from 'vuex'
  import PresetForm from '@/components/preset/form.vue'
@Component({
  components: {
    PresetForm,
  },
  methods: {
    ...mapActions('preset', {
      find: 'find',
      update: 'update',
    }),
  },
})
  export default class Edit extends Vue {
  find!: (id) => Promise<any>
  update!: (preset) => void
  preset: any = JSON.parse(JSON.stringify(this.$store.state.preset.preset))

  mounted () {
    const id = this.$route.params.id
    if (this.preset.id !== id) {
      this.find({ id }).then(() => {
        this.preset = JSON.parse(JSON.stringify(this.$store.state.preset.preset))
      })
    }
  }
  }
</script>
