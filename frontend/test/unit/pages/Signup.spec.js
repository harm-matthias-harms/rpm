import Vue from 'vue'
import { mount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import Vuex from 'vuex'
import VeeValidate from 'vee-validate'
import VueRouter from 'vue-router'
import { store } from '../utils/vuex-store'
import SignUp from '@/pages/sign_up.vue'

Vue.use(Vuex)
Vue.use(VeeValidate)
Vue.use(VueRouter)
Vue.use(Vuetify)

describe('SignUp', () => {
  let wrapper
  let router
  beforeEach(() => {
    router = new VueRouter()
    wrapper = mount(SignUp, {
      stubs: {
        NuxtLink: RouterLinkStub,
        RouterLink: RouterLinkStub,
      },
      provide: {
        $validator () {
          return new VeeValidate.Validator()
        },
      },
      sync: false,
      store,
      router,
    })
  })
  test('is a Vue instance', () => {
    expect(wrapper.isVueInstance()).toBeTruthy()
  })
  test('loads user on error', () => {
    const storeCopy = store
    storeCopy.state.user.user = {
      username: 'name',
      email: 'mail@mail.com',
      password: '123',
    }
    storeCopy.state.user.registerError = true
    wrapper = mount(SignUp, {
      stubs: {
        NuxtLink: RouterLinkStub,
        RouterLink: RouterLinkStub,
      },
      store: storeCopy,
      router,
    })
    expect(wrapper.vm.user.username).toEqual('name')
    expect(wrapper.vm.user.email).toEqual('mail@mail.com')
    expect(wrapper.vm.user.password).toEqual('123')
  })
  test('redirects if account allready exists', () => {
    const storeCopy = store
    storeCopy.state.user.registerError = true
    storeCopy.state.user.registerErrorReason = 'already exists'
    wrapper = mount(SignUp, {
      stubs: {
        NuxtLink: RouterLinkStub,
        RouterLink: RouterLinkStub,
      },
      store: storeCopy,
      router,
    })
    expect(wrapper.vm.$route.path).toEqual('/')
  })
  test('custom rules', () => {
    wrapper.find('input[name="username"]').setValue('asdasd&')
    wrapper.find('input[name="password"]').setValue('asd asd')
    // TODO the custom rules can't be tested at the moment due to vue-test sync/async bugs with jest
  })
})
