import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import Exams from '@/components/medical_case/Exams.vue'

Vue.use(Vuetify)

describe('Exams', () => {
  let wrapper
  beforeEach(() => {
    wrapper = shallowMount(Exams, {
      stubs: {
        NuxtLink: RouterLinkStub,
        RouterLink: RouterLinkStub
      }
    })
  })
  test('is a Vue instance', () => {
    expect(wrapper).toBeTruthy()
  })
})
