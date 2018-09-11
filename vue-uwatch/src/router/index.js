import Vue from 'vue'
import Router from 'vue-router'
const _import = require('./_import_' + process.env.NODE_ENV)
// in development-env not use lazy-loading, because lazy-loading too many pages will cause webpack hot update too slow. so only in production use lazy-loading;
// detail: https://panjiachen.github.io/vue-element-admin-site/#/lazy-loading

Vue.use(Router)

/* Layout */
import Layout from '../views/layout/Layout'
import LayoutWeb from '../views/layout/LayoutWeb'

/** note: submenu only apppear when children.length>=1
*   detail see  https://panjiachen.github.io/vue-element-admin-site/#/router-and-nav?id=sidebar
**/

/**
* hidden: true                   if `hidden:true` will not show in the sidebar(default is false)
* alwaysShow: true               if set true, will always show the root menu, whatever its child routes length
*                                if not set alwaysShow, only more than one route under the children
*                                it will becomes nested mode, otherwise not show the root menu
* redirect: noredirect           if `redirect:noredirect` will no redirct in the breadcrumb
* name:'router-name'             the name is used by <keep-alive> (must set!!!)
* meta : {
    roles: ['admin','editor']     will control the page roles (you can set multiple roles)
    title: 'title'               the name show in submenu and breadcrumb (recommend set)
    icon: 'svg-name'             the icon show in the sidebar,
    noCache: true                if true ,the page will no be cached(default is false)
  }
**/
export const constantRouterMap = [
  { path: '/login', component: _import('login/index'), hidden: true },
  { path: '/authredirect', component: _import('login/authredirect'), hidden: true },
  { path: '/404', component: _import('errorPage/404'), hidden: true },
  { path: '/401', component: _import('errorPage/401'), hidden: true },
  {
    path: '',
    component: LayoutWeb,
    redirect: '/dashboard',
    children: [{
      path: 'dashboard',
      component: _import('dashboard/index'),
      name: 'Uwatch',
      meta: { title: 'Uwatch', icon: 'international', noCache: true }
    }]
  },
  { path: '/dashboard/search', component: _import('dashboard/search'), hidden: true },
  { path: '/dashboard/detail', component: _import('dashboard/detail'), hidden: true },
  {
    path: '/admin',
    component: Layout,
    redirect: '/admin/index',
    children: [{
      path: 'index',
      component: _import('admin/index'),
      name: 'admin',
      meta: { title: '首页', icon: 'dashboard', noCache: true }
    }]
  },
  { path: '/monitor', component: _import('monitor/index'), hidden: true }
]

export default new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRouterMap
})

export const asyncRouterMap = [
  {
    path: '/autofill',
    component: Layout,
    redirect: 'noredirect',
    name: 'autofill',
    meta: {
      title: '故障自愈',
      icon: 'example'
    },
    children: [
      { path: 'tasklist', icon: 'tab', component: _import('recoverytask/taskList'), name: 'taskList', meta: { title: '自愈详情' }},
      { path: '/task/:id', icon: 'tab', component: _import('recoverytask/taskDetail'), name: 'taskDetail', meta: { title: '任务详情' }, hidden: true, noCache: true },
      {
        path: '/autofill/access',
        component: _import('autofill/access/index'),
        name: 'autofill-access',
        meta: {
          title: '接入自愈'
        },
        children: [
          { path: 'list', component: _import('autofill/access/list'), name: 'autofill-access-list', meta: { title: '接入自愈' }},
          { path: 'add', component: _import('autofill/access/add'), name: 'autofill-access-add', meta: { title: '添加自愈方案' }, hidden: true },
          { path: 'edit/:id', component: _import('autofill/access/add'), name: 'autofill-access-edit', meta: { title: '编辑自愈方案' }, hidden: true },
          { path: 'detection/:id', component: _import('autofill/access/detection'), name: 'autofill-access-detection', meta: { title: '检测自愈方案' }, hidden: true }
        ]
      }
    ]
  },
  {
    path: '/configure',
    component: Layout,
    redirect: 'noredirect',
    hidden: true,
    name: 'configure',
    meta: {
      title: '参数配置',
      icon: 'chart'
    },
    children: [
      { path: 'graph', icon: 'chart', component: _import('configure/graph/list'), name: 'configure-graph-list', meta: { title: '图表配置', icon: 'chart', noCache: true }}
    ]
  },
  {
    path: '/work',
    component: Layout,
    redirect: 'noredirect',
    hidden: false,
    name: 'work',
    meta: {
      title: '前端管理',
      icon: 'component'
    },
    children: [
      { path: 'counter', icon: 'tab', component: _import('counter/counterList'), name: 'counter', meta: { title: '配置Counter' }},
      { path: 'counter/add', icon: 'tab', component: _import('counter/counterList'), name: 'counter-add', meta: { title: '新建Counter' }, hidden: true },
      { path: 'nav', icon: 'tab', component: _import('nav/navList'), name: 'nav', meta: { title: '一级导航' }},
      { path: 'menu', icon: 'tab', component: _import('nav/menuList'), name: 'menu', meta: { title: '二级导航' }}
    ]
  },
  {
    path: '/icon',
    component: Layout,
    hidden: true,
    children: [{
      path: 'index',
      component: _import('svg-icons/index'),
      name: 'icons',
      meta: { title: 'icons', icon: 'icon', noCache: true }
    }]
  },

  {
    path: '/components',
    component: Layout,
    redirect: 'noredirect',
    hidden: true,
    name: 'component-demo',
    meta: {
      title: 'components',
      icon: 'component'
    },
    children: [
      { path: 'tinymce', component: _import('components-demo/tinymce'), name: 'tinymce-demo', meta: { title: 'tinymce' }},
      { path: 'markdown', component: _import('components-demo/markdown'), name: 'markdown-demo', meta: { title: 'markdown' }},
      { path: 'json-editor', component: _import('components-demo/jsonEditor'), name: 'jsonEditor-demo', meta: { title: 'jsonEditor' }},
      { path: 'dnd-list', component: _import('components-demo/dndList'), name: 'dndList-demo', meta: { title: 'dndList' }},
      { path: 'splitpane', component: _import('components-demo/splitpane'), name: 'splitpane-demo', meta: { title: 'splitPane' }},
      { path: 'avatar-upload', component: _import('components-demo/avatarUpload'), name: 'avatarUpload-demo', meta: { title: 'avatarUpload' }},
      { path: 'dropzone', component: _import('components-demo/dropzone'), name: 'dropzone-demo', meta: { title: 'dropzone' }},
      { path: 'sticky', component: _import('components-demo/sticky'), name: 'sticky-demo', meta: { title: 'sticky' }},
      { path: 'count-to', component: _import('components-demo/countTo'), name: 'countTo-demo', meta: { title: 'countTo' }},
      { path: 'mixin', component: _import('components-demo/mixin'), name: 'componentMixin-demo', meta: { title: 'componentMixin' }},
      { path: 'back-to-top', component: _import('components-demo/backToTop'), name: 'backToTop-demo', meta: { title: 'backToTop' }}
    ]
  },

  {
    path: '/charts',
    component: Layout,
    redirect: 'noredirect',
    hidden: true,
    name: 'charts',
    meta: {
      title: 'charts',
      icon: 'chart'
    },
    children: [
      { path: 'keyboard', component: _import('charts/keyboard'), name: 'keyboardChart', meta: { title: 'keyboardChart', noCache: true }},
      { path: 'line', component: _import('charts/line'), name: 'lineChart', meta: { title: 'lineChart', noCache: true }},
      { path: 'mixchart', component: _import('charts/mixChart'), name: 'mixChart', meta: { title: 'mixChart', noCache: true }}
    ]
  },

  {
    path: '/form',
    component: Layout,
    redirect: 'noredirect',
    hidden: true,
    name: 'form',
    meta: {
      title: 'form',
      icon: 'form'
    },
    children: [
      { path: 'create-form', component: _import('form/create'), name: 'createForm', meta: { title: 'createForm', icon: 'table' }},
      { path: 'edit-form', component: _import('form/edit'), name: 'editForm', meta: { title: 'editForm', icon: 'table' }}
    ]
  },

  {
    path: '/error',
    component: Layout,
    redirect: 'noredirect',
    hidden: true,
    name: 'errorPages',
    meta: {
      title: 'errorPages',
      icon: '404'
    },
    children: [
      { path: '401', component: _import('errorPage/401'), name: 'page401', meta: { title: 'page401', noCache: true }},
      { path: '404', component: _import('errorPage/404'), name: 'page404', meta: { title: 'page404', noCache: true }}
    ]
  },

  {
    path: '/error-log',
    component: Layout,
    redirect: 'noredirect',
    hidden: true,
    children: [{ path: 'log', component: _import('errorLog/index'), name: 'errorLog', meta: { title: 'errorLog', icon: 'bug' }}]
  },

  {
    path: '/excel',
    component: Layout,
    redirect: '/excel/export-excel',
    hidden: true,
    name: 'excel',
    meta: {
      title: 'excel',
      icon: 'excel'
    },
    children: [
      { path: 'export-excel', component: _import('excel/exportExcel'), name: 'exportExcel', meta: { title: 'exportExcel' }},
      { path: 'export-selected-excel', component: _import('excel/selectExcel'), name: 'selectExcel', meta: { title: 'selectExcel' }},
      { path: 'upload-excel', component: _import('excel/uploadExcel'), name: 'uploadExcel', meta: { title: 'uploadExcel' }}
    ]
  },

  {
    path: '/zip',
    component: Layout,
    redirect: '/zip/download',
    alwaysShow: true,
    meta: { title: 'zip', icon: 'zip' },
    hidden: true,
    children: [{ path: 'download', component: _import('zip/index'), name: 'exportZip', meta: { title: 'exportZip' }}]
  },
  {
    path: '/theme',
    component: Layout,
    redirect: 'noredirect',
    children: [{ path: 'index', component: _import('theme/index'), name: 'theme', meta: { title: '主题', icon: 'theme' }}]
  },

  { path: '*', redirect: '/404', hidden: true }
]
