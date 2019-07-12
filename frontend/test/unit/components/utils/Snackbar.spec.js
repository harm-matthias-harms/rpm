import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import Vuex from 'vuex'
import Snackbar from '@/components/utils/Snackbar.vue'
import { store } from '@/test/unit/utils/vuex-store'

Vue.use(Vuetify)
Vue.use(Vuex)

describe('Index', () => {
  let wrapper
  beforeEach(() => {
    wrapper = shallowMount(Snackbar, {
      stubs: {
        NuxtLink: RouterLinkStub,
        RouterLink: RouterLinkStub
      },
      store
    })
  })
  test('is a Vue instance', () => {
    expect(wrapper.isVueInstance()).toBeTruthy()
  })
  test('snackbar has empty default state', () => {
    expect(wrapper.vm.show).toBeFalsy()
    expect(wrapper.vm.message).toEqual('')
  })
  test('on Commit set Snackbar', done => {
    wrapper.vm.$store.commit('snackbar/SET', 'test')
    wrapper.vm
      .$nextTick()
      .then(() => {
        expect(wrapper.vm.message).toEqual('test')
        expect(wrapper.vm.show).toBeTruthy()
        done()
      })
      .catch(done)
    // expect(wrapper.vm.$store.state.snackbar.message).toEqual('test')
  })
})
