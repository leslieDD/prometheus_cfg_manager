// getBaseLabels, putBaseLabels, postBaseLabels, searchBaseLabels
import request from '@/utils/request'

export function getBaseLabels (info) {
  return request({
    url: '/v1/base/labels',
    method: 'get',
    params: info
  })
}

export function putBaseLabels (info) {
  return request({
    url: '/v1/base/labels',
    method: 'put',
    data: info
  })
}

export function postBaseLabels (info) {
  return request({
    url: '/v1/base/labels',
    method: 'post',
    data: info
  })
}

export function deleteBaseLabels (info) {
  return request({
    url: '/v1/base/labels',
    method: 'delete',
    params: info
  })
}


  
