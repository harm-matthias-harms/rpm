import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import VeeValidate from 'vee-validate'
import VitalSign from '@/components/medical_case/vitalSign.vue'

Vue.use(Vuetify)
Vue.use(VeeValidate)

describe('Medical case form', () => {
  let wrapper
  beforeEach(() => {
    wrapper = shallowMount(VitalSign, {
      stubs: {
        NuxtLink: RouterLinkStub,
        RouterLink: RouterLinkStub,
      },
      propsData: {
        vitalSign: {
          data: {},
          childs: [],
        },
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
    wrapper.vm.addChild()
    expect(wrapper.vm.vitalSign.childs.length).toBe(1)
  })
})
