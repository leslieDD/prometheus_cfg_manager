// getBaseLabels, putBaseLabels, postBaseLabels, searchBaseLabels
import request from '@/utils/request'

export function getIDC (info) {
  return request({
    url: '/v1/idc',
    method: 'get',
    params: info
  })
}

export function getIDCs (info) {
  return request({
    url: '/v1/idcs',
    method: 'get',
    params: info
  })
}

export function postIDC (info) {
  return request({
    url: '/v1/idcs',
    method: 'post',
    params: info
  })
}

export function putIDC (info) {
  return request({
    url: '/v1/idc',
    method: 'put',
    params: info
  })
}

export function deleteIDC (info) {
  return request({
    url: '/v1/idc',
    method: 'delete',
    params: info
  })
}

export function deleteIDC (info) {
  return request({
    url: '/v1/idc',
    method: 'delete',
    params: info
  })
}

export function getIDCTree (info) {
  return request({
    url: '/v1/idc/tree',
    method: 'get',
    params: info
  })
}

export function getLine (info) {
  return request({
    url: '/v1/idc/line',
    method: 'get',
    params: info
  })
}

export function getLines (info) {
  return request({
    url: '/v1/idc/lines',
    method: 'get',
    params: info
  })
}

export function postLine (info) {
  return request({
    url: '/v1/idc/line',
    method: 'post',
    params: info
  })
}

export function putLine (info) {
  return request({
    url: '/v1/idc/line',
    method: 'put',
    params: info
  })
}

export function delLine (info) {
  return request({
    url: '/v1/idc/line',
    method: 'delete',
    params: info
  })
}

export function getLineIpAddrs (info) {
  return request({
    url: '/v1/idc/line/ipaddrs',
    method: 'get',
    params: info
  })
}

export function putLineIpAddrs (info) {
  return request({
    url: '/v1/idc/line/ipaddrs',
    method: 'put',
    params: info
  })
}

