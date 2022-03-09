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


