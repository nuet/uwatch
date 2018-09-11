import { getToken, setToken, removeToken } from '@/utils/auth'
import { getUserList, getTeamList, getDepartmentList, getRole } from '@/api/system_user'

const user = {
  state: {
    userInfo: {},
    user: '',
    status: '',
    code: '',
    token: getToken(),
    name: '',
    avatar: 'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif',
    introduction: '',
    roles: [],
    roleNames: [],
    setting: {
      articlePlatform: []
    },
    teamList: [],
    userList: [],
    departmentList: []
  },

  mutations: {
    SET_USERINFO: (state, userinfo) => {
      state.userInfo = userinfo
    },
    SET_USER: (state, user) => {
      state.user = user
    },
    SET_CODE: (state, code) => {
      state.code = code
    },
    SET_TOKEN: (state, token) => {
      state.token = token
    },
    SET_INTRODUCTION: (state, introduction) => {
      state.introduction = introduction
    },
    SET_SETTING: (state, setting) => {
      state.setting = setting
    },
    SET_STATUS: (state, status) => {
      state.status = status
    },
    SET_NAME: (state, name) => {
      state.name = name
    },
    SET_AVATAR: (state, avatar) => {
      state.avatar = avatar
    },
    SET_ROLES: (state, roles) => {
      state.roles = roles
    },
    SET_ROLENAMES: (state, roleNames) => {
      state.roleNames = roleNames
    },
    SET_USERLIST: (state, lists) => {
      state.userList = lists
    },
    SET_TEAMLIST: (state, lists) => {
      state.teamList = lists
    },
    SET_DEPARTMENTLIST: (state, lists) => {
      state.departmentList = lists
    }
  },

  actions: {
    // 获取用户信息
    GetInfo({ commit, state }) {
      return new Promise((resolve, reject) => {
        const value = JSON.parse(getToken())
        commit('SET_USERINFO', value.user)
        commit('SET_USER', value.user.name)
        commit('SET_TOKEN', value.user.token)
        getRole().then(response => {
          const data = response.data
          console.log(response.data.Role)
          commit('SET_ROLES', response.data.Role)
          commit('SET_ROLENAMES', response.data.RoleName)
          resolve(data)
        }).catch(error => {
          reject(error)
        })
      })
    },

    // 登出
    LogOut({ commit, state }) {
      commit('SET_TOKEN', '')
      commit('SET_ROLES', [])
      removeToken()
      window.location.href = process.env.BASE_API + '/logout?webUrl=' + window.location.href
    },

    // 前端 登出
    FedLogOut({ commit }) {
      return new Promise(resolve => {
        commit('SET_TOKEN', '')
        removeToken()
        resolve()
      })
    },

    // 动态修改权限
    ChangeRole({ commit }, role) {
      return new Promise(resolve => {
        commit('SET_ROLES', [role])
        commit('SET_TOKEN', role)
        setToken(role)
        resolve()
      })
    },

    // 获取员工列表
    GetUserList({ commit }) {
      return new Promise((resolve, reject) => {
        getUserList().then(response => {
          if (response.data !== 'ERROR') {
            commit('SET_USERLIST', JSON.parse(response.data))
          }
          resolve()
        }).catch(error => {
          reject(error)
        })
      })
    },

    // 获取UIC team列表
    GetTeamList({ commit }) {
      return new Promise((resolve, reject) => {
        getTeamList().then(response => {
          if (response.data !== 'ERROR') {
            commit('SET_TEAMLIST', JSON.parse(response.data))
          }
          resolve()
        }).catch(error => {
          reject(error)
        })
      })
    },

    // 获取部门列表
    GetDepartmentList({ commit }) {
      return new Promise((resolve, reject) => {
        getDepartmentList().then(response => {
          if (response.data !== 'ERROR') {
            commit('SET_DEPARTMENTLIST', JSON.parse(response.data))
          }
          resolve(JSON.parse(response.data))
        }).catch(error => {
          reject(error)
        })
      })
    }
  }
}

export default user
