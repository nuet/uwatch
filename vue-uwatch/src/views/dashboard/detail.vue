<template>
  <div>
    <titleBar  @showdata="getMsgFromChild" style="position: fixed;width: 100%;"></titleBar>
    <div style="width:100%;height:100%;	position: absolute;top: 60px;left: 0;right: 0;bottom: 0;">
      <section class="detail_bg " style="background-color: #0a76a4">
        <div>
          <div class="demo-input-size" style="width: 500px; margin: 0 auto;padding-top: 20px;">
          <span style="color: #409EFF;padding-left:5px;">是否开启自动刷新</span>
          <el-switch
          v-model="refresh"
          :change="beforeDestroy()"
          active-color="#0a76a4"
          inactive-color="#C0C0C9">
          </el-switch>
          <span style="color:#000;padding-left:10px;">|</span>
          <span style="color: #409EFF; margin-left:12px;">获取前</span>
          <el-input-number v-model="SearchGrafanaKeyTimes.Times" controls-position="right" @change="initGraph" :min="1" :max="24" size="mini"></el-input-number>
          <span style="color:#409EFF">小时</span>
          </div>
        </div>
        <hr>
        <div>
          <!--<ve-line :data="chartData"></ve-line>-->
          <div class="col-md-12" style="margin-bottom: 30px;margin-top:10px;">
          <div :id="this.container1"  class="ccontainer"></div>
          <div :id="this.container2"  class="ccontainer"></div>
          <div :id="this.container3"  class="ccontainer"></div>
          </div>
          <div class="col-md-12" style="margin-bottom: 30px;margin-top:10px;">
          <div :id="this.idle" class="chart detail"></div>
          <div :id="this.mem"  class="chart detail" ></div>
          <div :id="this.df"  class="chart detail" ></div>
          <div :id="this.user"  class="chart detail"></div>
          <div :id="this.detail1"  class="chart detail"></div>
          <div :id="this.detail2"  class="chart detail" ></div>
          <div :id="this.detail3"  class="chart detail" ></div>
          <div :id="this.detail4"  class="chart detail"></div>
          <div :id="this.detail5"  class="chart detail" ></div>
          <div :id="this.detail6"  class="chart detail" ></div>
          <div :id="this.detail7"  class="chart detail" ></div>
          <div :id="this.detail8"  class="chart detail" ></div>
          </div>
        </div>
      </section>
    </div>
  </div>
</template>

<script>
  import Chart from '@/components/Charts/keyboard'
  import titleBar from './titleBar'
  import '@/assets/dashboard/css/detail.css'
  import { getFalconCounter, getGraphData } from '@/api/search_keyword'
  import echarts from 'echarts'
  import { parseTime } from '@/utils'
  export default {
    name: 'keyboardChart',
    components: { Chart, titleBar },
    data() {
      return {
        SelectGrafanaKeyWord: {
          Query: ''
        },
        SelectCounterQuery: {
          query: ''
        },
        SearchGrafanaKeyTimes: {
          Times: 1
        },
        SearchGrafanaKeyWord: {
          Query: ''
        },
        SearchGrafanaKeyWordQuery: {
          page: 1,
          limit: 100,
          sort: '-Id',
          query: '',
          times: '',
          status: 1
        },
        version: true,
        chart_idle: null,
        chart_user: null,
        chart_df: null,
        chart_detail1: null,
        chart_detail2: null,
        chart_detail3: null,
        chart_detail4: null,
        chart_detail5: null,
        chart_detail6: null,
        chart_detail7: null,
        chart_detail8: null,
        chart_container1: null,
        chart_container2: null,
        chart_container3: null,
        idle: 'idle',
        mem: 'mem',
        df: 'df',
        user: 'user',
        detail1: 'detail1',
        detail2: 'detail2',
        detail3: 'detail3',
        detail4: 'detail4',
        detail5: 'detail5',
        detail6: 'detail6',
        detail7: 'detail7',
        detail8: 'detail8',
        container1: 'container1',
        container2: 'container2',
        container3: 'container3',
        idleShow: false,
        memShow: false,
        userShow: false,
        dfShow: false,
        detail1Show: false,
        detail2Show: false,
        detail3Show: false,
        detail4Show: false,
        detail5Show: false,
        detail6Show: false,
        detail7Show: false,
        detail8Show: false,
        container1Show: true,
        container2Show: true,
        container3Show: true,
        refresh: 0,
        refreshStatus: 0
      }
    },
    created() {
      console.log('endpoint===>', this.$route.query.endpoint)
      this.initGraph()
      this.topinitGraph(this.$route.query.endpoint)
    },
    methods: {
      topinitGraph(query) {
        this.SelectGrafanaKeyWord.Query = query
        this.SelectCounterQuery.query = JSON.stringify(this.SelectGrafanaKeyWord)
        getFalconCounter(this.SelectCounterQuery).then(response => {
          console.log('getFalconEndpoint+Counte===>', response.data)
          this.SearchGrafanaKeyWord.Query = response.data[0].counter
          this.SearchGrafanaKeyWordQuery.query = JSON.stringify(this.SearchGrafanaKeyWord)
          this.SearchGrafanaKeyWordQuery.times = this.SearchGrafanaKeyTimes.Times
          this.SearchGrafanaKeyWordQuery.status = 1
          getGraphData(this.SearchGrafanaKeyWordQuery, this.version).then(response => {
            console.log('DETAIL----return===>', response.data)
            if (response.data !== null) {
              for (var v = 0; v < response.data.length; v++) {
                if (response.data[v].counter.length === 0) {
                  continue
                } else {
                  var showLine = this.getTopGraphIdValue(v) + 'Show'
                  if (showLine === 'container1Show') {
                    this.container1Show = true
                    document.getElementById(this.container1).style.display = 'block'
                  } else if (showLine === 'container2Show') {
                    this.container2Show = true
                    document.getElementById(this.container2).style.display = 'block'
                  } else if (showLine === 'container3Show') {
                    this.container3Show = true
                    document.getElementById(this.container3).style.display = 'block'
                  } else {
                    this.container3Show = true
                    document.getElementById(this.container3).style.display = 'block'
                  }
                  const ylist = []
                  const xlist = []
                  if (response.data[v].title === '内存使用率') {
                    for (var q = 0; q < response.data[v].values.length; q++) {
                      xlist.push(parseTime(response.data[v].values[q].timestamp * 1000))
                      ylist.push((100 - response.data[v].values[q].value / parseInt(response.data[v].dividend)).toFixed(2))
                    }
                  } else {
                    for (var p = 0; p < response.data[v].values.length; p++) {
                      xlist.push(parseTime(response.data[v].values[p].timestamp * 1000))
                      ylist.push((response.data[v].values[p].value / (response.data[v].dividend)))
                    }
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

                  var Ids = this.getTopGraphId(v)
                  Ids.setOption({
                    toolbox: {
                      feature: {
                      }
                    },
                    series: [
                      {
                        name: response.data[v].title,
                        type: 'gauge',
                        detail: { formatter: '{value}%' },
                        data: [{ value: ylist[response.data[v].values.length - 1], name: response.data[v].title }]
                      }
                    ]
                  })
                }
              }
            } else {
              console.log('未获取到load.1min')
            }
          })
        }).catch(err => {
          console.log(err)
        })
      },
      initGraph() {
        this.SelectGrafanaKeyWord.Query = this.$route.query.endpoint
        this.SelectCounterQuery.query = JSON.stringify(this.SelectGrafanaKeyWord)
        getFalconCounter(this.SelectCounterQuery).then(response => {
          console.log('getFalconEndpoint+Counte===>', response.data)
          this.SearchGrafanaKeyWord.Query = response.data[0].value
          this.SearchGrafanaKeyWordQuery.query = JSON.stringify(this.SearchGrafanaKeyWord)
          this.SearchGrafanaKeyWordQuery.times = this.SearchGrafanaKeyTimes.Times
          this.SearchGrafanaKeyWordQuery.status = 1
          getGraphData(this.SearchGrafanaKeyWordQuery, this.version).then(response => {
            console.log('return===>', (response.data !== null))
            if (response.data !== null) {
              for (var v = 0; v < response.data.length; v++) {
                if (response.data[v].values.length === 0 || response.data[v].title === '1分钟负载' || response.data[v].title === '5分钟负载' || response.data[v].title === '15分钟负载') {
                  continue
                } else {
                  var showLine = this.getGraphIdValue(v) + 'Show'
                  if (showLine === 'idleShow') {
                    this.idleShow = true
                    document.getElementById(this.idle).style.display = 'block'
                  } else if (showLine === 'memShow') {
                    this.memShow = true
                    document.getElementById(this.mem).style.display = 'block'
                  } else if (showLine === 'userShow') {
                    this.userShow = true
                    document.getElementById(this.user).style.display = 'block'
                  } else if (showLine === 'detail1Show') {
                    this.detail1Show = true
                    document.getElementById(this.detail1).style.display = 'block'
                  } else if (showLine === 'detail2Show') {
                    this.detail2Show = true
                    document.getElementById(this.detail2).style.display = 'block'
                  } else if (showLine === 'detail3Show') {
                    this.detail3Show = true
                    document.getElementById(this.detail3).style.display = 'block'
                  } else if (showLine === 'detail4Show') {
                    this.detail4Show = true
                    document.getElementById(this.detail4).style.display = 'block'
                  } else if (showLine === 'detail5Show') {
                    this.detail5Show = true
                    document.getElementById(this.detail5).style.display = 'block'
                  } else if (showLine === 'detail6Show') {
                    this.detail6Show = true
                    document.getElementById(this.detail6).style.display = 'block'
                  } else if (showLine === 'detail7Show') {
                    this.detail7Show = true
                    document.getElementById(this.detail7).style.display = 'block'
                  } else if (showLine === 'detail8Show') {
                    this.detail8Show = true
                    document.getElementById(this.detail8).style.display = 'block'
                  } else {
                    this.dfShow = true
                    document.getElementById(this.df).style.display = 'block'
                  }

                  const ylist = []
                  const xlist = []

                  if (response.data[v].title === '内存使用率') {
                    for (var q = 0; q < response.data[v].values.length; q++) {
                      xlist.push(parseTime(response.data[v].values[q].timestamp * 1000))
                      ylist.push((100 - response.data[v].values[q].value / parseInt(response.data[v].dividend)).toFixed(2))
                    }
                  } else {
                    for (var p = 0; p < response.data[v].values.length; p++) {
                      xlist.push(parseTime(response.data[v].values[p].timestamp * 1000))
                      ylist.push(response.data[v].values[p].value / parseInt(response.data[v].dividend))
                    }
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

                  var Ids = this.getGraphId(v)
                  console.log('showLine==>', this.idle)
                  console.log('ylistylistylistylist==>', ylist)
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
              }
            } else {
              if (this.version === false) {
                this.$notify({
                  title: '失败',
                  message: '检索失败，数据暂时未配置',
                  type: 'warning',
                  duration: 2000
                })
                this.$router.push({ path: '/dashboard/search', query: { search: this.$route.query.search }})
              } else {
                this.$notify({
                  title: '提示',
                  message: '检索失败，数据暂时未配置',
                  type: 'warning',
                  duration: 2000
                })
                this.$router.push({ path: '/dashboard/search', query: { search: this.$route.query.search }})
              }
            }
          })
        }).catch(err => {
          console.log(err)
        })
      },
      // 通过子组件监听获取关闭dialog状态
      getMsgFromChild(v) {
        this.$router.push({ path: '/dashboard' })
      },
      getGraphId(id) {
        if (id === 0) {
          this.chart_idle = echarts.init(document.getElementById(this.idle))
          return this.chart_idle
        } else if (id === 1) {
          this.chart_user = echarts.init(document.getElementById(this.user))
          return this.chart_user
        } else if (id === 2) {
          this.chart_df = echarts.init(document.getElementById(this.df))
          return this.chart_df
        } else if (id === 3) {
          this.chart_detail1 = echarts.init(document.getElementById(this.detail1))
          return this.chart_detail1
        } else if (id === 4) {
          this.chart_detail2 = echarts.init(document.getElementById(this.detail2))
          return this.chart_detail2
        } else if (id === 5) {
          this.chart_detail3 = echarts.init(document.getElementById(this.detail3))
          return this.chart_detail3
        } else if (id === 6) {
          this.chart_detail4 = echarts.init(document.getElementById(this.detail4))
          return this.chart_detail4
        } else if (id === 7) {
          this.chart_detail5 = echarts.init(document.getElementById(this.detail5))
          return this.chart_detail5
        } else if (id === 8) {
          this.chart_detail6 = echarts.init(document.getElementById(this.detail6))
          return this.chart_detail6
        } else if (id === 9) {
          this.chart_detail7 = echarts.init(document.getElementById(this.detail7))
          return this.chart_detail7
        } else if (id === 10) {
          this.chart_detail8 = echarts.init(document.getElementById(this.detail8))
          return this.chart_detail8
        } else {
          this.chart_mem = echarts.init(document.getElementById(this.mem))
          return this.chart_mem
        }
      },
      getTopGraphId(id) {
        if (id === 0) {
          this.chart_container1 = echarts.init(document.getElementById(this.container1))
          return this.chart_container1
        } else if (id === 1) {
          this.chart_container2 = echarts.init(document.getElementById(this.container2))
          return this.chart_container2
        } else if (id === 2) {
          this.chart_container3 = echarts.init(document.getElementById(this.container3))
          return this.chart_container3
        } else {
          this.chart_container3 = echarts.init(document.getElementById(this.container3))
          return this.chart_container3
        }
      },
      beforeDestroy: function() {
        console.log(this.refresh)
        if (this.refresh) {
          this.count = 1
          this.t = setInterval(this.timer, 6000)
        } else {
          this.count = 0
          this.refreshStatus = 0
          clearInterval(this.t)
        }
      },
      timer: function() {
        if (this.count > 0) {
          this.count++
          this.refreshStatus = 1
          this.initGraph(this.refreshStatus)
          this.topinitGraph(this.$route.query.endpoint)
        }
      },
      getGraphIdValue(id) {
        if (id === 0) {
          return 'idle'
        } else if (id === 1) {
          return 'user'
        } else if (id === 2) {
          return 'df'
        } else if (id === 3) {
          return 'detail1'
        } else if (id === 4) {
          return 'detail2'
        } else if (id === 5) {
          return 'detail3'
        } else if (id === 6) {
          return 'detail4'
        } else if (id === 7) {
          return 'detail5'
        } else if (id === 8) {
          return 'detail6'
        } else if (id === 9) {
          return 'detail7'
        } else if (id === 10) {
          return 'detail8'
        } else {
          return 'mem'
        }
      },
      getTopGraphIdValue(id) {
        if (id === 0) {
          return 'container1'
        } else if (id === 1) {
          return 'container2'
        } else if (id === 2) {
          return 'container3'
        } else {
          return 'container3'
        }
      }
    }
  }
</script>

