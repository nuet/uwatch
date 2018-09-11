<template>
  <div>
    <div class="container-header">
      <ul class="nowTime">
        <li></li>
        <li></li>
      </ul>
      <div class="location">
        <i class="icon iconfont icon-buoumaotubiao23"></i>
        <span class="areaName">深圳</span>
      </div>
      <h3>Uwatch监控大屏</h3>
    </div>
    <div class="container-content">
      <div class="content">
        <div class="main">
          <div class="resource-right">
            <!--<img src="images/teaching-resouce.gif" alt="">-->
            <div class="com-count-title" style="padding: 10px 20px;">IDC资源总量
              <ul class="re-all-data one">

              </ul>
            </div>
            <div class="com-count-title" style="padding: 0 20px">办公网资源总量
              <ul class="re-all-data two">

              </ul>
            </div>
            <div class="re-top-right">
              <ul class="six-border">
                <span class="data-type">虚拟机</span>
                <span class="data-type-two">宿主机</span>
              </ul>
            </div>
            <div class="com-screen-content">
              <div id="main6" style="width:100%;height:540px;"></div>
              <div class="inner-pie">
                <table>
                  <tbody>
                  <!--<tr><td>视频：</td><td>20.4%</td></tr>-->
                  <!--<tr><td>文档：</td><td>26.0%</td></tr>-->
                  <!--<tr><td>图片：</td><td>26.7%</td></tr>-->
                  <!--<tr><td>音频：</td><td>26.9%</td></tr>-->
                  </tbody>
                </table>
              </div>
            </div>
            <span class="left-top"></span>
            <span class="right-top"></span>
            <span class="left-bottom"></span>
            <span class="right-bottom"></span>
          </div>
          <div class="acting-school">
            <div class="com-count-title">最新告警</div>
            <div class="com-screen-content">
              <table class="course-table mt30">
                <thead>
                <tr>
                  <td style="width:98px;">时间</td>
                  <td style="width:50px;">级别</td>
                  <td style="width:100px;">状态</td>
                  <td style="width:200px;">机器标识</td>
                  <td style="width:350px;">标签</td>
                </tr>
                </thead>
                <tbody id="main7">
                </tbody>
              </table>
            </div>
            <span class="left-top"></span>
            <span class="right-top"></span>
            <span class="left-bottom"></span>
            <span class="right-bottom"></span>
          </div>
        </div>
      </div>
      <div class="left">
        <div class="class-build">
          <div class="com-count-title">源站可用率</div>
          <div class="com-screen-content">
            <div class="class-build-left">
              <div id="main1" style="width:100%;height:290px;"></div>
            </div>
            <div class="class-build-right">
              <div id="main2" style="width:100%;height:290px;"></div>
            </div>
            <ul class="echart-title">
              <li>会过</li>
              <li>卷皮</li>
            </ul>
          </div>
          <span class="left-top"></span>
          <span class="right-top"></span>
          <span class="left-bottom"></span>
          <span class="right-bottom"></span>
        </div>
        <div class="living-status">
          <div class="com-count-title">实时动态</div>
          <div class="com-screen-content">
            <ul class="status-data">
            </ul>
            <ul class="status-name">
              <li>核心流量</li>
              <li>办公网流量</li>
              <li>会过流量</li>
              <li>卷皮流量</li>
            </ul>
            <div class="living-status-left ">
              <div id="main3" style="width:100%;height:185px;"></div>
            </div>
            <div class="living-status-right">
              <div id="main4" style="width:100%;height:185px;"></div>
            </div>
          </div>
          <span class="left-top"></span>
          <span class="right-top"></span>
          <span class="left-bottom"></span>
          <span class="right-bottom"></span>
        </div>
        <div class="sg-distribution">
          <ul class="filter-type">
            <li class="active" data-type="in" @click="changeType">会过入口流量</li>
            <li data-type="out" @click="changeType">会过出口流量</li>
          </ul>
          <div class="com-screen-content">
            <div id="main5" style="width:100%;height:250px;"></div>
          </div>
          <span class="left-top"></span>
          <span class="right-top"></span>
          <span class="left-bottom"></span>
          <span class="right-bottom"></span>
        </div>
      </div>
      <div class="right">
        <div class="teaching-analysis">
          <div class="com-count-title">故障报告-Q2</div>
          <div class="com-screen-content">
            <div class="class-build-left">
              <div id="main8" style="width:100%;height:290px;"></div>
            </div>
            <div class="class-build-right">
              <div id="main9" style="width:100%;height:290px;"></div>
            </div>
            <ul class="echart-text">
              <li><span>故障总数</span></li>
              <li><span>总影响时长</span></li>
            </ul>
            <ul class="echart-title">
              <li>故障分类</li>
              <li>故障影响</li>
            </ul>
          </div>
          <span class="left-top"></span>
          <span class="right-top"></span>
          <span class="left-bottom"></span>
          <span class="right-bottom"></span>

        </div>
        <div class="this-week">
          <div class="com-count-title">SLA-Q2</div>
          <div class="com-screen-content">
            <div class="course this-week-status">
              <div class="status_title">
                <span></span>
                <p>会过</p>
              </div>
              <ul class="">
                <li class="data-count">
                </li>
                <li class="status">
                </li>
              </ul>
            </div>
            <div class="course use-rate">
              <div class="status_title">
                <span></span>
                <p>卷皮</p>
              </div>
              <ul class="use-rate">
                <li class="data-count">
                </li>
                <li class="status">
                </li>
              </ul>
            </div>
          </div>
          <span class="left-top"></span>
          <span class="right-top"></span>
          <span class="left-bottom"></span>
          <span class="right-bottom"></span>
        </div>
        <div class="grand">
          <ul class="filter-type" style="padding: 15px;">
            <li class="active" data-type="cpu" @click="changeType">CPU使用排行</li>
            <li data-type="mem" @click="changeType">MEM使用排行</li>
          </ul>
          <div class="com-screen-content">
            <span class="rYName">-使用率-</span>
            <div id="main10" style="width:100%;height:368px;"></div>
          </div>
          <span class="left-top"></span>
          <span class="right-top"></span>
          <span class="left-bottom"></span>
          <span class="right-bottom"></span>
        </div>
      </div>
      <div class="clearfix"></div>
    </div>
    <!--<div v-if="showFullscreenButton" class="fullscreen-toolbox">-->
      <!--<button type="button" class="btn btn-secondary" @click="fullScreen">-->
        <!--<i class="fa fa-arrows-alt"></i>-->
        <!--全屏-->
      <!--</button>-->
    <!--</div>-->
  </div>
</template>
<style>
  body{
    min-width: 1880px;
    background: #081325 url(./images/interact.png) center no-repeat;
    background-size: 100% 100%;
    overflow: hidden;
  }
</style>
<script src="./js/index.js"></script>
<!--<script src="./js/teaching.js"></script>-->
