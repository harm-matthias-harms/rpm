import Cookie from 'js-cookie'
import jwtDecode from 'jwt-decode'
export default function(context) {
  let blockUnAuthorized = ['']
  let cookie = Cookie.get('Authorization')
  Cookie.set("test", "test")
  console.log(cookie)
  if (cookie) {
    const decoded = jwtDecode(cookie)
    context.store.commit('user/SET_AUTHENTICATE', {
      id: decoded.id,
      username: decoded.username,
      expire: decoded.exp
    })
  }
  if (blockUnAuthorized.includes(context.route.name)) {
    context.store.commit(
      'snackbar/SET',
      'You are unauthorized to enter this area, please log in first.'
    )
    return context.redirect('/')
  }
  console.log(context)
}
