import request from '@/utils/request'

export function getList(query) {
  return request({
    url: 'v1/uw_graph',
    method: 'get',
    params: query
  })
}
export function UpdateAutoStatus(data) {
  return request({
    url: 'v1/uw_graph/' + data.Id,
    method: 'put',
    data
  })
}
export function createGraph(data) {
  return request({
    url: 'v1/uw_graph',
    method: 'post',
    data
  })
}

export function updateGraph(data) {
  return request({
    url: 'v1/uw_graph/' + data.id,
    method: 'put',
    data
  })
}

export function deleteGraph(id) {
  return request({
    url: 'v1/uw_graph/' + id,
    method: 'delete'
  })
}
