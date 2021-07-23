import Cookies from 'js-cookie'

const cookie_name = 'userInfo'

export function getToken () {
  // return Cookies.get(TokenKey)
  const ui = Cookies.getJSON(cookie_name)
  if (ui) {
    return ui.session.token
  } else {
    return null
  }
}

export function getUserInfo() {
  return Cookies.getJSON(cookie_name)
}

export function setToken (userInfo) {
  Cookies.set(cookie_name, userInfo)
}

export function removeToken () {
  Cookies.remove(cookie_name)
}
