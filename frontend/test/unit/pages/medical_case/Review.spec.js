import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import { store } from '../../utils/vuex-store'
import ReviewMedicalCases from '@/pages/medical_cases/review.vue'

Vue.use(Vuetify)

describe('Index Medical Cases', () => {
  let wrapper
  beforeEach(() => {
    wrapper = shallowMount(ReviewMedicalCases, {
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
