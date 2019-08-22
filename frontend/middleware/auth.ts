import Cookie from 'js-cookie'
import jwtDecode from 'jwt-decode'
export default function(context) {
  // const
  const blockUnAuthorized = ['']
  // checl for cookie
  const cookie = Cookie.get('Authorization')
  if (cookie) {
    const decoded = jwtDecode(cookie)
    context.store.commit('user/SET_AUTHENTICATE', {
      id: decoded.id,
      username: decoded.username,
      expire: decoded.exp
    })
    if (new Date(decoded.exp * 1000) < new Date()) {
      Cookie.remove('Authorization')
      context.store.commit('user/LOGOUT')
      context.store.commit('snackbar/SET', "You're session has expired")
      return context.redirect('/')
    }
  }
  // Check if has AccessRights
  if (!context.store.state.user.isAuthenticated) {
    if (blockUnAuthorized.includes(context.route.name)) {
      context.store.commit(
        'snackbar/SET',
        'You are unauthorized to enter this area, please log in first'
      )
      return context.redirect('/')
    }
  }
}
