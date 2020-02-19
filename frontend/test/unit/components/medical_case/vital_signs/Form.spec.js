import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import VeeValidate from 'vee-validate'
import VitalSign from '@/components/medical_case/vital_signs/form.vue'
import { store } from '../../../utils/vuex-store'

Vue.use(Vuetify)
Vue.use(VeeValidate)

describe('Medical case vital sign form', () => {
  let wrapper
  beforeEach(() => {
    wrapper = shallowMount(VitalSign, {
      stubs: {
        NuxtLink: RouterLinkStub,
        RouterLink: RouterLinkStub
      },
      propsData: {
        vitalSign: {
          data: {},
          childs: []
        }
      },
      provide: {
        $validator () {
          return new VeeValidate.Validator()
        }
      },
      store,
      sync: false
    })
  })
  test('is a Vue instance', () => {
    expect(wrapper.isVueInstance()).toBeTruthy()
    wrapper.vm.addChild()
    expect(wrapper.vm.vitalSign.childs.length).toBe(1)
    wrapper.vm.presetID = '001'
  })
})
