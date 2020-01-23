import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import Editor from '@/components/utils/Editor.vue'

const $moment = () => ({ format: () => '1970-01-01T00:00:00' })

Vue.use(Vuetify)

describe('Editor', () => {
  let wrapper

  beforeEach(() => {
    wrapper = shallowMount(Editor, {
      stubs: {
        NuxtLink: RouterLinkStub,
        RouterLink: RouterLinkStub
      },
      propsData: {
        editor: {
          id: '001',
          username: 'username'
        },
        updatedAt: '1970-01-01T00:00:00'
      },
      mocks: {
        $moment
      }
    })
    // wrapper.vm.$moment = moment
    // inject('moment', moment)
  })
  test('is a Vue instance', () => {
    expect(wrapper.isVueInstance()).toBeTruthy()
    expect(wrapper.props().updatedAt).toEqual('1970-01-01T00:00:00')
    expect(wrapper.props().editor).toMatchObject({
      id: '001',
      username: 'username'
    })
  })
})
