import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import NewMedicalCase from '@/pages/medical_cases/new.vue'
import { store } from '../../utils/vuex-store'

Vue.use(Vuetify)

describe('New Medical Case', () => {
  let wrapper
  beforeEach(() => {
    wrapper = shallowMount(NewMedicalCase, {
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
