import request from '@/utils/request'

export function restartSrv () {
  return request({
    url: '/v1/reload',
    method: 'post'
  })
}

export function restartDefSrv () {
  return request({
    url: '/v1/def/reload',
    method: 'post'
  })
}
