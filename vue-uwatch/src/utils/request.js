import axios from 'axios'
import { Message } from 'element-ui'
import store from '@/store'

// create an axios instance
const service = axios.create({
  baseURL: process.env.BASE_API, // api的base_url
  timeout: 5000 // request timeout
})

// request interceptor
service.interceptors.request.use(config => {
  // Do something before request is sent
  if (store.state.user.user && store.state.user.token) {
    config.headers.Authorization = `token ${store.state.user.token}`
  }
  return config
}, error => {
  // Do something with request error
  console.log(error) // for debug
  Promise.reject(error)
})

// respone interceptor
service.interceptors.response.use(
  response => {
    const res = response.data
    if (res.code !== 200 && res.code !== 201) {
      // Message({
      //   message: res.msg,
      //   type: 'error',
      //   duration: 5 * 1000
      // })
      // 50008:非法的token 50012:其他客户端登录了  50014:Token 过期了
      if (res.code === 50008 || res.code === 50012 || res.code === 50014) {
        store.dispatch('FedLogOut').then(() => {
          location.reload()// 为了重新实例化vue-router对象 避免bug
        })
      }
      return Promise.reject('error')
    } else {
      console.log(res)
      return res
    }
  },
  error => {
    console.log('err' + error)// for debug
    Message({
      message: error.message,
      type: 'error',
      duration: 5 * 1000
    })
    return Promise.reject(error)
  })

export default service
