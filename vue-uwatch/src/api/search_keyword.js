import request from '@/utils/request'

export function getSearchKeyWord(query) {
  console.log(query)
  return request({
    url: 'v1/search_keyword',
    method: 'get',
    params: query
  })
}

export function getGraphData(query, v) {
  if (v) {
    return request({
      url: '/findFalconData',
      method: 'get',
      params: query
    })
  } else {
    return request({
      url: '/uw_graph_data',
      method: 'get',
      params: query
    })
  }
}

export function getOneGraphData(query, v) {
  if (v) {
    return request({
      url: '/findOneFalconData',
      method: 'get',
      params: query
    })
  } else {
    return request({
      url: '/findOneFalconData',
      method: 'get',
      params: query
    })
  }
}

export function getCounter(query) {
  return request({
    url: '/getCounter',
    method: 'get',
    params: query
  })
}

export function getFalconCounter(query) {
  return request({
    url: '/getFalconCounter',
    method: 'get',
    params: query
  })
}

export function getGraphCounter(query) {
  return request({
    url: '/getGraphCounter',
    method: 'get',
    params: query
  })
}

export function getCmdbHost(query) {
  return request({
    url: '/getCmdbHost',
    method: 'get',
    params: query
  })
}
