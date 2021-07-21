import request from '@/utils/request'

export function resetSystem (GrantInfo) {
  return request({
    url: '/v1/operate/reset/system',
    method: 'post',
    data: GrantInfo
  })
}

export function preResetSystem (GrantInfo) {
  return request({
    url: '/v1/operate/reset/secret',
    method: 'post',
    data: GrantInfo
  })
}

export function getOperateLog (queryInfo) {
  return request({
    url: '/v1/operate/log',
    method: 'get',
    params: queryInfo
  })
}