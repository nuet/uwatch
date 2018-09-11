/**
 * Created by Administrator on 2018-05-29.
 */
import '../css/reset.css'
import '../icon/iconfont.css'
import '../css/common.css'
import '../css/course-view.css'
import echarts from 'echarts'
import axios from 'axios'
import { fetchAlarmData, fetchHgGraphData } from '@/api/monitor'
export default {
  name: 'monitor',
  data() {
    return {
      interval: null,
      myChart1: null,
      myChart2: null,
      myChart3: null,
      myChart4: null,
      myChart5: null,
      myChart6: null,
      myChart7: null,
      myChart8: null,
      myChart9: null,
      myChart10: null,
      Timer5: null,
      Timer8: null,
      Timer9: null,
      option5Index: -1,
      option8Index: -1,
      option9Index: -1,
      grandTimes: null,
      grandNum: 0,
      grandData: [],
      hgGraphTimes: null,
      hgButtonTimes: null,
      hgGraphNum: 0,
      hgGraphData: {},
      alarmTimes: null,
      failTimes: null,
      showFullscreenButton: true
    }
  },
  mounted() {
    this.initChart()
  },
  created() {
    this.interval = setInterval(() => {
      this.getNowFormatDate()
    }, 1000)
    this.grandTimes = setInterval(() => {
      this.initGrand()
    }, 10000)
    this.hgButtonTimes = setInterval(() => {
      this.hgButtonTime()
    }, 22000)
    this.alarmTimes = setInterval(() => {
      this.initAlarmData()
    }, 60000)
    this.failTimes = setInterval(() => {
      this.initFailData()
    }, 3600000)
    this.hgGraphTimes = setInterval(() => {
      this.getHgGraphData()
    }, 300000)
  },
  beforeDestroy() {
    this.myChart1.dispose()
    this.myChart2.dispose()
    this.myChart3.dispose()
    this.myChart4.dispose()
    this.myChart5.dispose()
    this.myChart6.dispose()
    this.myChart8.dispose()
    this.myChart9.dispose()
    this.myChart10.dispose()
    clearInterval(this.interval)
    clearInterval(this.grandTimes)
    clearInterval(this.hgButtonTimes)
    clearInterval(this.alarmTimes)
    clearInterval(this.failTimes)
    clearInterval(this.hgGraphTimes)
  },
  methods: {
    initChart() {
      this.myChart1 = echarts.init(document.getElementById('main1'));
      this.myChart2 = echarts.init(document.getElementById('main2'));
      this.myChart3 = echarts.init(document.getElementById('main3'));
      this.myChart4= echarts.init(document.getElementById('main4'));
      this.myChart5 = echarts.init(document.getElementById('main5'));
      this.myChart6= echarts.init(document.getElementById('main6'));
      //this.myChart7= echarts.init(document.getElementById('main7'));
      this.myChart8 = echarts.init(document.getElementById('main8'));
      this.myChart9= echarts.init(document.getElementById('main9'));
      this.myChart10 = echarts.init(document.getElementById('main10'));
      this.initWebUsage()
      this.initLivingData()
      this.initHGData()
      this.initFailData()
      this.initGrand()
      this.initResourceData()
      this.initAlarmData()
    },
    initWebUsage() {
      var originalData=[];
      originalData = [{
        value: 0,
        name: '失败率'
      }, {
        value: 100,
        name: '可用率'
      }];
      var colorList = ['#112958','#14ad5d'];
      // 总和
      echarts.util.each(originalData, function(item, index) {
        item.itemStyle = {
          normal: {
            color: colorList[index]
          }
        };
      });
      var option1={
        tooltip: {
          trigger: 'item',
          formatter: "<p class='center'>{b}</p><p class='center'>{c}</p>",
          transitionDuration:0,
          backgroundColor : 'rgba(83,93,105,0.8)',
          borderColor : '#535b69',
          borderRadius : 8,
          borderWidth: 2,
          padding: [5,10],
          position: ['50%', '50%']
        },
        graphic:{
          type: 'text',
          left: 'center', // 相对父元素居中
          top: 'middle',  // 相对父元素居中
          style: {
            fill: '#1d70f2',
            text: [
              100
            ],
            font: '40px myFirstFont'
          }
        },
        series: [
          {
            type:'pie',
            radius: ['48%', '60%'],
            hoverAnimation: false,
            avoidLabelOverlap: true,
            label: {
              normal: {
                show: false,
                position: 'center'
              }
            },
            labelLine: {
              normal: {
                show: false
              }
            },
            data:originalData
          }
        ],
        color:['#112958','#14ad5d']
      };
      this.myChart1.setOption(option1);
      this.myChart2.setOption(option1);
    },
    initLivingData() {
      var Tpl='';
      Tpl='<li></li>'+
        '<li></li>'+
        '<li></li>'+
        '<li></li>';
      $('.living-status .status-data').html(Tpl);
      var placeHolderStyle = {
        normal: {
          borderWidth: 5,
          shadowBlur: 40,
          borderColor: "#132235",
          shadowColor: 'rgba(0, 0, 0, 0)', //边框阴影
          color: "#132235"
        }
      };
      var option3 = {
        title:{
          show:true,
          text:'成单量 '+'单/秒',
          right:2,
          top:40,
          textStyle:{
            fontSize:16,
            color:'#1abffa'
          }
        },
        color: ['#1abffa','#11a754'],
        tooltip: {
          show:false,
          trigger: 'item'
        },
        legend: {
          type: 'scroll',
          orient: 'vertical',
          right: '5%',
          top: '40%',
          itemWidth: 10,
          itemHeight: 5,
          itemGap: 10,
          textStyle: {
            color: ['#1abffa','#11a754'],
            fontSize:12
          },
          data: ['会过 ','卷皮 ']
        },
        series: [
          {
            name: '卷皮 ',
            type: 'pie',
            clockWise: false,
            radius: ['50%', '55%'],
            center: ['30%','50%'],
            clockWise: false, //顺时加载
            hoverAnimation: false, //鼠标移入变大
            itemStyle: {
              normal: {
                label: {
                  show: false
                },
                labelLine: {
                  show: false
                },
                borderWidth: 5,
                shadowBlur: 40,
                borderColor: "#1c76f2",
                shadowColor: 'rgba(0, 0, 0, 0)' //边框阴影
              }
            },
            data: [
              {
                value: 50,
                name: '卷皮'
              },
              {
                value: 50*20/100,
                name: '',
                itemStyle: placeHolderStyle
              },
              {
                value:50*30/100,
                name: '',
                itemStyle: {
                  normal: {
                    color: 'none',
                    borderColor:'none'
                  }
                }
              }
            ]
          }, {
            name:'会过 ',
            type: 'pie',
            clockWise: false,
            hoverAnimation: false,
            radius: ['30%', '35%'],
            center: ['30%','50%'],
            itemStyle: {
              normal: {
                label: {
                  show: false
                },
                labelLine: {
                  show: false
                },
                borderWidth: 5,
                shadowBlur: 40,
                borderColor: "#098c43",
                shadowColor: 'rgba(0, 0, 0, 0)' //边框阴影
              }
            },
            data: [
              {
                value: 80,
                name: '会过'
              },
              {
                value: 80*20/100,
                name: '',
                itemStyle: placeHolderStyle
              },
              {
                value:80*30/100,
                name: '',
                itemStyle: {
                  normal: {
                    color: 'none',
                    borderColor:'none'
                  }
                }
              }
            ]
          }]
      };
      this.myChart3.setOption(option3);
      var option4 = {
        title:{
          show:true,
          text:'异常访问',
          right:4,
          top:40,
          textStyle:{
            fontSize:16,
            color:'#ff3761'
          }
        },
        color: ['#ff3761'],
        tooltip: {
          trigger: 'item',
          show:false
          // formatter:"<p class='center'>{b}</p><p class='center'>{c}</p>"
        },
        legend: {
          type: 'scroll',
          orient: 'vertical',
          right: '15%',
          top: '40%',
          itemWidth: 10,
          itemHeight: 5,
          itemGap: 10,
          textStyle: {
            color: '#ff3761',
            fontSize:12
          },
          data: ['404']
        },
        series: [{
          name: '404',
          type: 'pie',
          clockWise: false,
          radius: ['50%', '55%'],
          center: ['30%','50%'],
          hoverAnimation: false,
          itemStyle: {
            normal: {
              label: {
                show: false
              },
              labelLine: {
                show: false
              },
              borderWidth: 5,
              shadowBlur: 40,
              borderColor: "#ff3761",
              shadowColor: 'rgba(0, 0, 0, 0)' //边框阴影
            }
          },
          data: [
            {
              value: 80,
              name: '异常访问'
            },
            {
              value: 80*20/100,
              name: '',
              itemStyle: placeHolderStyle
            } ,{
              value:80*30/100,
              name: '',
              itemStyle: {
                normal: {
                  color: 'none',
                  borderColor:'none'
                }
              }
            }
          ]
        }]
      };
      this.myChart4.setOption(option4);
    },
    initHGData() {
      $('.sg-distribution .filter-type li').eq(this.hgGraphNum).addClass('active').siblings().removeClass('active')
      this.hgGraphNum ++
      fetchHgGraphData().then(response => {
        this.hgGraphData = response.data
        this.hgGraph(response.data, 'in')
      }).catch(err => {
        console.log(err)
      })
    },
    getHgGraphData () {
      fetchHgGraphData().then(response => {
        this.hgGraphData = response.data
      }).catch(err => {
          console.log(err)
      })
    },
    hgButtonTime() {
      $('.sg-distribution .filter-type li').eq(this.hgGraphNum).addClass('active').siblings().removeClass('active')
      var dataType = $('.sg-distribution .filter-type li.active').data('type')
      if(dataType == 'in'){
        this.hgGraph(this.hgGraphData, 'in')
      }else if(dataType == 'out'){
        this.hgGraph(this.hgGraphData, 'out')
      }
      this.hgGraphNum ++
      if(this.hgGraphNum >= 2){
        this.hgGraphNum = 0
      }
    },
    hgGraph(hgdata, type) {
      clearInterval(this.Timer5)
      var data
      if (type == 'in') {
        data = hgdata['intraffic']
      } else {
        data = hgdata['outtraffic']
      }
      var category = [], roomCnt = [], s
      for(var i=0;i<data.length;i++){
        s = new Date(data[i].timestamp * 1000)
        category.push(this.add_zero(s.getHours()) + ':' + this.add_zero(s.getMinutes()))
        roomCnt.push(data[i].value);
      }
      var option5= {
        title: {},
        tooltip: {
          trigger: 'item',
          transitionDuration:0,
          backgroundColor : 'rgba(83,93,105,0.8)',
          borderColor : '#535b69',
          borderRadius : 8,
          borderWidth: 2,
          padding: [5,10],
          showContent: true,
          formatter:'{b} : {c}'
        },
        grid: {
          show: false,
          top:'6%',
          left: '8%',
          right:'5%'
        },
        xAxis: {
          type: 'category',
          boundaryGap: false,
          axisLine: {
            show:false,
            lineStyle: {
              color: '#a5a6bb'
            }
          },
          axisLabel: {
            interval:0,
            show: true,
            inside: false,
            rotate: 45,
            margin: 2,
            textStyle: {
              fontSize: 12,
              color: '#a5a6bb'
            }
          },
          splitLine: {
            show:false
          },
          axisTick: {
            show: false
          },
          data: category,
          offset:15
        },
        yAxis: {
          type: 'value',
          axisLine: {
            show:false,
            lineStyle: {
              color: '#a5a6bb'
            }
          },
          axisLabel: {
            margin:10,
            textStyle: {
              fontSize: 12,
              color: '#a5a6bb'
            }
          },
          splitLine: {
            lineStyle: {
              color: '#273651',
              width:1,
              shadowColor: 'rgba(0, 0, 0, 0.7)',
              shadowBlur: 10,
              shadowOffsetX: 5,
              opacity:'0.3'
            }
          },
          axisTick: {
            show: false
          }
        },
        series: [{
          type: 'line',
          smooth: false,
          symbolSize:6,
          lineStyle: {
            normal: {
              width:2
            }
          },
          areaStyle: {
            normal: {
              color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [{
                offset: 0,
                color: 'rgba(25, 153, 227, 0.7)'
              }, {
                offset: 0.8,
                color: 'rgba(25, 153, 227, 0)'
              }], false),
              shadowColor: 'rgba(0, 0, 0, 0.1)',
              shadowBlur: 10
            }
          },
          itemStyle: {
            normal: {
              color: '#1999e3'
            }
          },
          data:roomCnt
        }]
      };
      this.myChart5.setOption(option5)
      this.Timer5 = this.hoverTurn(option5, this.myChart5, this.option5Index);
    },
    initFailData() {
      var _this = this
      var season = 'Q1'
      clearInterval(this.Timer8)
      clearInterval(this.Timer9)
      var list = {
        "Q1": [0, 1, 2],
        "Q2": [3, 4, 5],
        "Q3": [6, 7, 8],
        "Q4": [9, 10, 11]
      }
      var date = new Date()
      var month = date.getMonth()
      for(var s in list) {
        if ($.inArray(month, list[s]) >= 0) {
          season = s
          break
        }
      }
      var failUrl = 'http://uwork.juanpi.org/uwork/uwatchFetchData/failure'
      axios.get(failUrl, {
        params: {
          season: season
        }
      }).then(function (response) {
        var data = JSON.parse(response.data)
        var fail = data.data
        var dataType = [],typeAll = fail.length
        var dataTime = [],timeAll = 0
        var types = {}
        for (var i = 0; i < fail.length; i ++) {
          var tname = fail[i].ftype
          if (types.hasOwnProperty(tname)) {
            types[tname].count += 1
            types[tname].affect += fail[i].affect_time
          } else {
            types[tname] = {'count': 1, 'affect': fail[i].affect_time}
          }
          timeAll += fail[i]['affect_time']
        }

        for(var key in types) {
          dataType.push({'name': key, 'value': types[key].count})
          dataTime.push({'name': key, 'value': types[key].affect})
        }
        $('.echart-text li:first-child').append('<p>' + typeAll + '</p>')
        var option8= {
          tooltip: {
            trigger: 'item',
            transitionDuration: 0,
            backgroundColor: 'rgba(83,93,105,0.8)',
            borderColor: '#535b69',
            borderRadius: 8,
            borderWidth: 2,
            padding: [5, 10],
            formatter: "<p class='center'>{b}</p><p class='center'>{c}({d}%)</p>"
          },
          title: {
            text: '',
            textStyle: {
              color: '#fff',
              fontSize: '18'
            },
            x: 'center',
            y: 'center'
          },
          series: [
            {
              type: 'pie',
              radius: ['50%', '65%'],
              avoidLabelOverlap: true,
              hoverAnimation: true,
              label: {
                normal: {
                  show: false,
                  position: 'center',
                  color: '#fff',
                  fontSize: '8'
                },
                emphasis: {
                  show: false,
                  formatter: '{b} \n\n{d}%',
                  textStyle: {
                    fontSize: '18',
                    fontWeight: 'bold'
                  }
                }

              },
              labelLine: {
                normal: {
                  show: false
                }
              },
              data: dataType
            }
          ]
        }
        _this.myChart8.setOption(option8)
        _this.Timer8 = _this.hoverTurn(option8, _this.myChart8, _this.option8Index)

        $('.echart-text li:nth-child(2)').append('<p>' + timeAll + '</p>');
        var option9 = {
          tooltip: {
            trigger: 'item',
            transitionDuration:0,
            backgroundColor : 'rgba(83,93,105,0.8)',
            borderColor : '#535b69',
            borderRadius : 8,
            borderWidth: 2,
            padding: [5,10],
            formatter: "<p class='center'>{b}</p><p class='center'>{c}({d}%)</p>"
          },
          title: {
            text: '',
            textStyle: {
              color: '#fff',
              fontSize: '18'
            },
            x: 'center',
            y: 'center'
          },
          series: [
            {
              type:'pie',
              radius: ['45%', '65%'],
              hoverAnimation: false,
              avoidLabelOverlap: true,
              label: {
                normal: {
                  show: false,
                  position: 'center',
                  color:'#fff',
                  fontSize: '8'
                },
                emphasis: {
                  show: false,
                  textStyle: {
                    fontSize: '18',
                    fontWeight: 'bold'
                  }
                }
              },
              labelLine: {
                normal: {
                  show: false
                }
              },
              data: dataTime
            }
          ]
        };
        _this.myChart9.setOption(option9)
        _this.Timer9 = _this.hoverTurn(option9, _this.myChart9, _this.option9Index)

        var lists = {
          'Q1': ['一月', '二月', '三月'],
          'Q2': ['四月', '五月', '六月'],
          'Q3': ['七月', '八月', '九月'],
          'Q4': ['十月', '十一月', '十二月']
        }
        var TplD = '', HgTplD = '', Tpl = '', avg = 0, Hgavg = 0
        for(key in list[season]) {
          TplD += '<span>' + parseFloat(data.sla_per_mon[list[season][key]]).toFixed(2) + '</span>'
          HgTplD += '<span>' + parseFloat(data.sla_per_monhg[list[season][key]]).toFixed(2) + '</span>'
          Tpl += '<span>' + lists[season][key] + '</span>'
          avg += parseFloat(data.sla_per_mon[list[season][key]])
          Hgavg += parseFloat(data.sla_per_monhg[list[season][key]])
        }
        $('.this-week-status .data-count').html(TplD)
        $('.use-rate .data-count').html(HgTplD)
        $('.this-week-status .status').html(Tpl)
        $('.use-rate .status').html(Tpl)
        $('.this-week-status .status_title span').html((avg/3).toFixed(2))
        $('.use-rate .status_title span').html((Hgavg/3).toFixed(2))
      }).catch(function (error) {
        console.log(error)
      })
    },
    initGrand() {
      $('.grand .filter-type li').eq(this.grandNum).addClass('active').siblings().removeClass('active')
      var dataType = $('.grand .filter-type li.active').data('type')
      if(dataType == 'cpu'){
        this.grand(this.grandData, 'cpu')
      }else if(dataType == 'mem'){
        this.grand(this.grandData, 'mem')
      }
      this.grandNum ++
      if(this.grandNum >= 2){
        this.grandNum = 0
      }
    },
    grand(data, dataType) {
      var data = [
        {
          "resourceCnt": 91,
          "schoolName": "10.205.136.30"
        },
        {
          "resourceCnt": 84,
          "schoolName": "10.205.136.33"
        },
        {
          "resourceCnt": 84,
          "schoolName": "10.205.136.45"
        },
        {
          "resourceCnt": 83,
          "schoolName": "10.205.136.27"
        },
        {
          "resourceCnt": 82,
          "schoolName": "10.205.130.38"
        },
        {
          "resourceCnt": 82,
          "schoolName": "10.205.130.44"
        },
        {
          "resourceCnt": 81,
          "schoolName": "10.205.136.19"
        },
        {
          "resourceCnt": 80,
          "schoolName": "10.205.136.23"
        }
      ]
      var schoolName = [], resourceCnt = [], total = 0, max = 0;
      if (data) {
        for (var i = 0; i < data.length; i++) {
          schoolName.push(data[i].schoolName);
          resourceCnt.push(data[i].resourceCnt);
          total += data[i].resourceCnt;
        }
        var maxData = [100, 100, 100, 100, 100, 100, 100, 100];
      }

      var option10 = {
        grid: {
          left: '40',
          top: '0',
          right: '24',
          bottom: '10',
          containLabel: true
        },
        xAxis: [{
          show: false
        }],
        yAxis: [{
          splitLine: {show: false},
          axisLine: {
            show: false,
            lineStyle: {
              color: '#1798cf'
            }
          },
          axisLabel: {
            fontSize: '16'
          },
          axisTick: {
            show: false
          },
          offset: 20,
          data: schoolName.reverse()
        }, {
          splitLine: {show: false},
          axisLine: {
            show: false,
            lineStyle: {
              color: '#fff'
            }
          },
          axisLabel: {
            fontSize: '24',
            fontFamily: 'myFirstFont'
          },
          axisTick: {
            show: false
          },
          offset: 50,
          data: resourceCnt.reverse()
        }, {
          name: '',
          nameGap: '50',
          nameTextStyle: {
            color: '#ffffff',
            fontSize: '24'
          },
          axisLine: {
            lineStyle: {
              color: 'rgba(0,0,0,0)'
            }
          },
          data: []
        }],
        series: [{
          name: '条',
          type: 'bar',
          yAxisIndex: 0,
          data: resourceCnt,
          label: {
            normal: {
              show: false,
              position: 'right',
              formatter: function (param) {
                // return param.value + '%';
                return param.value;
              },
              textStyle: {
                color: '#fff',
                fontSize: '16'
              }
            }
          },
          barWidth: 6,
          itemStyle: {
            normal: {
              barBorderRadius: 6,
              color: new echarts.graphic.LinearGradient(0, 0, 1, 0, [{
                offset: 0,
                color: '#1d6bf1'
              }, {
                offset: 1,
                color: '#1ad0fc'
              }])
            }
          },
          z: 2
        }, {
          name: '白框',
          type: 'bar',
          yAxisIndex: 1,
          barGap: '-100%',
          data: maxData,
          barWidth: 10,
          itemStyle: {
            normal: {
              color: '#0a182e',
              barBorderRadius: 10
            }
          },
          z: 1
        }, {
          name: '外框',
          type: 'bar',
          yAxisIndex: 2,
          barGap: '-100%',
          data: maxData,
          barWidth: 12,
          itemStyle: {
            normal: {
              color: "#4e6a8c",
              borderColor: '#4e6a8c',
              barBorderRadius: 12
            }
          },
          z: 0
        }, {
          name: '内圆',
          type: 'scatter',
          hoverAnimation: false,
          data: [0, 0, 0, 0, 0, 0, 0, 0],
          yAxisIndex: 2,
          symbolSize: 14,
          itemStyle: {
            normal: {
              color: '#1367fb',
              opacity: 1
            }
          },
          z: 3
        }, {
          name: '外圆',
          type: 'scatter',
          hoverAnimation: true,
          data: [0, 0, 0, 0, 0, 0, 0, 0],
          yAxisIndex: 2,
          symbolSize: 20,
          itemStyle: {
            normal: {
              color: '#0a182e',
              borderColor: '#4e6a8c',
              opacity: 1
            }
          },
          z: 0
        }]
      };
      this.myChart10.setOption(option10);
    },
    initResourceData() {
      var data = {
        "officeTotal": 360,
        "idcTotal": 445
      }
      var tplOne = '', tplTwo = '', Tpl = '<li>KVM' +
          '<p class="top"></p>' +
          '<p class="bottom"></p>' +
          '</li>' +
          '<li>Docker' +
          '<p class="top"></p>' +
          '<p class="bottom"></p></li>' +
          '<li>VmWare' +
          '<p class="top"></p>' +
          '<p class="bottom"></p>' +
          '</li>',
        idcTotal, officeTotal;

      if (data) {
        idcTotal = data.idcTotal.toString();
        officeTotal = data.officeTotal.toString();
        for (var i = 0; i < idcTotal.length; i++) {
          tplOne += '<li>' + idcTotal[i] + '</li>';
        }
        for (var i = 0; i < officeTotal.length; i++) {
          tplTwo += '<li>' + officeTotal[i] + '</li>';
        }
        Tpl = '<li>KVM' +
          '<p class="top">' + 86 + '</p>' +
          '<p class="bottom">' + 10 + '</p>' +
          '</li>' +
          '<li>Docker' +
          '<p class="top">' + 111 + '</p>' +
          '<p class="bottom">' + 24 + '</p></li>' +
          '<li>VmWare' +
          '<p class="top">' + 249 + '</p>' +
          '<p class="bottom">' + 13 + '</p>' +
          '</li>' +
          '<li>Other' +
          '<p class="top">' + 0 + '</p>' +
          '<p class="bottom">' + 59 + '</p>' +
          '</li>';
      }

      $('.re-top-right .six-border').append(Tpl);
      $('.re-all-data.one').html(tplOne);
      $('.re-all-data.two').html(tplTwo)

      var pieData = [
        {
          "resourceCnt": 361,
          "subjectName": "Ucloud物理机"
        },
        {
          "resourceCnt": 41,
          "subjectName": "Ucloud云主机"
        },
        {
          "resourceCnt": 15,
          "subjectName": "Aliyun云主机"
        },
        {
          "resourceCnt": 28,
          "subjectName": "Qcloud物理机"
        },
        {
          "resourceCnt": 138,
          "subjectName": "Qcloud云主机"
        },
        {
          "resourceCnt": 186,
          "subjectName": "深圳Office"
        },
        {
          "resourceCnt": 174,
          "subjectName": "武汉Office"
        }
      ]
      var subjectName = [];
      if (pieData) {
        for (var i = 0; i < pieData.length; i++) {
          subjectName.push({name: pieData[i].subjectName, value: pieData[i].resourceCnt, selected: true});
        }
      }
      var option6 = {
        color: ['#4bc965', '#e0b455', '#eb4d4e', '#9c90ff', '#2774ba', '#2ecec4', '#abb659', '2a86aa', '2a86aa'],
        tooltip: {
          trigger: 'item',
          transitionDuration:0,
          backgroundColor : 'rgba(83,93,105,0.8)',
          borderColor : '#535b69',
          borderRadius : 8,
          borderWidth: 2,
          padding: [5,10],
          formatter: "{b}: {c} ({d}%)"
        },
        series: [
          {    // 饼图的属性配置
            name: 'outPie',
            type: 'pie',
            center: ['52%', '40%'],
            radius: ['70%', '65.5%'],
            hoverAnimation: false,
            // avoidLabelOverlap: false,
            startAngle: 20,
            zlevel: 0,
            itemStyle: {
              normal: {
                borderColor: "#192342",
                borderWidth: 10
              },
              emphasis: {
                shadowBlur: 2,
                shadowOffsetX: 0,
                shadowColor: 'rgba(0, 0, 0, 0.8)'
              }
            },
            // 图形样式
            label: {
              normal: {
                show: true,
                formatter: function (param) {
                  return param.name + '(' + param.value + ')';
                },
                position: 'top',
                fontSize: '14'
              },
              emphasis: {
                show: true,
                formatter: function (param) {
                  return param.name + '(' + param.value + ')';
                },
                textStyle: {
                  fontSize: '16',
                  fontWeight: 'bold'
                }
              }
            },
            labelLine: {
              normal: {
                show: false
              }
            },
            data: subjectName
          }
        ]
      };
      this.myChart6.setOption(option6);
    },
    initAlarmData() {
      fetchAlarmData().then(response => {
        var data = response.data
        var Tpl=''
        for(var i = 0; i < 5; i++){
          Tpl +=
            '<tr>' +
            '<td><div style="width:98px;">' + data[i].Time + '</td>' +
            '<td><div style="width:50px;">' + data[i].Level + '</td>' +
            '<td><div style="width:100px;">' + data[i].Status + '</td>' +
            '<td><div style="width:200px;">' + data[i].Name + '</td>' +
            '<td><div style="width:350px;">' + data[i].Flag + '</td>' +
            '</tr>';
        }
        $('#main7').html(Tpl);
      }).catch(err => {
        console.log(err)
      })
    },
    changeType(e) {
      var _this = $(e.target)
      _this.addClass('active').siblings().removeClass('active')
      var dataType = _this.data('type')
      if(dataType == 'in'){
        this.hgGraph(this.hgGraphData, 'in')
      }else if(dataType=='out'){
        this.hgGraph(this.hgGraphData, 'out')
      }else if(dataType == 'cpu'){
        this.grand(this.grandData, 'cpu')
      }else if(dataType == 'mem'){
        this.grand(this.grandData, 'mem')
      }
    },
    hoverTurn(a,b,c){
      var intervalVal =setInterval(function () {
        var dataLen = a.series[0].data.length;
        // 取消之前高亮的图形
        b.dispatchAction({
          type: 'downplay',
          seriesIndex: 0,
          dataIndex: c
        });
        c = (c + 1) % dataLen;
        // 高亮当前图形
        b.dispatchAction({
          type: 'highlight',
          seriesIndex: 0,
          dataIndex: c
        });
        // 显示 tooltip
        b.dispatchAction({
          type: 'showTip',
          seriesIndex: 0,
          dataIndex: c
        });
      }, 2000);
      return intervalVal;
    },
    getNowFormatDate() {
      var date = new Date();
      var year = date.getFullYear();
      var month = date.getMonth() + 1;
      var strDate = date.getDate();
      var Hour =  date.getHours();       // 获取当前小时数(0-23)
      var Minute =  date.getMinutes();     // 获取当前分钟数(0-59)
      var Second = date.getSeconds();     // 获取当前秒数(0-59)
      var show_day=new Array('星期日','星期一','星期二','星期三','星期四','星期五','星期六');
      var day=date.getDay();
      if (Hour<10) {
        Hour = "0" + Hour;
      }
      if (Minute <10) {
        Minute = "0" + Minute;
      }
      if (Second <10) {
        Second = "0" + Second;
      }
      if (month >= 1 && month <= 9) {
        month = "0" + month;
      }
      if (strDate >= 0 && strDate <= 9) {
        strDate = "0" + strDate;
      }
      var currentdate = '<div><p>'+year + '年' + month +'月' + strDate+'号</p><p>'+show_day[day]+'</p></div>';
      var HMS = Hour + ':' + Minute +':' + Second;
      $('.nowTime li:nth-child(1)').html(HMS);
      $('.nowTime li:nth-child(2)').html(currentdate);
    },
    add_zero(val) {
      var ret = val
      if (val < 10) {
        ret = '0'+val
      }
      return ret
    },
    fullScreen() {
      var elem = document.body
      if (elem.webkitRequestFullScreen) {
        elem.webkitRequestFullScreen()
      } else if (elem.mozRequestFullScreen) {
        elem.mozRequestFullScreen()
      } else if (elem.requestFullScreen) {
        elem.requestFullscreen()
      } else {
        notice.notice_show("浏览器不支持全屏API或已被禁用", null, null, null, true, true)
      }
      this.showFullscreenButton = false
    }
  }
}