import request from '@/utils/request'

export function ctlCreate () {
  return request({
    url: '/v1/control/create',
    method: 'post'
  })
}

export function ctlReloadPrometheus () {
  return request({
    url: '/v1/control/reload',
    method: 'post'
  })
}

export function ctlCreateAndReload () {
  return request({
    url: '/v1/control/create/and/reload',
    method: 'post'
  })
}

