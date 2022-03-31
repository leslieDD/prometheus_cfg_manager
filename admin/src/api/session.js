import request from '@/utils/request'

export function getSession (getInfo) {
  return request({
    url: '/v1/manager/session',
    method: 'get',
    params: getInfo,
  })
}

export function delSession (info) {
  return request({
    url: '/v1/manager/session',
    method: 'delete',
    params: info,
  })
}

export function getSessionParams () {
  return request({
    url: '/v1/manager/session/params',
    method: 'get'
  })
}


export function updateSessionParams (data) {
  return request({
    url: '/v1/manager/session/params',
    method: 'put',
    data: data,
  })
}

