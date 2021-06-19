import request from '@/utils/request'

export function publish () {
  return request({
    url: '/v1/publish',
    method: 'post'
    // params: {}
  })
}

export function loadFileContent(fileInfo) {
    return request({
        url: '/v1/load/file-content',
        method: 'post',
        data: fileInfo
    })
}

export function loadAllFile() {
    return request({
        url: '/v1/load/all-file-list',
        method: 'get'
    })
}