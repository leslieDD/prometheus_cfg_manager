import request from '@/utils/request'

export function getProTmpl () {
  return request({
    url: '/v1/prometheus/tmpl',
    method: 'get'
  })
}

export function putProTmpl (data) {
  return request({
    url: '/v1/prometheus/tmpl',
    method: 'put',
    data: data
  })
}

export function getGoStruct () {
  return request({
    url: '/v1/prometheus/struct',
    method: 'get'
  })
}
