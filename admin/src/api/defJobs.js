import request from '@/utils/request'

export function getDefJobsWithSplitPage (queryInfo) {
  return request({
    url: '/v1/jobs/default/split',
    method: 'get',
    params: queryInfo
  })
}

export function postDefJob (jobInfo) {
  return request({
    url: '/v1/job/default',
    method: 'post',
    data: jobInfo
  })
}

export function putDefJob (jobInfo) {
  return request({
    url: '/v1/job/default',
    method: 'put',
    data: jobInfo
  })
}

export function deleteDefJob (jobID) {
  return request({
    url: '/v1/job/default',
    method: 'delete',
    params: {id: jobID}
  })
}

export function enabledDefaultJob (info) {
  return request({
    url: '/v1/job/default/status',
    method: 'put',
    data: info
  })
}

export function publishJobs () {
  return request({
    url: '/v1/job/default/publish',
    method: 'post'
  })
}