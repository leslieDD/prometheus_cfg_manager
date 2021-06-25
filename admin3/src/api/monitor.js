import request from '@/utils/request'

export function getTree () {
  return request({
    url: '/v1/tree',
    method: 'get'
    // params: {}
  })
}

export function getRuleDetail(info) {
  return request({
    url: '/v1/tree/node',
    method: 'get',
    params: info
  })
}
  
export function getDefaultLabels() {
  return request({
    url: '/v1/tree/default/lables',
    method: 'get'
  })
}

export function postNodeInfo(nodeInfo) {
  console.log('postNodeInfo', nodeInfo)
  return request({
    url: '/v1/tree/node',
    method: 'post',
    data: nodeInfo
  })
}

export function putNodeInfo(nodeInfo) {
  console.log('putNodeInfo', nodeInfo)
  return request({
    url: '/v1/tree/node',
    method: 'put',
    data: nodeInfo
  })
}