import request from '@/utils/request'

export function getTree () {
  return request({
    url: '/v1/tree',
    method: 'get'
    // params: {}
  })
}

export function getRuleDetail (info) {
    return request({
      url: '/v1/tree/node',
      method: 'get',
      params: info
    })
  }
  