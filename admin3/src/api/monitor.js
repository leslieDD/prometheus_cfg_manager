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
  return request({
    url: '/v1/tree/node',
    method: 'post',
    data: nodeInfo
  })
}

export function putNodeInfo(nodeInfo) {
  return request({
    url: '/v1/tree/node',
    method: 'put',
    data: nodeInfo
  })
}

export function deleteNodeLable(labelInfo, lType) {
return request({
    url: '/v1/tree/node/label',
    method: 'delete',
    data: labelInfo,
    params: lType
})
}

export function createTreeNode(nodeInfo) {
return request({
    url: '/v1/tree/create/node',
    method: 'post',
    data: nodeInfo
})
}

export function updateTreeNode(nodeInfo) {
  return request({
    url: '/v1/tree/update/node',
    method: 'put',
    data: nodeInfo
  })
}

export function removeTreeNode(nodeInfo) {
return request({
    url: '/v1/tree/remove/node',
    method: 'delete',
    data: nodeInfo
})
}


export function rulesPublish() {
return request({
    url: '/v1/rules/publish',
    method: 'post'
})
}