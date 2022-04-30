import request from '@/utils/request'

export function getInstanceTargets (queryData) {
  return request({
    url: '/v1/instance/targets',
    method: 'get',
    params: queryData
  })
}

export function putInstanceTargets (putData) {
  return request({
    url: '/v1/instance/targets',
    method: 'put',
    data: putData
  })
}

