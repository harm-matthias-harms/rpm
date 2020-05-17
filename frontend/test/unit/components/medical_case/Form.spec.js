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
        RouterLink: RouterLinkStub
      },
      propsData: {
        medicalCase: {
          id: '002',
          general: { scenario: ['scenario'] },
          patient: { type: 'type' },
          medical: { signs: 'signs' },
          makeup: { makeup: 'make-up' },
          vitalSigns: []
        },
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
    wrapper.vm.expansionPanel = []
    expect(wrapper).toBeTruthy()
    wrapper.vm.addVitalSign()
    expect(wrapper.vm.medicalCase.vitalSigns.length).toBe(1)
    wrapper.vm.setExpansionPanel()
    wrapper.vm.medicalCase.vitalSigns[0].data.temperature = ''
    wrapper.vm.medicalCase.vitalSigns[0].data.expectations.foe = ''
    wrapper.vm.medicalCase.vitalSigns[0].childs = [{ data: { expectations: {} }, childs: [] }]
    wrapper.vm.submit()
    expect(wrapper.vm.medicalCase.vitalSigns[0].data.temperature).not.toBeDefined()
  })
})
