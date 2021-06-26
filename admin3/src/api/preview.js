import request from '@/utils/request'

export function preview () {
  return request({
    url: '/v1/preview',
    method: 'get'
  })
}

export function preRulesView () {
    return request({
      url: '/v1/preview/rules',
      method: 'get'
    })
  }
  