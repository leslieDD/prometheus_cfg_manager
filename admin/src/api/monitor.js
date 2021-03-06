import request from '@/utils/request'

export function getTree () {
  return request({
    url: '/v1/tree',
    method: 'get'
    // params: {}
  })
}

export function getRuleDetail (info) {
  return request({
    url: '/v1/tree/node',
    method: 'get',
    params: info
  })
}

export function getDefaultLabels () {
  return request({
    url: '/v1/tree/default/lables',
    method: 'get'
  })
}

export function getDefaultEnableLables () {
  return request({
    url: '/v1/tree/default/enable/lables',
    method: 'get'
  })
}

export function postNodeInfo (nodeInfo) {
  return request({
    url: '/v1/tree/node',
    method: 'post',
    data: nodeInfo
  })
}

export function putNodeInfo (nodeInfo) {
  return request({
    url: '/v1/tree/node',
    method: 'put',
    data: nodeInfo
  })
}

export function deleteNodeLable (labelInfo, lType) {
  return request({
    url: '/v1/tree/node/label',
    method: 'delete',
    data: labelInfo,
    params: lType
  })
}

export function createTreeNode (nodeInfo) {
  return request({
    url: '/v1/tree/create/node',
    method: 'post',
    data: nodeInfo
  })
}

export function updateTreeNode (nodeInfo) {
  return request({
    url: '/v1/tree/update/node',
    method: 'put',
    data: nodeInfo
  })
}

export function removeTreeNode (nodeInfo, skipSelf) {
  return request({
    url: '/v1/tree/remove/node',
    method: 'delete',
    data: nodeInfo,
    params: skipSelf
  })
}


export function rulesPublish () {
  return request({
    url: '/v1/rules/publish',
    method: 'post'
  })
}

export function emptyRulesPublish () {
  return request({
    url: '/v1/rules/empty/publish',
    method: 'post'
  })
}

export function disableTreeNode (disableData) {
  return request({
    url: '/v1/tree/node/status',
    method: 'put',
    data: disableData
  })
}

export function uploadYamlFile (formData, gInfo) {
  return request({
    url: '/v1/tree/upload/file/yaml',
    method: 'post',
    data: formData,
    params: gInfo,
  })
}