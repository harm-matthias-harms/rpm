<template>
  <div>
    <h5 class="subtitle-1 black--text">
      <v-icon color="secondary">
        fas fa-user
      </v-icon>
      {{ editor.username }}
    </h5>
    <h6 class="body-2">
      edited at {{ $moment(updatedAt).format("MMM DD'YY") }}
    </h6>
  </div>
</template>

<script lang="ts">
  import { Prop, Component, Vue } from 'vue-property-decorator'
@Component
  export default class Editor extends Vue {
  @Prop({ type: Object, required: true }) readonly editor!: object
  @Prop({ type: String }) readonly updatedAt?: string
  }
</script>
