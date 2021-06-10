import request from '@/utils/request'

export function publish () {
  return request({
    url: '/v1/publish',
    method: 'post'
    // params: {}
  })
}
