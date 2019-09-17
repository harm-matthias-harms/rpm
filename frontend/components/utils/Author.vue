<template>
  <div>
    <h5 class="subtitle-1 black--text">
      <v-icon color="primary">
        fas fa-user
      </v-icon>
      {{ author.username }}
    </h5>
    <h6 class="body-2">
      created at {{ $moment(createdAt).format("MMM DD'YY") }}
    </h6>
  </div>
</template>

<script lang="ts">
import { Prop, Component, Vue } from 'vue-property-decorator'
@Component
export default class Author extends Vue {
  @Prop({ type: Object, required: true }) readonly author!: object
  @Prop({ type: String }) readonly createdAt?: string
}
</script>
