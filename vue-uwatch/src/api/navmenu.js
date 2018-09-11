import request from '@/utils/request'

export function getNavList(query) {
  return request({
    url: 'v1/uw_nav',
    method: 'get',
    params: query
  })
}

export function getList(query) {
  return request({
    url: 'v1/uw_navmenu',
    method: 'get',
    params: query
  })
}

export function createNavMenu(data) {
  return request({
    url: 'v1/uw_navmenu',
    method: 'post',
    data
  })
}

export function updateNavMenu(data) {
  return request({
    url: 'v1/uw_navmenu/' + data.id,
    method: 'put',
    data
  })
}

export function deleteNavMenu(id) {
  return request({
    url: 'v1/uw_navmenu/' + id,
    method: 'delete'
  })
}

export function getNavAll(query) {
  return request({
    url: 'uw_navall',
    method: 'get',
    params: query
  })
}
