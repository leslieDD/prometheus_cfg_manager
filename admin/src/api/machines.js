import request from '@/utils/request'

export function getMechine (machineId) {
  return request({
    url: '/v1/machine',
    method: 'get',
    params: { id: machineId }
  })
}

export function getMachines () {
  return request({
    url: '/v1/machines',
    method: 'get'
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
