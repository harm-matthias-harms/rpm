<template>
  <div>
    <PrintRoleplayer
      :inject="inject"
      v-for="(inject, idx) in injects"
      :key="idx"
    />
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import { mapActions } from 'vuex'

import { Inject } from '~/store/inject/type'
import PrintRoleplayer from '@/components/inject/print/Rolepayer.vue'

import _ from 'lodash'

@Component({
  layout: 'print',
  components: {
    PrintRoleplayer,
  },
  methods: {
    ...mapActions('inject', {
      findInject: 'find',
    }),
  },
})
export default class PrintVitalSigns extends Vue {
  findInject!: ({ id, exerciseID, disableLoader }) => Promise<Inject>

  injects: Inject[] = []

  async created() {
    const ids: string | (string | null)[] =
      typeof this.$route.query.injects === 'string'
        ? [this.$route.query.injects]
        : this.$route.query.injects

    for (const id of ids) {
      const inject: Inject = await this.$axios.$get(
        `/api/exercises/${this.$route.params.id}/injects/${id}`
      )
      this.injects.push(inject)
    }
    setTimeout(window.print, 500)
  }

  get _() {
    return _
  }
}
</script>

<style scoped>
html {
  -webkit-print-color-adjust: exact;
}

@page {
  size: 8.27in 11.69in;
  margin: 0;
}

html,
body {
  display: block;
  min-height: 0;
  background-color: #ffffff;
}
</style>