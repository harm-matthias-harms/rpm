<template>
  <v-row class="mx-3">
    <v-col class="px-0">
      <v-icon color="secondary">
        fas fa-user
      </v-icon>
    </v-col>
    <v-col class="px-2">
      <h5 class="subtitle-1 black--text">
        {{ editor.username }}
        <br>
        <span class="body-2">created at {{ $moment(createdAt).format("MMM DD'YY") }}</span>
      </h5>
    </v-col>
  </v-row>
</template>

<script lang="ts">
  import { Prop, Component, Vue } from 'vue-property-decorator'
@Component
  export default class Editor extends Vue {
  @Prop({ type: Object, required: true }) readonly editor!: object
  @Prop({ type: String }) readonly updatedAt?: string
  }
</script>

<style scoped>
.subtitle-1 {
  line-height: 1rem !important;
}
.body-2 {
  color: #6a737c;
}
</style>
