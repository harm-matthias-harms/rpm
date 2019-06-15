/// <reference types="jest" />
import Vue from 'vue'
import { mount, RouterLinkStub } from '@vue/test-utils'
import CookieHint from '@/components/utils/CookieHint.vue'
import Vuetify from 'vuetify'
import LocalStorageMock from '@/test/unit/utils/local-storage-mock'

describe('Index', () => {
  let wrapper
  beforeEach(() => {
    Vue.use(Vuetify)
    wrapper = mount(CookieHint, {
      stubs: {
        NuxtLink: RouterLinkStub
      }
    })
    global.localStorage = LocalStorageMock
  })
  test('is a Vue instance', () => {
    expect(wrapper.isVueInstance()).toBeTruthy()
  })
  test('snackbar should be false by default', () => {
    expect(wrapper.snackbar).toBeFalsy()
  })
  test('sets localstorage on close', () => {
    wrapper.vm.acceptCookieHint()
    expect(wrapper.snackbar).toBeFalsy()
    expect(localStorage.getItem('accepted_cookie_hint')).toEqual('true')
  })
})
