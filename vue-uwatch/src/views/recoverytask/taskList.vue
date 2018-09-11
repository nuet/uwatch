<template>
  <div class="app-container calendar-list-container">
    <el-alert
      title=""
      type="success"
      :closable="false">
      <div style="padding:10px;color:#666;">
        <el-date-picker
          v-model="daterange"
          type="daterange"
          range-separator="-"
          start-placeholder="开始日期"
          end-placeholder="结束日期"
          align="right"
          unlink-panels
          :picker-options="pickerOptions">
          @change="getList"
        >
        </el-date-picker>
        <el-select  placeholder="状态" v-model="listQuery.query.status" style="width: 120px" @change="getList">
          <el-option
            v-for="item in status_list"
            :key="item.value"
            :label="item.label"
            :value="item.value">
          </el-option>
        </el-select>
        <el-input  style="width: 200px;" v-model="listQuery.query.source" class="filter-item" placeholder="请输入IP" ><el-button slot="append" icon="el-icon-search" @click="getList"></el-button></el-input>
        <el-button  class="filter-item" @click="clean">
          <i class="fa fa-refresh" ></i>
        </el-button>
        <span class="s-data-bar s-item2">
                <i></i>共&nbsp;<em id="instance_count">{{total}}</em>&nbsp;次自愈
        </span>
      </div>
    </el-alert>
    <el-table :key='tableKey' :data="list" v-loading="listLoading" element-loading-text="给我一点时间" border fit highlight-current-row
              style="width: 100%">
      <el-table-column  align="center" label="类型" show-overflow-tooltip>
        <template slot-scope="scope">
          <span v-if="scope.row.AutoFill">{{scope.row.AutoFill.Title}}</span>
        </template>
      </el-table-column>
      <el-table-column width="160px"  align="center" label="产生时间">
        <template slot-scope="scope">
          <span>{{GMTToStr(scope.row.CreateTime)}}</span>
        </template>
      </el-table-column>
      <el-table-column width="80px" align="center" label="自愈耗时(ms)">
        <template slot-scope="scope">
          <span >{{scope.row.Duration}}</span>
        </template>
      </el-table-column>
      <el-table-column  width="160px" align="center" label="IP">
        <template slot-scope="scope">
          <span>{{scope.row.Source}}</span>
        </template>
      </el-table-column>
      <el-table-column width="300px" align="center" label="告警信息">
        <template slot-scope="scope">
          <span>{{scope.row.Item}}</span>
        </template>
      </el-table-column>
      <el-table-column  width="160px" align="center" label="状态">
        <template slot-scope="scope">
          <span v-if="scope.row.Status == 0">新增</span>
          <span v-else-if="scope.row.Status == 1">执行中</span>
          <span v-else-if="scope.row.Status == 2">执行成功</span>
          <span v-else="scope.row.Status == 3">执行失败</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="操作" width="120" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <router-link :to="{name: 'taskDetail', params: {id: scope.row.Id}}" >
            <el-button type="primary" icon="el-icon-search" style="width:100px" size="mini">查看详情</el-button>
          </router-link>
        </template>
      </el-table-column>
    </el-table>

    <div class="pagination-container">
      <el-pagination background @size-change="handleSizeChange" @current-change="handleCurrentChange" :current-page="listQuery.page" :page-sizes="[10,20,30, 50]" :page-size="listQuery.limit" layout="total, sizes, prev, pager, next, jumper" :total="total">
      </el-pagination>
    </div>


  </div>
</template>

<script>
  import { getTaskList } from '@/api/uwTask'

  const defaultQuery = {
    status: '',
    source: '',
    start: '',
    end: ''
  }

  export default {
    name: 'complexTable',
    data() {
      return {
        tableKey: 0,
        list: null,
        total: null,
        listLoading: true,
        listQuery: {
          page: 1,
          limit: 20,
          query: Object.assign({}, defaultQuery)
        },
        status_list: [
          {
            value: '0',
            label: '新增'
          }, {
            value: '1',
            label: '执行中'
          }, {
            value: '2',
            label: '执行成功'
          }, {
            value: '3',
            label: '执行失败'
          }
        ],
        daterange: [
          new Date(new Date().getTime() - 24 * 60 * 60 * 1000),
          new Date()
        ],
        pickerOptions: {
          shortcuts: [{
            text: '最近一周',
            onClick(picker) {
              const end = new Date()
              const start = new Date()
              start.setTime(start.getTime() - 3600 * 1000 * 24 * 7)
              picker.$emit('pick', [start, end])
            }
          }, {
            text: '最近一个月',
            onClick(picker) {
              const end = new Date()
              const start = new Date()
              start.setTime(start.getTime() - 3600 * 1000 * 24 * 30)
              picker.$emit('pick', [start, end])
            }
          }, {
            text: '最近三个月',
            onClick(picker) {
              const end = new Date()
              const start = new Date()
              start.setTime(start.getTime() - 3600 * 1000 * 24 * 90)
              picker.$emit('pick', [start, end])
            }
          }]
        }
      }
    },
    created() {
      this.getList()
    },
    methods: {
      getList() {
        this.listLoading = true
        this.listQuery.query.start = this.daterange[0]
        this.listQuery.query.end = this.daterange[1]
        getTaskList(this.listQuery).then(response => {
          this.list = response.data.list
          this.total = response.data.total
          this.listLoading = false
        })
      },
      handleFilter() {
        this.listQuery.page = 1
        this.getList()
      },
      handleSizeChange(val) {
        this.listQuery.limit = val
        this.getList()
      },
      handleCurrentChange(val) {
        this.listQuery.page = val
        this.getList()
      },
      clean() {
        this.listQuery.query = Object.assign({}, defaultQuery)
        this.daterange = [
          new Date(new Date().getTime() - 24 * 60 * 60 * 1000),
          new Date()
        ]
        this.getList()
      },
      GMTToStr(time) {
        const date = new Date(time)
        const Str = date.getFullYear() + '-' +
          (date.getMonth() + 1) + '-' +
          date.getDate() + ' ' +
          date.getHours() + ':' +
          date.getMinutes() + ':' +
          this.checkTime(date.getSeconds())
        return Str
      },
      checkTime(i) {
        if (i < 10) {
          i = '0' + i
        }
        return i
      }
    },
    watch: {
      daterange(val) {
        this.getList()
      }
    }
  }
</script>
<style>
  .s-data-bar {
    display: inline-block;
    line-height: 32px;
    vertical-align: top;
    font-size: 14px;
    color: #337ab7;
    margin-left: 20px;
  }
  .s-item2 i {
    background: #ffc261;
  }
  .s-data-bar i {
    display: inline-block;
    width: 8px;
    height: 8px;
    -webkit-border-radius: 50%;
    border-radius: 50%;
    margin-right: 10px;
  }
  .s-data-bar em {
    display: inline-block;
    vertical-align: -1px;
    font-style: normal;
    font-size: 14px;
    color: #666;
  }
</style>
