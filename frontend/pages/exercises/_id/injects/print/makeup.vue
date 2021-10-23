<template>
  <div>
    <div
      v-for="(chunk1, idx) in _.chunk(_.chunk(injects, 2), 2)"
      :key="idx"
      class="page portrait"
    >
      <div class="is-third-page" v-for="(chunk, idx) in chunk1" :key="idx">
        <v-row>
          <v-col class="is-dotted">
            <Makeup :inject="chunk[0]" v-if="chunk.length >= 1"></Makeup>
          </v-col>
          <v-col>
            <Makeup :inject="chunk[1]" v-if="chunk.length >= 2"></Makeup>
          </v-col>
        </v-row>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import { mapActions } from 'vuex'

import { Inject } from '~/store/inject/type'
import Makeup from '@/components/inject/print/Makeup.vue'

import _ from 'lodash'

@Component({
  layout: 'print',
  components: {
    Makeup,
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
    this.$nextTick(() => {
      this.$nextTick(() => {
        window.print()
      })
    })
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

.page {
  position: relative;
  overflow: hidden;
  padding: 0.8in 0.8in;
  page-break-before: always;
}

.page.landscape {
  width: 11.69in;
  height: 8.27in;
  transform: translateY(11.69in) rotate(270deg);
  transform-origin: 0% 0%;
}

.page.portrait {
  width: 8.27in;
  height: 11.69in;
}

.is-third-page {
  height: 5.045in;
}

.is-third-page:not(:last-child) {
  border-bottom: 1px dotted black;
  padding-bottom: 0.5in;
}

.is-third-page:not(:first-child) {
  padding-top: 0.5in;
}

.is-dotted:first-child {
  border-right: 1px dotted black;
}
</style>