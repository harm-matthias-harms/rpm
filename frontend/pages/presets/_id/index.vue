<template>
  <v-container>
    <v-row justify="center">
      <v-col
        lg="8"
        md="10"
        sm="12"
      >
        <v-card>
          <v-card-text>
            <h4 class="display-1 font-weight-light mb-2 black--text">
              {{ preset.title }}
              <v-icon
                color="primary"
                @click="editPreset(preset)"
              >
                edit
              </v-icon>
              <DeleteButton
                v-if="preset.author.id == $store.state.user.user.id"
                :item="preset"
                :go-back="true"
              />
            </h4>

            <v-row>
              <v-col
                v-if="preset.author && preset.author.username"
                class="col-auto"
              >
                <Author
                  :author="preset.author"
                  :created-at="preset.createdAt"
                />
              </v-col>
              <v-col
                v-if="preset.editor && preset.editor.username"
                class="col-auto mr-auto"
              >
                <Editor
                  :editor="preset.editor"
                  :updated-at="preset.editedAt"
                />
              </v-col>
            </v-row>
          </v-card-text>
          <VitalSigns
            class="black--text"
            :vital-signs="preset.vitalSigns"
          />
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import { mapState, mapActions } from 'vuex'
import Author from '@/components/utils/Author.vue'
import Editor from '@/components/utils/Editor.vue'
import VitalSigns from '@/components/preset/vitalSigns.vue'
import DeleteButton from '@/components/preset/Delete.vue'
@Component({
  components: {
    Author,
    Editor,
    VitalSigns,
    DeleteButton
  },
  computed: {
    ...mapState('preset', {
      preset: 'preset'
    })
  },
  methods: {
    ...mapActions('preset', {
      find: 'find'
    })
  }
})
export default class ShowPreset extends Vue {
  find!: (id) => void
  preset!: any

  mounted () {
    const id = this.$route.params.id
    if (this.preset.id !== id) {
      this.find({ id })
    }
  }

  editPreset (preset) {
    this.$router.push('/presets/' + preset.id + '/edit')
  }
}
</script>
