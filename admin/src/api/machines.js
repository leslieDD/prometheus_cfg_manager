import request from '@/utils/request'

export function getMechine (machineId) {
  return request({
    url: '/v1/machine',
    method: 'get',
    params: { id: machineId }
  })
}

export function getMachines (info) {
  return request({
    url: '/v1/machines',
    method: 'get',
    params: info
  })
}

export function postMachine (data) {
  return request({
    url: '/v1/machine',
    method: 'post',
    data
  })
}

export function putMachine (data) {
  return request({
    url: '/v1/machine',
    method: 'put',
    data
  })
}

export function deleteMachine (machineId) {
  return request({
    url: '/v1/machine',
    method: 'delete',
    params: { id: machineId }
  })
}

export function enabledMachine (mInfo) {
  return request({
    url: '/v1/machine/status',
    method: 'put',
    data: mInfo
  })
}

export function allIPList () {
  return request({
    url: '/v1/machines/all',
    method: 'get'
  })
}

export function batchDeleteMachine (ids) {
  return request({
    url: '/v1/machines/selection',
    method: 'delete',
    data: ids
  })
}

export function batchEnableMachine (ids) {
  return request({
    url: '/v1/machines/selection/enable',
    method: 'put',
    data: ids
  })
}

export function batchDisableMachine (ids) {
  return request({
    url: '/v1/machines/selection/disable',
    method: 'put',
    data: ids
  })
}

export function batchImportIpAddrsWeb (data) {
  return request({
    url: '/v1/machines/batch/import',
    method: 'put',
    data: data,
  })
}

export function batchImportDomainWeb (data) {
  return request({
    url: '/v1/machines/batch/import/domain',
    method: 'put',
    data: data,
  })
}

export function updatePosition () {
  return request({
    url: '/v1/machines/position',
    method: 'put'
  })
}

export function outputAllIP () {
  return request({
    url: '/v1/output/ip/list',
    method: 'get'
  })
}
