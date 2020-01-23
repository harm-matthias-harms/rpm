import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import VitalSigns from '@/components/preset/vitalSigns.vue'

Vue.use(Vuetify)

const vitalSign = { temperature: 37 }

describe('VitalSigns', () => {
  let wrapper
  beforeEach(() => {
    wrapper = shallowMount(VitalSigns, {
      stubs: {
        NuxtLink: RouterLinkStub,
        RouterLink: RouterLinkStub
      },
      propsData: {
        vitalSigns: vitalSign
      }
    })
  })
  test('is a Vue instance', () => {
    expect(wrapper.isVueInstance()).toBeTruthy()
    expect(wrapper.props().vitalSigns).toMatchObject(vitalSign)
  })
})
