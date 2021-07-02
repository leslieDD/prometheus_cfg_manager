import request from '@/utils/request'

export function getJobGroup (info) {
  return request({
    url: '/v1/job/group',
    method: 'get',
    params: info
  })
}

export function putJobGroup (jobSubGroupInfo) {
  return request({
    url: '/v1/job/group',
    method: 'put',
    data: jobSubGroupInfo
  })
}

export function addJobGroup (jobSubGroupInfo) {
  return request({
    url: '/v1/job/group',
    method: 'post',
    data: jobSubGroupInfo
  })
}

export function delJobGroup (info) {
  return request({
    url: '/v1/job/group',
    method: 'delete',
    params: info
  })
}

export function getJobMachines (jobGroupID) {
  return request({
    url: '/v1/job/machines',
    method: 'get',
    params: {id: jobGroupID}
  })
}

export function getGroupLabels (jobGroupID) {
  return request({
    url: '/v1/job/group/labels',
    method: 'get',
    params: {id: jobGroupID}
  })
}

export function putGroupLabels (jobGroupLablesInfo) {
  return request({
    url: '/v1/job/group/labels',
    method: 'put',
    data: jobGroupLablesInfo
  })
}

export function getGroupMachine (jobGroupID) {
  return request({
    url: '/v1/job/group/machine',
    method: 'get',
    params: {id: jobGroupID}
  })
}

export function putGroupMachine (jobGroupMachinesInfo) {
  return request({
    url: '/v1/job/group/machine',
    method: 'put',
    params: jobGroupMachinesInfo
  })
}