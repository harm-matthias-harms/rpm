<template>
  <v-row justify="center">
    <v-col
      lg="6"
      md="10"
      sm="12"
    >
      <v-card>
        <v-card-title primary-title>
          New Preset
        </v-card-title>
        <v-card-text>
          <PresetForm
            :preset="preset"
            :at-submit="create"
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
      create: 'create',
    }),
  },
})
  export default class NewPreset extends Vue {
  preset = { vitalSigns: {} }
  create!: (preset) => void
  }
</script>
