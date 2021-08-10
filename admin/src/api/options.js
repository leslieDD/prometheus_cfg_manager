import request from '@/utils/request'

export function getOptions () {
  return request({
    url: '/v1/options',
    method: 'get'
  })
}

export function putOptions (options) {
  return request({
    url: '/v1/options',
    method: 'put',
    data: options
  })
}
