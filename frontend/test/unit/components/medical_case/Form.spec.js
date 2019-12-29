import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import VeeValidate from 'vee-validate'
import Form from '@/components/medical_case/form.vue'

Vue.use(Vuetify)
Vue.use(VeeValidate)

describe('Medical case form', () => {
  let wrapper
  beforeEach(() => {
    wrapper = shallowMount(Form, {
      stubs: {
        NuxtLink: RouterLinkStub,
        RouterLink: RouterLinkStub,
      },
      propsData: {
        medicalCase: {
          generalInformation: {},
          medicalHistory: {},
          expectations: {},
          vitalSigns: [],
          files: [],
        },
        atSubmit: () => {},
      },
      provide: {
        $validator () {
          return new VeeValidate.Validator()
        },
      },
      sync: false,
    })
  })
  test('is a Vue instance', () => {
    expect(wrapper.isVueInstance()).toBeTruthy()
    wrapper.vm.addVitalSign()
    expect(wrapper.vm.medicalCase.vitalSigns.length).toBe(1)
  })
})
