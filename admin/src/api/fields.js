import request from '@/utils/request'

export function getBaseFields (info) {
  return request({
    url: '/v1/base/fields',
    method: 'get',
    params: info
  })
}

export function putBaseFields (info) {
  return request({
    url: '/v1/base/fields',
    method: 'put',
    data: info
  })
}

export function postBaseFields (info) {
  return request({
    url: '/v1/base/fields',
    method: 'post',
    data: info
  })
}

export function deleteBaseFields (info) {
  return request({
    url: '/v1/base/fields',
    method: 'delete',
    params: info
  })
}

export function enabledBaseFields (info) {
  return request({
    url: '/v1/base/fields/status',
    method: 'put',
    data: info
  })
}
  
