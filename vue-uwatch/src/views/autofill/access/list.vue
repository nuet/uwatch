<template>
  <div class="app-container calendar-list-container">
    <div class="filter-container">
      <el-alert
        title=""
        type="success"
        :closable="false">
        <div style="float:left;width:100%;color:#666;">
          <h2>
            <router-link :to="{ path:'add'}">
              <el-button type="success" icon="el-icon-plus">接入自愈</el-button>
            </router-link>
          </h2>针对每一种告警类型配置对应的处理方案。
        </div>
      </el-alert>
      <div style="margin-top:20px;">
        <el-input @keyup.enter.native="handleFilter" style="width: 200px;" class="filter-item" placeholder="自愈方案名称" v-model="QueryItems.Query"></el-input>
        <el-button @click.stop="on_refresh" class="filter-item">
          <i class="fa fa-refresh"></i>
        </el-button>
      </div>
    </div>
    <el-table :key='tableKey' :data="list" v-loading="listLoading" element-loading-text="给我一点时间" border fit highlight-current-row
              style="width: 100%">
      <el-table-column align="center" label="ID" width="65px">
        <template slot-scope="scope">
          <span>{{scope.row.Id}}</span>
        </template>
      </el-table-column>
      <el-table-column  align="center" label="自愈方案名称" show-overflow-tooltip>
        <template slot-scope="scope">
          <router-link :to="{ path:'edit/'+scope.row.Id}" >
            <span>{{scope.row.Title}}</span>
          </router-link>
        </template>
      </el-table-column>
      <el-table-column width="80px" align="center" label="告警数量">
        <template slot-scope="scope">
          <span>{{scope.row.Value}}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="自愈场景">
        <el-table-column width="330px" align="center" label="Metric" show-overflow-tooltip>
          <template slot-scope="scope">
            <span v-if="scope.row.Tag==''">{{scope.row.Metric}}</span>
            <span v-else>{{scope.row.Metric}}/{{scope.row.Tag}}</span>
          </template>
        </el-table-column>
        <el-table-column width="60px" align="center" label="比较">
          <template slot-scope="scope">
            <span>{{scope.row.Operator}}</span>
          </template>
        </el-table-column>
        <el-table-column width="110px" align="center" label="条件值">
          <template slot-scope="scope">
            <span>{{scope.row.Value}}</span>
          </template>
        </el-table-column>
      </el-table-column>
      </el-table-column>
      <el-table-column  width="160px" align="center" label="创建时间">
        <template slot-scope="scope">
          <span>{{scope.row.InsertTime}}</span>
        </template>
      </el-table-column>
      <el-table-column  width="160px" align="center" label="更新时间">
        <template slot-scope="scope">
          <span>{{scope.row.UpdateTime}}</span>
        </template>
      </el-table-column>
      <el-table-column width="80px" align="center" label="创建人">
        <template slot-scope="scope">
          <span>{{scope.row.Author}}</span>
        </template>
      </el-table-column>
      <el-table-column  width="80px" align="center" label="启用">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.Isvalid==1" type="danger">是</el-tag>
          <el-tag v-else type="info">否</el-tag>
        </template>
      </el-table-column>
      <el-table-column align="center" label="操作" width="330" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <el-popover ref="popover4" placement="left-start" width="620" trigger="click" >
            <div style="margin-left:10px;font-size:15px">
              <i class="el-icon-date" style="color:#20A0FF"></i>  <span style="margin-left:8px;color:#20A0FF">自愈方案详情</span>
            </div>
            <div class="login-bodya">
              <div class="loginWarpa">
                <div class="login-forma">
                  <el-form ref="form"  label-width="0">
                    <el-form-item  class="login-itema" style="margin-top:-20px;margin-left:-25px">
                      <label >方案名称 ：</label><span style="color:teal">{{projectData.Title}}</span>
                    </el-form-item>
                    <el-form-item prop="old_password" class="login-itema" style="margin-top:-20px;margin-left:-25px">
                      <label >自愈场景 ：</label>
                      <span style="color:teal">{{projectData.Metric}}</span>
                      <span style="color:teal" v-if="projectData.Tag!=''">/{{projectData.Tag}}</span>
                      <span style="color:teal">{{projectData.Operator}}</span>
                      <span style="color:teal">{{projectData.Value}}</span>
                    </el-form-item>
                  </el-form>
                </div>
              </div>
            </div>
            <div v-for="(oper,i) in projectData.Operation">
              <div style="margin-left:10px;font-size:15px">
                <i class="el-icon-menu" style="color:#13CE66"></i> <span style="margin-top:-20px;margin-left:8px;color:#13CE66">自愈处理 {{i+1}}</span>
              </div>
              <div class="login-bodya">
                <div class="loginWarpa">
                  <div class="login-forma">
                    <el-form ref="form"  label-width="0">
                      <el-form-item prop="old_password" class="login-itema" style="margin-top:-20px;margin-left:-25px">
                        <label >步骤名称 ：</label> <span style="color:teal">{{oper.Title}}</span>
                      </el-form-item>

                      <el-form-item prop="old_password" class="login-itema" style="margin-top:-20px;margin-left:-25px">
                        <label >用户 ：</label> <span style="color:teal">{{oper.User}}</span>
                      </el-form-item>

                      <el-form-item prop="old_password" class="login-itema" style="margin-top:-20px;margin-left:-25px">
                        <label >机器列表 ：</label> <span style="color:teal"> {{oper.Hosts}} </span>
                      </el-form-item>

                      <el-form-item prop="old_password" class="login-itema" style="margin-top:-20px;margin-left:-25px">
                        <label >脚本命令 ：</label> <span style="color:teal"> {{oper.Command}} </span>
                      </el-form-item>
                    </el-form>
                  </div>
                </div>
              </div>
            </div>

            <div style="margin-left:10px;font-size:15px">
              <i class="el-icon-setting" style="color:#FF4949"></i><span style="margin-top:-20px;margin-left:8px;color:#FF4949">自愈通知</span>
            </div>
            <div class="login-bodya">
              <div class="loginWarpa">
                <div class="login-forma">
                  <el-form ref="form"  label-width="0">
                    <el-form-item prop="old_password" class="login-itema" style="margin-top:-20px;margin-left:-25px">
                      <label >超时时间 ：</label><span style="color:teal">{{projectData.Timeout}} 分</span>
                    </el-form-item>

                    <el-form-item prop="old_password" class="login-itema" style="margin-top:-20px;margin-left:-25px">
                      <label >通知方式 ：</label>
                      <span style="color:teal" v-if="projectData.BeginNotice==1">开始</span>
                      <span style="color:teal" v-if="projectData.SuccNotice==1">成功</span>
                      <span style="color:teal" v-if="projectData.FailNotice==1">失败</span>
                    </el-form-item>

                    <el-form-item prop="old_password" class="login-itema" style="margin-top:-20px;margin-left:-25px">
                      <label >通知人员 ：</label>
                      <span style="color:teal" v-if="projectData.NoticeTeam!=''">接收组： {{projectData.NoticeTeam}} </span>
                      <span style="color:teal" v-if="projectData.NoticeUser!=''">接收人： {{projectData.NoticeUser}} </span>
                    </el-form-item>
                  </el-form>
                </div>
              </div>
            </div>
          </el-popover>
          <el-button type="primary" v-popover:popover4 size="small" icon="el-icon-search" @click="open(scope.row.Id)">查看
          </el-button>
          <router-link :to="{ path:'detection/'+scope.row.Id}" tag="span">
            <el-button type="success" size="small" icon="el-icon-setting">检测</el-button>
          </router-link>
          <router-link :to="{ path:'edit/'+scope.row.Id}"  tag="span">
            <el-button type="primary" size="small" icon="el-icon-edit">修改</el-button>
          </router-link>
          <el-button type="danger" size="small" icon="el-icon-delete" @click="delete_data(scope.row.Id)">删除
          </el-button>
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
  import PanelTitle from '@/components/PanelTitle'
  import { getList } from '@/api/autofill'
  import { fetchEditAutoFill, delAutoFill } from '@/api/autofill_access'
  export default {
    components: { PanelTitle },
    data() {
      return {
        tableKey: 0,
        listLoading: true,
        listQuery: {
          page: 1,
          limit: 20,
          sort: '-Id',
          query: ''
        },
        QueryItems: {
          Query: ''
        },
        list: null,
        total: null,
        // 项目详情
        projectData: []
      }
    },
    created() {
      this.getList()
    },
    methods: {
      getList() {
        this.listLoading = true
        getList(this.listQuery).then(response => {
          this.list = response.data.list
          this.total = response.data.total
          this.listLoading = false
        })
      },
      open(id) {
        fetchEditAutoFill(id).then(response => {
          console.log(response.data)
          this.projectData = response.data
        }).catch(err => {
          console.log(err)
        })
      },
      handleSizeChange(val) {
        this.listQuery.limit = val
        this.getList()
      },
      handleCurrentChange(val) {
        this.listQuery.page = val
        this.getList()
      },
      handleFilter() {
        this.listQuery.query = JSON.stringify(this.QueryItems)
        this.getList()
      },
      // 刷新
      on_refresh() {
        this.getList()
      },
      delete_data(id) {
        this.$confirm('此操作将删除该数据, 是否继续?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          delAutoFill(id).then(response => {
            this.listLoading = true
            if (response.data === 'OK') {
              this.$notify({
                title: '成功',
                message: '删除成功',
                type: 'success',
                duration: 2000
              })
              this.getList()
            } else {
              this.listLoading = false
              this.$notify({
                title: '失败',
                message: response.data,
                type: 'error',
                duration: 2000
              })
            }
          })
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除'
          })
        })
      }
    }
  }

</script>

<style>
  .el-dialog {
    width: 30%
  }
</style>