import request from '@/utils/request'

export function login (userInfo) {
  return request({
    url: '/login',
    method: 'post',
    data: userInfo
  })
}

export function logout (userInfo) {
  return request({
    url: '/v1/logout',
    method: 'post',
    data: userInfo
  })
}
