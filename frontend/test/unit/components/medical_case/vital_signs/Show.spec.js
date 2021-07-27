import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import VitalSigns from '@/components/medical_case/vital_signs/show.vue'

Vue.use(Vuetify)

describe('Medical case show vital signs', () => {
  let wrapper
  beforeEach(() => {
    wrapper = shallowMount(VitalSigns, {
      stubs: {
        NuxtLink: RouterLinkStub,
        RouterLink: RouterLinkStub
      },
      propsData: {
        vitalSigns: [
          {
            data: { temperature: 1, expectations: {} },
            childs: []
          }
        ]
      }
    })
  })
  test('is a Vue instance', () => {
    expect(wrapper).toBeTruthy()
  })
})
