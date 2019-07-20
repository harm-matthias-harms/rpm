import Vue from 'vue'
import { mount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import Vuex from 'vuex'
import VeeValidate from 'vee-validate'
import VueRouter from 'vue-router'
import flushPromises from 'flush-promises'
import { store } from '../utils/vuex-store'
import SignUp from '@/pages/sign_up.vue'

Vue.use(Vuetify)
Vue.use(Vuex)
Vue.use(VeeValidate)
Vue.use(VueRouter)

describe('Index', () => {
  let wrapper
  let router
  beforeEach(() => {
    router = new VueRouter()
    Vue.use(Vuetify)
    wrapper = mount(SignUp, {
      stubs: {
        NuxtLink: RouterLinkStub,
        RouterLink: RouterLinkStub
      },
      provide: {
        $validator() {
          return new VeeValidate.Validator()
        }
      },
      sync: false,
      store,
      router
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
      password: '123'
    }
    storeCopy.state.user.registerError = true
    wrapper = mount(SignUp, {
      stubs: {
        NuxtLink: RouterLinkStub,
        RouterLink: RouterLinkStub
      },
      store: storeCopy,
      router
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
        RouterLink: RouterLinkStub
      },
      store: storeCopy,
      router
    })
    expect(wrapper.vm.$route.path).toEqual('/sign_in')
  })
  test('redirects on succes', () => {
    const storeCopy = store
    storeCopy.state.user.registerSuccess = true
    wrapper = mount(SignUp, {
      stubs: {
        NuxtLink: RouterLinkStub,
        RouterLink: RouterLinkStub
      },
      store: storeCopy,
      router
    })
    expect(wrapper.vm.$route.path).toEqual('/account_created')
  })
  test('custom rules', async () => {
    wrapper.find('input[name="username"]').setValue('asdasd&')
    await flushPromises()
    expect(wrapper.vm.errors.collect('username')).toEqual([
      'The username can contain letters (a-z), numbers (0-9), and periods (.).'
    ])
    wrapper.find('input[name="password"]').setValue('asd asd')
    await flushPromises()
    expect(wrapper.vm.errors.collect('password')).toEqual([
      "The password can't contain a whitespace."
    ])
  })
})
