import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import NewExercise from '@/pages/exercises/new.vue'
import { store } from '../../utils/vuex-store'

Vue.use(Vuetify)

describe('New Exercise', () => {
  let wrapper
  beforeEach(() => {
    wrapper = shallowMount(NewExercise, {
      stubs: {
        NuxtLink: RouterLinkStub,
        RouterLink: RouterLinkStub
      },
      store
    })
  })
  test('is a Vue instance', () => {
    expect(wrapper).toBeTruthy()
  })
})
