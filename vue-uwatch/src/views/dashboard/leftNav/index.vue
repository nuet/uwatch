<template>
  <div class="col-md-2" style="font-size:12px;height:900px;margin-top:20px;">
    <el-col :span="24">
      <el-menu
        default-active="2"
        class="el-menu-vertical-demo"
        @open="handleOpen"
        @close="handleClose"
        background-color="#545c64"
        text-color="#fff"
        active-text-color="#ffd04b">
        <el-submenu index="1">
          <template slot="title">
            <svg-icon icon-class="chart" />
            <span>流量图</span>
          </template>
          <el-menu-item-group>
            <el-menu-item index="1-1" @click="showdata('http://grafana1.juanpi.org/dashboard/db/idcxun-jian-ri-bao-liu-liang-tu',1)">IDC巡检日报流量图</el-menu-item>
            <el-menu-item index="1-2" @click="showdata('http://grafana1.juanpi.org/dashboard/db/ge-yu-ming-liu-liang-tong-ji',2)">各域名流量统计</el-menu-item>
            <el-menu-item index="1-3" @click="showdata('http://grafana1.juanpi.org/dashboard/db/juan-pi-ban-gong-wang-luo-liu-liang-jian-kong',3)">卷皮办公网流量监控图</el-menu-item>
            <el-menu-item index="1-4" @click="showdata('http://grafana1.juanpi.org/dashboard/db/shen-zhen-wu-yi-dian-xin-chang-kuan-jiao-huan-ji-pingjian-ce-tu',4)">办公网Ping检测图</el-menu-item>
            <el-menu-item index="1-5" @click="showdata('http://grafana1.juanpi.org/dashboard/db/ufileliu-liang-jian-kong',5)">Ufile流量监控</el-menu-item>
          </el-menu-item-group>
        </el-submenu>

        <el-submenu index="2">
          <template slot="title">
            <svg-icon icon-class="dashboard" />
            <span>可用性检测</span>
          </template>
          <el-menu-item-group>
            <el-menu-item index="2-1" @click="showdata('http://grafana1.juanpi.org/dashboard/db/lvsfang-wen-tong-ji',6)">LVS访问统计</el-menu-item>
            <el-menu-item index="2-2" @click="showdata('http://grafana1.juanpi.org/dashboard/db/vts_yuan-zhan-ke-yong-xing',7)">Nginx源站可用性</el-menu-item>
            <el-menu-item index="2-3" @click="showdata('http://grafana1.juanpi.org/dashboard/db/vts_yuan-zhan-ke-yong-xing',8)">vts_源站可用性</el-menu-item>
          </el-menu-item-group>
        </el-submenu>

        <el-submenu index="3">
          <template slot="title">
            <svg-icon icon-class="example" />
            <span>网宿CDN</span>
          </template>
          <el-menu-item-group>
            <el-menu-item index="3-1" @click="showdata('http://grafana1.juanpi.org/dashboard/db/wang-su-cdnliu-liang-tu')">CDN流量图</el-menu-item>
            <el-menu-item index="3-2" @click="showdata('http://grafana1.juanpi.org/dashboard/db/wang-su-cdnyuan-zhan-ke-yong-lu')">源站可用率</el-menu-item>
            <el-menu-item index="3-3" @click="showdata('http://grafana1.juanpi.org/dashboard/db/wang-su-cdnyuan-zhan-xiang-ying-shi-jian')">源站响应时间</el-menu-item>
            <el-menu-item index="3-4" @click="showdata('http://grafana1.juanpi.org/dashboard/db/wang-su-cdnhui-yuan-shi-jian-tong-ji')">回源时间统计</el-menu-item>
            <el-menu-item index="3-5" @click="showdata('http://grafana1.juanpi.org/dashboard/db/teng-xun-cdnke-yong-lu-tong-ji')">可用率统计</el-menu-item>
            <el-menu-item index="3-6" @click="showdata('http://grafana1.juanpi.org/dashboard/db/teng-xun-cdnzhuang-tai-ma-tong-ji')">状态码统计</el-menu-item>
          </el-menu-item-group>
        </el-submenu>
      </el-menu>
    </el-col>
  </div>
</template>
<script type="text/javascript">
  import '@/assets/dashboard/bootstrap/css/bootstrap.css' // global css
  import '@/assets/dashboard/style.css'
  export default{
    name: 'leftNav',
    data() {
      return {
        load_data: false,
        user_data: [],
        href: '',
        showGrafana: false,
        search: true
      }
    },
    components: {},
    created() {
      this.show()
    },
    methods: {
      show() {
        this.href = 'http://grafana1.juanpi.org/dashboard/db/idcxun-jian-ri-bao-liu-liang-tu'
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
      handleOpen(key, keyPath) {
        console.log(key, keyPath)
      },
      handleClose(key, keyPath) {
        console.log(key, keyPath)
      }
    }
  }
</script>
<style scoped>
  .clickshow {
    background: #C4E3F3;
    border: 1px solid #1E90FF;
  }
</style>