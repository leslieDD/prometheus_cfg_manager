import request from '@/utils/request'

export function getJobs () {
  return request({
    url: '/v1/jobs',
    method: 'get'
    // params: {}
  })
}

export function getJobsWithSplitPage (queryInfo) {
  return request({
    url: '/v1/jobs/split',
    method: 'get',
    params: queryInfo
  })
}

export function getJob (jobID) {
  return request({
    url: '/v1/job',
    method: 'get',
    params: {id: jobID}
  })
}

export function postJob (jobInfo) {
  return request({
    url: '/v1/job',
    method: 'post',
    data: jobInfo
  })
}

export function putJob (jobInfo) {
  return request({
    url: '/v1/job',
    method: 'put',
    data: jobInfo
  })
}

export function deleteJob (jobID) {
  return request({
    url: '/v1/job',
    method: 'delete',
    params: {id: jobID}
  })
}

export function swapJob (swapInfo) {
  return request({
    url: '/v1/job/swap',
    method: 'put',
    data: swapInfo
  })
}

export function publishJobs () {
  return request({
    url: '/v1/jobs/publish',
    method: 'post'
  })
}

export function enabledJob (jInfo) {
  return request({
    url: '/v1/jobs/status',
    method: 'put',
    data: jInfo
  })
}

export function clearIPForJob (clearInfo) {
  return request({
    url: '/v1/jobs/clear-ips',
    method: 'post',
    data: clearInfo
  })
}