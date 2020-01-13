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
          id: '002',
          generalInformation: { usar: true },
          medicalHistory: { problems: 'problem' },
          expectations: { expectations: 'expectations' },
          vitalSigns: [],
        },
        atSubmit: () => {},
        isNew: false,
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
    wrapper.vm.expansionPanel = []
    expect(wrapper.isVueInstance()).toBeTruthy()
    wrapper.vm.addVitalSign()
    expect(wrapper.vm.medicalCase.vitalSigns.length).toBe(1)
    wrapper.vm.setExpansionPanel()
    wrapper.vm.medicalCase.vitalSigns[0].data.height = ''
    wrapper.vm.medicalCase.vitalSigns[0].childs = [{ data: {}, childs: [] }]
    wrapper.vm.submit()
    expect(wrapper.vm.medicalCase.vitalSigns[0].data.height).not.toBeDefined()
  })
})
