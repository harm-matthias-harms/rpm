import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import Confirm from '@/components/utils/Confirm.vue'

Vue.use(Vuetify)

describe('Confirm', () => {
  let wrapper

  beforeEach(() => {
    wrapper = shallowMount(Confirm, {
      stubs: {
        NuxtLink: RouterLinkStub,
        RouterLink: RouterLinkStub
      },
      propsData: {
        dialog: true,
        text: 'text',
        item: {},
        atSubmit: () => {},
        atCancel: () => {}
      }
    })
  })
  test('is a Vue instance', () => {
    expect(wrapper.isVueInstance()).toBeTruthy()
  })
})
