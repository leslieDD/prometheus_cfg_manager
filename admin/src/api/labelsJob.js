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

export function getJobMachines (jobID) {
  return request({
    url: '/v1/job/machines',
    method: 'get',
    params: { id: jobID }
  })
}

export function getJobMachinesForGroup (groupID) {
  return request({
    url: '/v1/job/group/machines',
    method: 'get',
    params: { id: groupID }
  })
}

export function putJobMachinesForGroup (gID, jobID, groupInfo) {
  return request({
    url: '/v1/job/group/machines',
    method: 'put',
    params: { id: gID, job_id: jobID },
    data: groupInfo
  })
}

export function getGroupLabels (gInfo) {
  return request({
    url: '/v1/job/group/labels',
    method: 'get',
    params: gInfo
  })
}

export function postGroupLabels (gID, jobGroupLablesInfo) {
  return request({
    url: '/v1/job/group/labels',
    method: 'post',
    params: { id: gID },
    data: jobGroupLablesInfo
  })
}

export function putGroupLabels (gID, jobGroupLablesInfo) {
  return request({
    url: '/v1/job/group/labels',
    method: 'put',
    params: { id: gID },
    data: jobGroupLablesInfo
  })
}

export function delGroupLabels (delInfo) {
  return request({
    url: '/v1/job/group/labels',
    method: 'delete',
    params: delInfo
  })
}

export function getGroupMachine (jobGroupID) {
  return request({
    url: '/v1/job/group/machine',
    method: 'get',
    params: { id: jobGroupID }
  })
}

export function putGroupMachine (jobGroupMachinesInfo) {
  return request({
    url: '/v1/job/group/machine',
    method: 'put',
    params: jobGroupMachinesInfo
  })
}

export function getAllIPAndLabels (info) {
  return request({
    url: '/v1/job/group/machines-and-labels',
    method: 'get',
    params: info
  })
}

export function enabledJobGroup (info) {
  return request({
    url: '/v1/job/group/status',
    method: 'put',
    data: info
  })
}

export function enabledJobGroupLables (info) {
  return request({
    url: '/v1/job/group/labels/status',
    method: 'put',
    data: info
  })
}

export function delEmptySubJob (info) {
  return request({
    url: '/v1/job/group/subgroup/emtpy',
    method: 'delete',
    params: info
  })
}