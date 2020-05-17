<template>
  <div>
    <v-card>
      <v-card-title primary-title>
        <h3 class="display-1">
          Use a code
        </h3>
      </v-card-title>
      <v-card-text>
        <v-form @submit.prevent="signin({code: code})">
          <v-text-field
            v-model="code"
            label="Code"
            name="code"
          />
          <v-btn
            :disabled="!code"
            color="success"
            type="submit"
          >
            Use code
          </v-btn>
        </v-form>
      </v-card-text>
    </v-card>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import { mapActions } from 'vuex'
@Component({
  methods: {
    ...mapActions('user', {
      signin: 'signin'
    })
  }
})
export default class Code extends Vue {
  code: string = ''
}
</script>
