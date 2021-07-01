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
    method: 'post',
    data: data
  })
}
