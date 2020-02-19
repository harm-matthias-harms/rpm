import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import VeeValidate from 'vee-validate'
import TeamFrom from '@/components/team/form.vue'

Vue.use(Vuetify)
Vue.use(VeeValidate)

describe('Team form', () => {
  let wrapper
  beforeEach(() => {
    wrapper = shallowMount(TeamFrom, {
      stubs: {
        NuxtLink: RouterLinkStub,
        RouterLink: RouterLinkStub
      },
      propsData: {
        team: {},
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
    wrapper.vm.submit()
    expect(wrapper.isVueInstance()).toBeTruthy()
  })
})
