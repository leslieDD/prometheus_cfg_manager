import request from '@/utils/request'

export function loadTxt () {
  return request({
    url: '/v1/txt/programer/say',
    method: 'get'
  })
}

export function chgPasswd(pwdInfo){
  return request({
    url: '/v1/manager/user/password',
    method: 'post',
    data: pwdInfo
  })
}

export function getMenuPriv(userInfo){
  return request({
    url: '/v1/manager/user/menu/priv',
    method: 'get',
    params: userInfo
  })
}

export function logout(){
  return request({
    url: '/v1/logout',
    method: 'post'
  })
}