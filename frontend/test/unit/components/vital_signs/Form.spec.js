import Vue from 'vue'
import { mount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import VeeValidate from 'vee-validate'
import Form from '@/components/vital_signs/form.vue'

Vue.use(Vuetify)
Vue.use(VeeValidate)

describe('Vitalsigns form', () => {
  let wrapper
  beforeEach(() => {
    wrapper = mount(Form, {
      stubs: {
        NuxtLink: RouterLinkStub,
        RouterLink: RouterLinkStub,
      },
      propsData: {
        vitalSigns: {},
      },
      provide: {
        $validator () {
          return new VeeValidate.Validator()
        },
      },
      vuetify: new Vuetify(),
      sync: false,
    })
  })
  test('is a Vue instance', (done) => {
    expect(wrapper.isVueInstance()).toBeTruthy()
    wrapper.find('input[name="oos"]').setValue('test')

    wrapper.vm
      .$nextTick()
      .then(() => {
        expect(wrapper.emitted('update:vitalSigns')[0][0]).toBe(wrapper.vm.vitalSigns)
        done()
      })
      .catch(done)
  })
})
