import request from '@/utils/request'

export function getAllReLabels () {
  return request({
    url: '/v1/base/relabels/all',
    method: 'get'
  })
}

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

export function enabledRelabelCfg (info) {
  return request({
    url: '/v1/base/relabels/status',
    method: 'put',
    data: info
  })
}

export function checkViewCodePriv () {
  return request({
    url: '/v1/base/relabels/check/view-code-priv',
    method: 'get'
  })
}