import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import MedicalCases from '@/pages/medical_cases/index.vue'
import { store } from '../../utils/vuex-store'

Vue.use(Vuetify)

describe('Index Medical Cases', () => {
  let wrapper
  beforeEach(() => {
    wrapper = shallowMount(MedicalCases, {
      stubs: {
        NuxtLink: RouterLinkStub,
        RouterLink: RouterLinkStub
      },
      store
    })
  })
  test('is a Vue instance', () => {
    expect(wrapper.isVueInstance()).toBeTruthy()
  })
})
