import request from '@/utils/request'

export function getJobs () {
  return request({
    url: '/v1/jobs',
    method: 'get'
    // params: {}
  })
}
