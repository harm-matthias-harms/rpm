import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import Vuex from 'vuex'
import VeeValidate from 'vee-validate'
import VueRouter from 'vue-router'
import { store } from '../utils/vuex-store'
import SignUp from '@/pages/signup.vue'

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
    wrapper = shallowMount(SignUp, {
      stubs: {
        NuxtLink: RouterLinkStub,
        RouterLink: RouterLinkStub
      },
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
    wrapper = shallowMount(SignUp, {
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
    wrapper = shallowMount(SignUp, {
      stubs: {
        NuxtLink: RouterLinkStub,
        RouterLink: RouterLinkStub
      },
      store: storeCopy,
      router
    })
    expect(wrapper.vm.$route.path).toEqual('/signin')
  })
  test('redirects on succes', () => {
    const storeCopy = store
    storeCopy.state.user.registerSuccess = true
    wrapper = shallowMount(SignUp, {
      stubs: {
        NuxtLink: RouterLinkStub,
        RouterLink: RouterLinkStub
      },
      store: storeCopy,
      router
    })
    expect(wrapper.vm.$route.path).toEqual('/accountcreated')
  })
})
