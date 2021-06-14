import request from '@/utils/request'

export function restartSrv () {
  return request({
    url: '/v1/restart',
    method: 'post'
  })
}
