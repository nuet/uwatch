import request from '@/utils/request'

export function getList(query) {
  return request({
    url: 'v1/uw_nav',
    method: 'get',
    params: query
  })
}

export function createNav(data) {
  return request({
    url: 'v1/uw_nav',
    method: 'post',
    data
  })
}

export function updateNav(data) {
  return request({
    url: 'v1/uw_nav/' + data.id,
    method: 'put',
    data
  })
}

export function deleteNav(id) {
  return request({
    url: 'v1/uw_nav/' + id,
    method: 'delete'
  })
}
