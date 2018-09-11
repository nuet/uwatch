import request from '@/utils/request'

export function getList(query) {
  return request({
    url: 'v1/uw_counter',
    method: 'get',
    params: query
  })
}

export function createCounter(data) {
  return request({
    url: 'v1/uw_counter',
    method: 'post',
    data
  })
}

export function updateCounter(data) {
  return request({
    url: 'v1/uw_counter/' + data.id,
    method: 'put',
    data
  })
}

export function deleteCounter(id) {
  return request({
    url: 'v1/uw_counter/' + id,
    method: 'delete'
  })
}
