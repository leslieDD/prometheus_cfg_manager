import request from '@/utils/request'

export function getReLabels (info) {
  return request({
    url: '/v1/base/relabels',
    method: 'get',
    params: info
  })
}

export function putReLabels (info) {
  return request({
    url: '/v1/base/relabels',
    method: 'put',
    data: info
  })
}

export function postReLabels (info) {
  return request({
    url: '/v1/base/relabels',
    method: 'post',
    data: info
  })
}

export function deleteReLabels (info) {
  return request({
    url: '/v1/base/relabels',
    method: 'delete',
    params: info
  })
}

export function putReLabelsCode (info) {
  return request({
    url: '/v1/base/relabels/code',
    method: 'put',
    data: info
  })
}
