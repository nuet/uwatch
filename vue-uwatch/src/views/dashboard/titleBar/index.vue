<template>
  <el-menu
    class="el-menu-demo"
    mode="horizontal"
    background-color="rgb(29, 62, 116)"
    text-color="#fff"
    active-text-color="#ffffff"
    style="position: relative;z-index: 999;height:60px;" v-clickoutside="handleClose">
    <div style="float:left;padding-left:20px;width:auto;margin-right:45px;margin-top: 8px;margin-bottom: 6px;cursor: pointer" @click="goto('dashboard')">
      <img :src="logo" alt=""/>
    </div>
    <el-menu-item index="3" style="float:left;cursor: pointer;" @click="showMenu()"><i class="el-icon-tickets" aria-hidden="true" style="color: #fff"></i>常用监控</el-menu-item>
    <el-menu-item index="2" @click="goto('monitor')"><i class="el-icon-view" aria-hidden="true" style="color: #fff"></i>监控大屏</el-menu-item>
    <div style="float:right;padding-top:12px;padding-right:20px;">
      <el-button type="text" style="color: #fff">{{this.$store.state.user.user}}</el-button>
      <el-button type="text" style="color:#909399"  @click="logout">退出登录</el-button>
      <span style="color: #fff;">|</span>
      <a href="#admin" target="_blank"   type="text" size="small" >
        <el-button   type="text"  style="color: #fff">管理后台</el-button></a>
    </div>
    <el-card class="box-card" v-show="showMenuList" style="width: 100%;height: auto;background-color: #fff">
      <dl class="nav-dropdown-list ng-scope" v-for=" item in navList" :key="item.key" :name="item.name">
        <dt>
          <a :href="item[0].purl" target="_blank" class="nav-dropdown-link ng-binding ng-scope"> {{item[0].pname}}</a>
        </dt>
        <dd v-for="list in item">
          <a :href="list.url" target="_blank" class="nav-dropdown-link ng-binding ng-scope"><i class="el-icon-d-arrow-right" aria-hidden="true" style="padding-right:5px;color:#167be0;">&nbsp;</i>{{list.name}}</a>
        </dd>
      </dl>
    </el-card>
  </el-menu>
</template>
<script type="text/javascript">
  import logo from '@/assets/dashboard/logo.png'
  import store from '../../../store'
  import '@/assets/dashboard/bootstrap/css/bootstrap.css' // global css
  import '@/assets/dashboard/style.css'
  import { getNavAll } from '@/api/navmenu'
  const clickoutside = {
    // 初始化指令
    bind(el, binding, vnode) {
      function documentHandler(e) {
        // 这里判断点击的元素是否是本身，是本身，则返回
        if (el.contains(e.target)) {
          return false
        }
        // 判断指令中是否绑定了函数
        if (binding.expression) {
          // 如果绑定了函数 则调用那个函数，此处binding.value就是handleClose方法
          binding.value(e)
        }
      }
      // 给当前元素绑定个私有变量，方便在unbind中可以解除事件监听
      el.__vueClickOutside__ = documentHandler
      document.addEventListener('click', documentHandler)
    },
    update() {},
    unbind(el, binding) {
      // 解除事件监听
      document.removeEventListener('click', el.__vueClickOutside__)
      delete el.__vueClickOutside__
    }
  }
  export default {
    props: {},
    data() {
      return {
        logo,
        activeIndex2: '1',
        showMenuList: false,
        listQuery: {
          page: 1,
          limit: 10,
          sort: '-Id',
          query: ''
        },
        navList: []
      }
    },
    directives: { clickoutside },
    mounted() {
      getNavAll(this.listQuery).then(response => {
        console.log('getNavAll===>', response.data.list)
        this.navList = response.data.list
        this.listLoading = false
      })
    },
    methods: {
      tografana() {
        this.$emit('showgrafana', true)
        const loading = this.$loading({
          lock: true,
          text: 'Loading',
          spinner: 'el-icon-loading',
          background: 'rgba(0, 0, 0, 0.7)'
        })
        setTimeout(() => {
          loading.close()
        }, 2000)
      },
      // 退出
      logout() {
        this.$confirm('此操作将退出登录, 是否继续?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          store.dispatch('LogOut').then(() => {
            location.reload()// 为了重新实例化vue-router对象 避免bug
          })
        }).catch(() => {})
      },
      showdata(url) {
        this.$emit('showdata', url)
        const loading = this.$loading({
          lock: true,
          text: 'Loading',
          spinner: 'el-icon-loading',
          background: 'rgba(0, 0, 0, 0.7)'
        })
        setTimeout(() => {
          loading.close()
        }, 2000)
      },
      goto(path) {
        if (this.showMenuList === true) {
          this.showMenuList = false
        }
        path = '/' + path
        this.$router.push({ path: path })
      },
      handleClose(e) {
        if (this.showMenuList === true) {
          this.showMenuList = false
        }
      },
      showMenu() {
        if (this.showMenuList === true) {
          this.showMenuList = false
        } else {
          this.showMenuList = true
        }
      },
      cancelShowMenuList() {
        this.showMenuList = false
      }
    }
  }
</script>
<style scoped>
  .main-container {
    min-height: 100%;
    -webkit-transition: margin-left 0.28s;
    transition: margin-left 0.28s;
    margin-left: 0px;
  }
  .top-low {
    margin-top: 5px
  }
  .el-table th{
    background-color: #eef1f6;
  }
  /*.el-table thead {*/
    /*color:#606266;*/
  /*}*/
  /*.el-step__title.is-process{*/
    /*color:#409eff;*/
  /*}*/
  .el-step__title {
    font-size: 16px;
    line-height: 30px;
  }
  .navbar{
    margin-bottom: 0;
    z-index: 9;
    width: 100%;
    position: relative;
    background: #fff;
    font-size: 13px;
  }
  .navbar-inner {
    background: 0 0;
    border-radius: 0;
    border: none;
    box-shadow: none;
    width: 85%;
    margin: auto;
    /*padding: 5px;*/
  }
  .navbar .container {
    width: 100%;
    min-width: 960px;
    margin: 0 auto;
    max-width: 100%;
  }
  /*.navbar .brand {*/
    /*float:left;width: auto;padding: 5px 20px;height: 34px;line-height: 34px;color: #ccc;font-weight: 700;display: block;margin-left: -20px;*/
    /*text-decoration:none;*/
    /*font-size: 32px;*/
  /*}*/
  .navbar .nav.pull-right {
    float: left;
    margin-right: 0;
  }
  .pull-right {
    float: left !important;
  }
  .navbar .nav>li {
    float: left;
  }
  .nav > li > a {
    position: relative;
    display: block;
    padding: 10px 15px;
    font-size: 13px;
  }
  /*.nav > li > a:hover {*/
    /*color: #409eff;*/
  /*}*/
  .el-menu--horizontal {
    border-bottom: none;
  }

  .nav-dropdown-list {
    position: relative;
    display: inline-block;
    -webkit-box-flex: initial;
    -ms-flex: initial;
    flex: initial;
    margin: 0 40px 30px;
    vertical-align: top;
  }
  .nav-dropdown-list:before {
    content: "";
    position: absolute;
    display: block;
    top: 40px;
    bottom: 0;
    left: -44px;
    width: 1px;
    border-left: 1px solid #e1e1e1;
  }
  .nav-dropdown-list .nav-dropdown-link {
    position: relative;
    display: block;
    width: 125px;
    padding: 8px 0;
    padding-left: 5px;
    color: #4c4c4c;
  }
  .ng-scope  a:hover{
    color: #167be0;
    text-decoration: none;
  }
  .el-card, .el-message {
     border-radius: 0px;
    overflow: hidden;
  }
  .el-menu--horizontal>.el-menu-item {
    float: left;
    height: 60px;
    line-height: 60px;
    margin: 0;
    border-bottom: 0px solid transparent;
    color: #909399;
  }
  .el-menu-item [class^=el-icon-] {
    margin-right: 0px;
    margin-top: -3px;
    width: 24px;
    text-align: center;
    font-size: 17px;
    vertical-align: middle;
  }
</style>
