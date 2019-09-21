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
        preset: { vitalSigns: {} },
        atSubmit: () => {}
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
    expect(wrapper.isVueInstance()).toBeTruthy()
  })
})
