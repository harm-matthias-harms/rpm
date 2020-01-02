import Vue from 'vue'
import { shallowMount, RouterLinkStub } from '@vue/test-utils'
import Vuetify from 'vuetify'
import Author from '@/components/utils/Author.vue'

const $moment = () => ({ format: () => '1970-01-01T00:00:00' })

Vue.use(Vuetify)

describe('Editor', () => {
  let wrapper

  beforeEach(() => {
    wrapper = shallowMount(Author, {
      stubs: {
        NuxtLink: RouterLinkStub,
        RouterLink: RouterLinkStub,
      },
      propsData: {
        author: {
          id: '001',
          username: 'username',
        },
        createdAt: '1970-01-01T00:00:00',
      },
      mocks: {
        $moment,
      },
    })
  })
  test('is a Vue instance', () => {
    expect(wrapper.isVueInstance()).toBeTruthy()
    expect(wrapper.props().createdAt).toEqual('1970-01-01T00:00:00')
    expect(wrapper.props().author).toMatchObject({
      id: '001',
      username: 'username',
    })
  })
})
