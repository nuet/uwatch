<template>
  <div>
    <titleBar @showgrafana="getTitleChild" @showdata="getMsgFromChild" style="position: fixed;width: 100%;"></titleBar>
    <div style="width:100%;height:100%;	position: absolute;top: 60px;left: 0;right: 0;bottom: 0;">
      <section class="search-white-bg " style="">
        <div class="demo-input-size" style="width: 638px;padding-top:20px;padding-left:100px;">
          <span style="font-size: 18px;color:#909399">搜索结果</span>
        </div>
        <hr/>
        <div style="width:100%;height: 100%;" v-show="this.search">
          <div class="demo-input-size" style="width: 800px;padding-bottom:45px;padding-top:10px;padding-left:100px;">
            <el-input placeholder="请输入内容" v-model="SearchGrafanaKeyWord.Query" class="input-with-select" size="88px;"
                      style="border: 1px solid #3077f1;color: #606266;height: 50px;border-radius: 6px;font-size: 20px;"
                      @keyup.enter.native="remoteMethod">
              <el-button slot="append" icon="el-icon-search" style="
              background-color: #3077f1;
              color: rgb(241, 243, 246);
              vertical-align: middle;
              display: table-cell;
              position: relative;
              border-radius: 0px;
              border: 1px solid #3077f1;" @click="remoteMethod"></el-button>
            </el-input>
          </div>
          <div v-for="item in states">
            <div style="font-size:14px;padding-left:100px;padding-right: 100px;">
              <router-link style="" :to="{ path:'/dashboard/detail', query: { endpoint: item.title, search: SearchGrafanaKeyWord.Query }}">
                <span style="margin-top:-20px;color:#167be0; line-height: 30px;font-size: 16px;">主机节点/内网IP-<span
                  style="color:rgb(85, 183, 228)">{{item.title}}</span></span>
              </router-link>
              <div style="margin-top:8px; line-height: 20px;font-size: 16px;">
                主机名称:  <span style="color:rgba(5,6,3,0.55);"
                             v-if="item.BkHostName != ''">{{item.BkHostName}}</span><span v-else
                                                                                          style="color:rgba(5,6,3,0.55)">无</span>
                内网IP:  <span style="color:rgba(5,6,3,0.55)"
                             v-if="item.BkHostInnerip != ''">{{item.BkHostInnerip}}</span><span v-else
                                                                                                style="color:rgba(5,6,3,0.55)">无</span>
                主要维护人:  <span style="color:rgba(5,6,3,0.55)" v-if="item.Operator != ''">{{item.Operator}}</span><span
                v-else style="color:rgba(5,6,3,0.55)">无</span>
                备份维护人:  <span style="color:rgba(5,6,3,0.55)"
                              v-if="item.BkBakOperator != ''">{{item.BkBakOperator}}</span><span v-else
                                                                                                 style="color:rgba(5,6,3,0.55)">无</span>
                当前状态:  <span style="color:rgba(5,6,3,0.55)"
                             v-if="item.BkCurrentStatus != ''">{{item.BkCurrentStatus}}</span><span v-else
                                                                                                    style="color:rgba(5,6,3,0.55)">无</span>
                运行状态:  <span style="color:rgba(5,6,3,0.55)" v-if="item.BkStatus != ''">{{item.BkStatus}}</span><span
                v-else style="color:rgba(5,6,3,0.55)">无</span>
              </div>
              <div style="margin-top:8px; line-height: 20px;font-size: 16px;">
                厂商:  <span style="color:rgba(5,6,3,0.55)"
                           v-if="item.BkManufacturer != ''">{{item.BkManufacturer}}</span><span v-else
                                                                                                style="color:rgba(5,6,3,0.55)">无</span>
                型号:  <span style="color:rgba(5,6,3,0.55)"
                           v-if="item.BkProductName != ''">{{item.BkProductName}}</span><span v-else
                                                                                              style="color:rgba(5,6,3,0.55)">无</span>
                操作系统名称:  <span style="color:rgba(5,6,3,0.55)" v-if="item.BkOsName != ''">{{item.BkOsName}}</span><span
                v-else style="color:rgba(5,6,3,0.55)">无</span>
                业务拓扑:  <span style="color:rgba(5,6,3,0.55)" v-if="item.Module != ''">{{item.Module}}</span><span v-else
                                                                                                                 style="color:rgba(5,6,3,0.55)">无</span>
              </div>
              <div style="margin-top:8px;font-size: 16px;height:auto;">
                Counter:
                <span style="margin-top:8px; margin-right:8px;font-size: 16px;" v-for="ct in item.Counter">
                   <el-button  size="mini" v-if="ct.id <= 3" @click="showGraph(item.title+'^'+ct.name)" style="margin-bottom: 5px;" round>{{ct.name}}</el-button>
                </span>
                <v-selectpage :data="item.Counter" key-field="id"  class="form-control"
                                :tb-columns="showFields"
                                @values="singleValues"
                                title="Counter"
                                placeholder="查看更多Counter"
                                v-if="item.length >= 4"
                                style="width:850px;">
                </v-selectpage >
                <!--<el-popover-->
                  <!--v-if="item.length > 0"-->
                  <!--placement="right"-->
                  <!--width="845"-->
                  <!--trigger="click">-->
                  <!--<el-table-->
                    <!--height="600"-->
                    <!--border-->
                    <!--:data="item.Counter">-->
                    <!--<el-table-column  label="主机节点/内网IP" width="350">-->
                      <!--<template slot-scope="scope">-->
                        <!--<span>{{item.title}}</span>-->
                      <!--</template>-->
                    <!--</el-table-column>-->
                    <!--<el-table-column  property="counter" label="Counter" width="350"></el-table-column>-->
                    <!--<el-table-column-->
                      <!--label="操作"-->
                      <!--width="100">-->
                      <!--<template slot-scope="scope">-->
                        <!--<el-button  size="mini" @click="showGraph(item.title+'^'+scope.row.counter)"  round>查看</el-button>-->
                      <!--</template>-->
                    <!--</el-table-column>-->
                  <!--</el-table>-->
                  <!--<el-button slot="reference" size="mini"  round><i class="el-icon-circle-plus-outline"></i>所有Counter</el-button>-->
                <!--</el-popover>-->
              </div>
              <!--<el-tooltip placement="top-start" effect="dark" v-if="item.length > 2">-->
                <!--<el-button size="small" round><i class="el-icon-circle-plus-outline"></i>更多</el-button>-->
                <!--<div slot="content" style="margin-left:8px;">-->
                     <!--<span style="margin-top:8px; margin-right:8px;font-size: 16px;" v-for="ct in item.Counter">-->
                        <!--<el-button v-if="ct.id > 2" size="mini" @click="showGraph(item.title+'^'+ct.counter)" style="margin-bottom: 8px;" round>{{ct.counter}}</el-button>-->
                      <!--</span>-->
                <!--</div>-->
              <!--</el-tooltip>-->
            </div>
            <hr>
          </div>
          <div class="pagination-container" style="width:330px;margin: 0 auto;padding-bottom: 20px;">
            <el-pagination background @size-change="handleSizeChange" @current-change="handleCurrentChange"
                           :current-page="listQuery.page" :page-sizes="[10,20,30, 50]" :page-size="listQuery.limit"
                           layout="total, sizes, prev, pager, next, jumper" :total="total">
            </el-pagination>
          </div>
        </div>
      </section>
    </div>
    <el-dialog
      title="Falcon+"
      :visible.sync="dialogVisible"
      width="845px">
      <div id="graph" style="width:800px;height:350px;margin: 0 auto;"></div>
      <span slot="footer" class="dialog-footer">
        <!--<el-button type="primary" @click="dialogVisible = false">关 闭</el-button>-->
      </span>
    </el-dialog>
  </div>
</template>
<style>
  .panel-group {
    max-height: 770px;
    overflow: auto;
    margin-bottom: 0px;
  }

  .el-switch__label {
    -webkit-transition: .2s;
    transition: .2s;
    height: 20px;
    font-size: 14px;
    font-weight: 500;
    vertical-align: middle;
    color: rgb(84, 92, 100);
  }

  #box {
    width: 800px;
    margin: 0 auto;
    height: 32px;
    line-height: 30px;
    overflow: hidden;
    transition: all 0.5s;
  }

  .anim {
    transition: all 0.5s;
  }

  #con1 li {
    list-style: none;
    line-height: 32px;
    height: 30px;
  }

  .el-alert {
    width: 100%;
    line-height: 31px;
    height: 31px;
    padding: 0px 0px;
    margin: 0;
    -webkit-box-sizing: border-box;
    box-sizing: border-box;
    border-radius: 4px;
    position: relative;
    background-color: #0a76a4;
    overflow: hidden;
    opacity: 1;
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center;
    -webkit-transition: opacity .2s;
    transition: opacity .2s;
  }

  .el-input__inner {
    -webkit-appearance: none;
    background-color: #fff;
    background-image: none;
    border-radius: 6px;
    border-top-right-radius: 4px;
    border-bottom-right-radius: 4px;
    border: 1px solid #dcdfe6;
    border-top-color: rgb(220, 223, 230);
    border-right-color: rgb(220, 223, 230);
    border-bottom-color: rgb(220, 223, 230);
    border-left-color: rgb(220, 223, 230);
    -webkit-box-sizing: border-box;
    box-sizing: border-box;
    color: #606266;
    display: inline-block;
    font-size: 20px;
    font-size: inherit;
    height: 50px;
    line-height: 1;
    outline: 0;
    padding: 0 15px;
    -webkit-transition: border-color .2s cubic-bezier(.645, .045, .355, 1);
    transition: border-color .2s cubic-bezier(.645, .045, .355, 1);
    width: 100%;
  }

  .el-input-group--prepend .el-input__inner, .el-input-group__append {
    border-top-left-radius: 0;
    border-bottom-left-radius: 0;
  }

  .el-input-group__append, .el-input-group__prepend {
    background-color: #3077f1;
    color: #909399;
    vertical-align: middle;
    display: table-cell;
    position: relative;
    border: 1px solid #3077f1;
    border-left-width: 1px;
    border-left-style: solid;
    border-left-color: rgb(48, 119, 241);
    border-radius: 4px;
    border-top-left-radius: 0px;
    border-bottom-left-radius: 0px;
    padding: 0 20px;
    width: 1px;
    white-space: nowrap;
  }
</style>
<script type="text/javascript">
  import titleBar from './titleBar'
  import { Loading } from 'element-ui'
  import { getCmdbHost, getOneGraphData } from '@/api/search_keyword'
  import '@/assets/dashboard/style.css'
  import { parseTime } from '@/utils'
  import echarts from 'echarts'
  import Vue from 'vue'
  import vSelectPage from 'v-selectpage'
  Vue.use(vSelectPage)
  export default {
    data() {
      return {
        load_data: false,
        SearchGrafanaKeyWord: {
          Query: ''
        },
        SearchOneGrafanaKeyWord: {
          Query: ''
        },
        SelectCounterQuery: {
          query: ''
        },
        SelectGrafanaKeyWord: {
          Query: ''
        },
        showGrafana: false,
        SearchGrafana: false,
        search: true,
        href: '',
        Searchhref: '',
        keyword: '',
        refresh: 0,
        count: 0,
        t: null,
        showLine: '',
        refreshStatus: 0,
        value9: [],
        CounterList: [],
        options4: [],
        list: [],
        loading: false,
        states: [],
        version: false,
        graph: 'graph',
        animate: false,
        listQuery: {
          page: 1,
          limit: 10,
          Query: ''
        },
        total: null,
        query: null,
        SearchGrafanaKeyWordQuery: {
          page: 1,
          limit: 100,
          sort: '-Id',
          query: '',
          times: '',
          status: 1
        },
        dialogVisible: false,
        options: {
          lock: true,
          text: 'Loading',
          spinner: 'el-icon-loading',
          background: 'rgba(0, 0, 0, 0.7)'
        },
        showSelectCounter: false,
        showSelectTitle: '',
        showSelectName: '',
        showFields: [
          { title: 'Id', data: 'id' },
          { title: 'Counter', data: 'name' },
          { title: 'Endpoint', data: 'title' }
        ]
      }
    },
    components: {
      titleBar,
      vSelectPage
    },
    created() {
      this.SearchGrafanaKeyWord.Query = this.$route.query.search
      Loading.service(this.options)
      this.searchCmdbHost(this.$route.query.search)
    },
    methods: {
      showGraph(endpointCounter) {
        this.dialogVisible = true
        this.SearchOneGrafanaKeyWord.Query = endpointCounter
        this.SearchGrafanaKeyWordQuery.query = JSON.stringify(this.SearchOneGrafanaKeyWord)
        this.SearchGrafanaKeyWordQuery.times = 1
        this.SearchGrafanaKeyWordQuery.status = 1
        getOneGraphData(this.SearchGrafanaKeyWordQuery, this.version).then(response => {
          if (response.data !== null) {
            this.loadingInstance = Loading.service(this.options)
            this.$nextTick(() => {
              this.loadingInstance.close()
            })
            for (var v = 0; v < response.data.length; v++) {
              const ylist = []
              const xlist = []
              for (var q = 0; q < response.data[v].values.length; q++) {
                xlist.push(parseTime(response.data[v].values[q].timestamp * 1000))
                ylist.push(response.data[v].values[q].value)
              }
              var min = ylist[0]
              for (var ii = 1; ii < response.data[v].values.length; ii++) {
                if (ylist[ii] < min) {
                  min = ylist[ii]
                }
              }
              var max = ylist[0]
              for (var jj = 1; jj < response.data[v].values.length; jj++) {
                if (ylist[jj] > max) {
                  max = ylist[jj]
                }
              }
              min = Math.floor(min)
              max = Math.ceil(max)
              console.log('getOneGraphData', this.graph)
              var Ids = echarts.init(document.getElementById('graph'))

              Ids.setOption({
                toolbox: {
                  feature: {
                    restore: {},
                    saveAsImage: {}
                  }
                },
                backgroundColor: '#fbfbfb',
                title: {
                  text: response.data[v].title,
                  subtext: response.data[v].counter,
                  textStyle: {
                    color: '#C0C0C9'
                  }
                },
                tooltip: {
                  trigger: 'axis',
                  axisPointer: {
                    type: 'cross'
                  }
                },
                xAxis: {
                  type: 'category',
                  axisLine: {
                    lineStyle: {
                      color: '#409EFF'
                    }
                  },
                  boundaryGap: false,
                  data: xlist
                },
                yAxis: {
                  min: min,
                  max: max,
                  type: 'value',
                  axisLine: {
                    lineStyle: {
                      color: '#409EFF'
                    }
                  },
                  axisLabel: {
                    formatter: '{value}' + response.data[v].axislabel
                  },
                  axisPointer: {
                    snap: true
                  }
                },
                visualMap: {
                  show: false,
                  dimension: 0
                },
                series: [
                  {
                    name: response.data[v].counter,
                    type: 'line',
                    smooth: true,
                    data: ylist,
                    areaStyle: {
                      normal: {
                        color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [{
                          offset: 0,
                          color: '#8ec6ad'
                        }, {
                          offset: 1,
                          color: '#ffe'
                        }])
                      }
                    },
                    markLine: {
                      symbol: ['none', 'none'],
                      label: {
                        normal: { show: false }
                      },
                      lineStyle: {
                        normal: {
                          color: 'rgb(29, 62, 116)',
                          width: 2
                        }
                      },
                      data: [{
                        yAxis: response.data[v].avg
                      }]
                    }
                  }
                ]
              })
            }
          } else {
            this.loadingInstance = Loading.service(this.options)
            this.$nextTick(() => {
              this.loadingInstance.close()
            })
            if (this.version === false) {
              this.$notify({
                title: '失败',
                message: '检索失败，数据暂时未配置',
                type: 'warning',
                duration: 2000
              })
              this.$router.push({ path: '/dashboard/search' })
            } else {
              this.$notify({
                title: '提示',
                message: '检索失败，数据暂时未配置',
                type: 'warning',
                duration: 2000
              })
              this.$router.push({ path: '/dashboard/search' })
            }
          }
        })
      },
      remoteMethod(query) {
        query = this.SearchGrafanaKeyWord.Query
        Loading.service(this.options)
        if (query !== '') {
          this.searchCmdbHost(query)
        } else {
          this.$notify({
            title: '提示',
            message: '请输入要搜索的主机名或IP',
            type: 'warning',
            duration: 2000
          })
          this.loadingInstance = Loading.service(this.options)
          this.$nextTick(() => {
            this.loadingInstance.close()
          })
        }
      },
      handleClose(done) {
        this.$confirm('确认关闭？').then(_ => {
          done()
        }).catch(_ => {})
      },
      searchCmdbHost(query) {
        this.SelectGrafanaKeyWord.Query = query
        this.listQuery.query = JSON.stringify(this.SelectGrafanaKeyWord)
        console.log('listQuery===>', this.listQuery)
        getCmdbHost(this.listQuery).then(response => {
          this.states = response.data.list
          this.total = response.data.total
          this.loadingInstance = Loading.service(this.options)
          this.$nextTick(() => {
            this.loadingInstance.close()
          })
          console.log('getCmdbHost==>', response.data)
          if (response.data.total === 0) {
            this.$notify({
              title: '提示',
              message: '未搜索到相关信息！',
              type: 'warning',
              duration: 2000
            })
          } else {
            this.CounterList = response.data.list.Counter
          }
        }).catch(err => {
          console.log(err)
        })
      },
      handleSizeChange(val) {
        this.listQuery.limit = val
        this.searchCmdbHost(this.SelectGrafanaKeyWord.Query)
      },
      handleCurrentChange(val) {
        this.listQuery.page = val
        this.searchCmdbHost(this.SelectGrafanaKeyWord.Query)
      },
      // 通过子组件监听获取关闭dialog状态
      getMsgFromChild(v) {
        this.$router.push({ path: '/dashboard' })
      },
      searchData(v) {
        this.searchCmdbHost(v)
      },
      getTitleChild(v) {
        this.showGrafana = false
        this.search = v
      },
      singleValues(v) {
        console.log('vvvv=>', v[0]['title'], v[0]['name'])
        this.showGraph(v[0]['title'] + '^' + v[0]['name'])
      }
    }
  }
</script>
