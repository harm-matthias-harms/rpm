import Cookie from 'js-cookie'
import jwtDecode from 'jwt-decode'
export default function (context) {
  // const
  const blockUnAuthorized = ['presets', 'medical_cases', 'teams']
  // checl for cookie
  const cookie = Cookie.get('Authorization')
  if (cookie) {
    const decoded = jwtDecode(cookie)
    if (new Date(decoded.exp * 1000) < new Date()) {
      Cookie.remove('Authorization')
      context.store.commit('user/LOGOUT')
      context.store.commit('snackbar/SET', "You're session has expired")
      return context.redirect('/')
    } else {
      context.store.commit('user/SET_AUTHENTICATE', {
        id: decoded.id,
        username: decoded.username,
        expire: decoded.exp,
        code: decoded.code
      })
      if (!context.store.state.user.isLoaded) {
        context.store.dispatch('user/load', { id: decoded.id })
      }
    }
  }
  // Check if has AccessRights
  if (!context.store.state.user.isAuthenticated || context.store.state.user.isCodeUser) {
    if (blockUnAuthorized.some(value => context.route.name.includes(value))) {
      context.store.commit(
        'snackbar/SET',
        'You are unauthorized to enter this area, please log in first'
      )
      return context.redirect('/')
    }
  }
}
