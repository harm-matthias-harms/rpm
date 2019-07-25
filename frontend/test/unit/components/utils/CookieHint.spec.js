import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import CookieHint from '@/components/utils/CookieHint.vue'
import LocalStorageMock from '@/test/unit/utils/local-storage-mock'

Vue.use(Vuetify)

describe('Index', () => {
  let wrapper
  beforeEach(() => {
    wrapper = shallowMount(CookieHint, {
      stubs: {
        NuxtLink: RouterLinkStub,
        RouterLink: RouterLinkStub
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
