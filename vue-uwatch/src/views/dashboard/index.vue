<template>
  <div>
    <titleBar @showgrafana="getTitleChild" @showdata="getMsgFromChild"></titleBar>
    <!--background-image: linear-gradient(180deg,#1d3e74,#376290 370px,#489292 630px,#6fc6c7 1030px)-->
    <div style="width:100%;height:100%;	position: absolute;top: 0px;left: 0;right: 0;bottom: 0;background-position: 0 0;background-image: linear-gradient(180deg,#1d3e74,#376290 370px,#489292 630px,#6fc6c7 1030px);">
      <section class="white-bg " style="margin-left: -15px;">
            <div style="width:100%;height: 100%;" v-show="this.search">
              <div class="demo-input-size" style="width: 638px;margin: 0 auto;padding-top:18%;padding-bottom:45px;">
                <el-input placeholder="请输入内容" v-model="SearchGrafanaKeyWord.Query" class="input-with-select" size="88px;" @keyup.enter.native="remoteMethod">
                  <el-button slot="append" icon="el-icon-search" @click="remoteMethod" style="background-color: #3077f1;color: #f8f8f8;vertical-align: middle;display: table-cell;position: relative;border: 1px solid #3077f1;"></el-button>
                </el-input>
              </div>
              <div id="clock">
                <p class="date">{{ date }}</p>
                <p class="time">{{ time }}</p>
              </div>
            </div>
      </section>
      <!--<el-footer style="height:50px;background-color: rgb(84, 92, 100);">-->
      <!--<h5 style="font-size: 14px;margin: 0 auto;width: 201px;padding-top:20px;">技术保障组@平台工具线</h5>-->
      <!--</el-footer>-->
    </div>
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
  #box{
    width: 800px;
    margin: 0 auto;
    height: 32px;
    line-height: 30px;
    overflow: hidden;
    transition: all 0.5s;
  }
  .anim{
    transition: all 0.5s;
  }
  #con1 li{
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
    -webkit-transition: border-color .2s cubic-bezier(.645,.045,.355,1);
    transition: border-color .2s cubic-bezier(.645,.045,.355,1);
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
  #clock {
    font-family: 'Share Tech Mono', monospace;
    color: #ffffff;
    text-align: center;
    position: relative;
    width: 400px;
    /*left: 50%;*/
    /*top: 50%;*/
    margin: 0 auto;
    /*-webkit-transform: translate(-50%, -50%);*/
    /*transform: translate(-50%, -50%);*/
    color: #daf6ff;
    text-shadow: 0 0 20px #0aafe6, 0 0 20px rgba(10, 175, 230, 0);
  }
  #clock .time {
    letter-spacing: 0.05em;
    font-size: 80px;
    padding: 5px 0;
  }
  #clock .date {
    letter-spacing: 0.1em;
    font-size: 24px;
  }
  #clock .text {
    letter-spacing: 0.1em;
    font-size: 12px;
    padding: 20px 0 0;
  }
</style>
<script type="text/javascript">
  import titleBar from './titleBar'
  import leftNav from './leftNav'
  import { getSearchKeyWord, getGraphData, getGraphCounter, getFalconCounter } from '@/api/search_keyword'
  import '@/assets/dashboard/style.css'
  import echarts from 'echarts'
  import { parseTime } from '@/utils'
  export default {
    data() {
      return {
        load_data: false,
        SearchGrafanaKeyWord: {
          Query: ''
        },
        SearchGrafanaKeyTimes: {
          Times: 1
        },
        SearchGrafanaKeyWordQuery: {
          page: 1,
          limit: 100,
          sort: '-Id',
          query: '',
          times: '',
          status: 1
        },
        SelectCounterQuery: {
          query: ''
        },
        SelectGrafanaKeyWord: {
          Query: ''
        },
        showGrafana: false,
        SearchGrafana: false,
        IndexSearchhref: 'http://gochart.juanpi.org/show?domain=total',
        search: true,
        href: '',
        Searchhref: '',
        keyword: '',
        chart_idle: null,
        chart_user: null,
        chart_df: null,
        chart_mem: null,
        idle: 'idle',
        mem: 'mem',
        df: 'df',
        user: 'user',
        idleShow: true,
        memShow: true,
        userShow: true,
        dfShow: true,
        Host: 'api.huiguo.net^会过实时连接数',
        refresh: 0,
        count: 0,
        t: null,
        showLine: '',
        refreshStatus: 0,
        value9: [],
        options: [],
        CounterList: [],
        options4: [],
        list: [],
        loading: false,
        states: [],
        version: false,
        animate: false,
        query: null,
        time: '',
        date: ''
      }
    },
    components: {
      titleBar,
      leftNav
    },
    mounted() {
      setInterval(this.updateTime, 1000)
      this.updateTime()
    },
    methods: {
      updateTime() {
        var cd = new Date()
        const week = ['星期天', '星期一', '星期二', '星期三', '星期四', '星期五', '星期六']
        this.time = this.zeroPadding(cd.getHours(), 2) + ':' + this.zeroPadding(cd.getMinutes(), 2) + ':' + this.zeroPadding(cd.getSeconds(), 2)
        this.date = this.zeroPadding(cd.getFullYear(), 4) + '-' + this.zeroPadding(cd.getMonth() + 1, 2) + '-' + this.zeroPadding(cd.getDate(), 2) + ' ' + week[cd.getDay()]
      },
      zeroPadding(num, digit) {
        var zero = ''
        for (var i = 0; i < digit; i++) {
          zero += '0'
        }
        return (zero + num).slice(-digit)
      },
      remoteMethod(query) {
        query = this.SearchGrafanaKeyWord.Query
        if (query !== '') {
          console.log('states===>', this.version)
          if (this.version === false) {
            this.options = ''
            this.counterList_v1(query)
          } else {
            this.options = ''
            this.counterList_v2(query)
          }
          setTimeout(() => {
            this.list = this.states.map(item => {
              return { value: item, label: item }
            })
          }, 3000)
          this.list = this.states
          this.loading = true
          setTimeout(() => {
            this.loading = false
            this.options = this.list.filter(item => {
              return item.label.toLowerCase().indexOf(query.toLowerCase()) > -1
            })
          }, 1000)
        } else {
          this.$notify({
            title: '提示',
            message: '请输入要搜索的主机名或IP',
            type: 'warning',
            duration: 2000
          })
        }
      },
      getGraphCounterList() {
        getGraphCounter().then(response => {
          console.log('response-counterlist===>', response)
          if (response.data != null) {
            this.CounterList = response.data
          } else {
            console.log('err===>')
          }
        }).catch(err => {
          console.log(err)
        })
      },
      counterList_v1(query) {
        this.$emit('searchData', query)
        this.$router.push({ path: '/dashboard/search', query: { search: query }})
      },
      counterList_v2(query) {
        this.SelectGrafanaKeyWord.Query = query
        this.SelectCounterQuery.query = JSON.stringify(this.SelectGrafanaKeyWord)
        getFalconCounter(this.SelectCounterQuery).then(response => {
          this.states = response.data
        }).catch(err => {
          console.log(err)
        })
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
        }
      },
      // 通过子组件监听获取关闭dialog状态
      getMsgFromChild(v) {
        console.log('vvvv===>', v)
        this.SearchGrafanaKeyWord.Query = v
        this.initGraph()
      },
      getTitleChild(v) {
        this.showGrafana = false
        this.search = v
      },
      submitForm() {
        this.SearchGrafanaKeyWordQuery.query = JSON.stringify(this.SearchGrafanaKeyWord)
        this.getList()
      },
      getList() {
        getSearchKeyWord(this.SearchGrafanaKeyWordQuery).then(response => {
          if (response.data.list != null) {
            this.$notify({
              title: '成功',
              message: '检索成功',
              type: 'success',
              duration: 2000
            })
            this.showGrafana = true
            this.search = false
            this.href = response.data.list[0]['Url']
          } else {
            this.$notify({
              title: '失败',
              message: '关键词不存在',
              type: 'warning',
              duration: 2000
            })
          }
          this.Searchhref = response.data.list
        }).catch(err => {
          console.log(err)
        })
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
        } else {
          this.chart_mem = echarts.init(document.getElementById(this.mem))
          return this.chart_mem
        }
      },
      getGraphIdValue(id) {
        if (id === 0) {
          return 'idle'
        } else if (id === 1) {
          return 'user'
        } else if (id === 2) {
          return 'df'
        } else {
          return 'mem'
        }
      },
      initGraph() {
        this.idleShow = false
        this.memShow = false
        this.userShow = false
        this.dfShow = false
        this.SearchGrafanaKeyWordQuery.query = JSON.stringify(this.SearchGrafanaKeyWord)
        this.SearchGrafanaKeyWordQuery.times = this.SearchGrafanaKeyTimes.Times
        this.SearchGrafanaKeyWordQuery.status = 1
        console.log('iniiniini===>', this.SearchGrafanaKeyWord)
        getGraphData(this.SearchGrafanaKeyWordQuery, this.version).then(response => {
          console.log('return===>', (response.data !== null))
          if (response.data !== null) {
            for (var v = 0; v < response.data.length; v++) {
              var showLine = this.getGraphIdValue(v) + 'Show'
              if (showLine === 'idleShow') {
                this.idleShow = true
              } else if (showLine === 'memShow') {
                this.memShow = true
              } else if (showLine === 'userShow') {
                this.userShow = true
              } else {
                this.dfShow = true
              }

              const ylist = []
              const xlist = []

              for (var q = 0; q < response.data[v].values.length; q++) {
                xlist.push(parseTime(response.data[v].values[q].timestamp * 1000))
                ylist.push(response.data[v].values[q].value / parseInt(response.data[v].dividend))
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
                backgroundColor: '#21202D',
                title: {
                  text: response.data[v].title,
                  subtext: response.data[v].counter,
                  textStyle: {
                    color: '#C0C0C0'
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
                          color: '#ffe',
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
            if (this.version === false) {
              this.$notify({
                title: '失败',
                message: '检索失败，0.1版本该IP或域名数据暂时未配置',
                type: 'warning',
                duration: 2000
              })
            } else {
              this.$notify({
                title: '提示',
                message: '未搜索到相关信息！',
                type: 'warning',
                duration: 2000
              })
            }
          }
        })
      }
    }
  }
</script>
