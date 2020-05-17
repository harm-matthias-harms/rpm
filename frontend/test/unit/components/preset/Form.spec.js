import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import VeeValidate from 'vee-validate'
import PresetFrom from '@/components/preset/form.vue'

Vue.use(Vuetify)
Vue.use(VeeValidate)

describe('Preset form', () => {
  let wrapper
  beforeEach(() => {
    wrapper = shallowMount(PresetFrom, {
      stubs: {
        NuxtLink: RouterLinkStub,
        RouterLink: RouterLinkStub
      },
      propsData: {
        preset: { vitalSigns: { expectations: {} } },
        atSubmit: () => {},
        isNew: false
      },
      provide: {
        $validator () {
          return new VeeValidate.Validator()
        }
      },
      sync: false
    })
  })
  test('is a Vue instance', () => {
    expect(wrapper).toBeTruthy()
  })
  test('handle empty vital signs on submit', () => {
    wrapper.vm.preset.vitalSigns.oxygenSaturation = ''
    wrapper.vm.submit()
    expect(wrapper.vm.preset.vitalSigns).toMatchObject({})
  })
})
