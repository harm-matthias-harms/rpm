import Vue from 'vue'
import { mount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import Vuex from 'vuex'
import VeeValidate from 'vee-validate'
import VueRouter from 'vue-router'
import SignIn from '@/components/auth/SignIn.vue'
import { store } from '../../utils/vuex-store'

Vue.use(Vuex)
Vue.use(VeeValidate)
Vue.use(VueRouter)
Vue.use(Vuetify)

describe('SignIn', () => {
  let wrapper
  let router
  beforeEach(() => {
    router = new VueRouter()
    wrapper = mount(SignIn, {
      stubs: {
        NuxtLink: RouterLinkStub,
        RouterLink: RouterLinkStub
      },
      provide: {
        $validator () {
          return new VeeValidate.Validator()
        }
      },
      sync: false,
      store,
      router
    })
  })
  test('is a Vue instance', () => {
    expect(wrapper).toBeTruthy()
  })
  test('custom rules', () => {
    wrapper.find('input[name="username"]').setValue('asdasd&')
    wrapper.find('input[name="password"]').setValue('asd asd')
    // TODO the custom rules can't be tested at the moment due to vue-test sync/async bugs with jest
  })
})
