import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import RegisterSuccess from '@/components/auth/RegisterSuccess.vue'

Vue.use(Vuetify)

describe('Index', () => {
  let wrapper
  beforeEach(() => {
    wrapper = shallowMount(RegisterSuccess, {
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
