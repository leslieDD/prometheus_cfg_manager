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

export function batchDeleteMachine(ids){
  return request({
    url: '/v1/machines/selection',
    method: 'delete',
    data: ids
  })
}