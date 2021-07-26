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