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
    url: '/v1/idc',
    method: 'post',
    data: info
  })
}

export function putIDC (info) {
  return request({
    url: '/v1/idc',
    method: 'put',
    data: info
  })
}

export function putIDCRemark (info) {
  return request({
    url: '/v1/idc/remark',
    method: 'put',
    data: info
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
    data: info
  })
}

export function putLine (info) {
  return request({
    url: '/v1/idc/line',
    method: 'put',
    data: info
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
    data: info
  })
}

export function updateAllIPAddrsNetInfo (info) {
  return request({
    url: '/v1/idc/update/netinfo/all',
    method: 'put',
    data: info
  })
}

export function updatePartIPAddrsNetInfo (info) {
  return request({
    url: '/v1/idc/update/netinfo/part',
    method: 'put',
    data: info
  })
}

export function createLabelForAllIPs (info) {
  return request({
    url: '/v1/idc/create/label',
    method: 'put',
    data: info
  })
}

export function createIPForJob (data) {
  return request({
    url: '/v1/idc/expand',
    method: 'post',
    data: data
  })
}

export function putIdcExpand (data) {
  return request({
    url: '/v1/idc/expand/switch',
    method: 'put',
    data: data
  })
}

export function putLineExpand (data) {
  return request({
    url: '/v1/idc/line/expand/switch',
    method: 'put',
    data: data
  })
}

export function putIdcView (data) {
  return request({
    url: '/v1/idc/view/switch',
    method: 'put',
    data: data
  })
}

export function putLineView (data) {
  return request({
    url: '/v1/idc/line/view/switch',
    method: 'put',
    data: data
  })
}

export function outportXls() {
  return request({
    url: '/v1/idc/tree/xls',
    method: 'get',
  })
}
