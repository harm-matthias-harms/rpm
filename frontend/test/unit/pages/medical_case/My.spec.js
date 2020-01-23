import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import { store } from '../../utils/vuex-store'
import MyMedicalCases from '@/pages/medical_cases/my.vue'

Vue.use(Vuetify)

describe('My Medical Cases', () => {
  let wrapper
  beforeEach(() => {
    wrapper = shallowMount(MyMedicalCases, {
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
