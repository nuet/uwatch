<template>
  <div class="app-container calendar-list-container">
    <div class="filter-container">
    <el-alert
      title=""
      type="success"
      :closable="false">
      <div style="float:left;width:100%;color:#666;">
        <h2>获取Graph数据参数配置</h2>endpoint-主机IP,counter-告警类型
      </div>
    </el-alert>
      <div style="margin-top:20px;margin-bottom:10px;">
        <el-input @keyup.enter.native="handleFilter" style="width: 200px;" class="filter-item" placeholder="主机IP" v-model="QueryItems.Query"></el-input>
        <el-button class="filter-item" type="primary" @click="handleFilter">搜索</el-button>
        <el-button class="filter-item" style="margin-left: 10px;" @click="handleCreate" type="primary" icon="el-icon-edit">添加</el-button>
      </div>
    </div>
    <el-table :key='tableKey' :data="list" v-loading="listLoading" element-loading-text="给我一点时间" border fit highlight-current-row
              style="width: 100%">
      <el-table-column align="center" label="序号" width="65">
        <template slot-scope="scope">
          <span>{{scope.row.Id}}</span>
        </template>
      </el-table-column>
      <el-table-column width="200px" align="center" label="名称">
        <template slot-scope="scope">
          <span>{{scope.row.Title}}</span>
        </template>
      </el-table-column>
      <el-table-column width="200px" align="center" label="endpoint">
        <template slot-scope="scope">
          <span>{{scope.row.Endpoint}}</span>
        </template>
      </el-table-column>
      <el-table-column width="110px" align="center" label="counter">
        <template slot-scope="scope">
          <span>{{scope.row.Counter}}</span>
        </template>
      </el-table-column>
      <el-table-column width="150px" align="center" label="中文名称">
        <template slot-scope="scope">
          <span>{{scope.row.Counter_cn}}</span>
        </template>
      </el-table-column>
      <el-table-column width="80px" align="center" label="最小值">
        <template slot-scope="scope">
          <span>{{scope.row.Min}}</span>
        </template>
      </el-table-column>
      <el-table-column width="80px" align="center" label="最大值">
        <template slot-scope="scope">
          <span>{{scope.row.Max}}</span>
        </template>
      </el-table-column>
      <el-table-column width="100px" align="center" label="临界值">
        <template slot-scope="scope">
          <span>{{scope.row.Avg}}</span>
        </template>
      </el-table-column>
      <el-table-column  width="200px" align="center" label="添加时间">
        <template slot-scope="scope">
          <span>{{scope.row.InsertTime.substring(0,10)}}</span>
        </template>
      </el-table-column>
      <el-table-column  width="200px" align="center" label="更新时间">
        <template slot-scope="scope">
          <span>{{scope.row.UpdateTime.substring(0,10)}}</span>
        </template>
      </el-table-column>
      <el-table-column   align="center" label="是否开启">
        <template slot-scope="scope">
          <el-button type="success" size="mini" @click="updateStatus(scope.row,scope.row.Status,scope.row.Endpoint)" v-if="scope.row.Status">关闭</el-button>
          <el-button type="info" size="mini" @click="updateStatus(scope.row,scope.row.Status,scope.row.Endpoint)" v-else>开启</el-button>
        </template>
      </el-table-column>
      <el-table-column align="center" label="操作" width="230" class-name="small-padding fixed-width">
        <template slot-scope="scope">
            <el-button type="primary" size="mini" @click="handleUpdate(scope.row)">编辑</el-button>
            <el-button v-if="scope.row.status!='deleted'" size="mini" type="danger" @click="delData(scope.row)">删除
            </el-button>
        </template>
      </el-table-column>
    </el-table>
    <div class="pagination-container">
      <el-pagination background @size-change="handleSizeChange" @current-change="handleCurrentChange" :current-page="listQuery.page" :page-sizes="[10,20,30, 50]" :page-size="listQuery.limit" layout="total, sizes, prev, pager, next, jumper" :total="total">
      </el-pagination>
    </div>

    <el-dialog :title="this.autoStatus" :visible.sync="updatedialogFormVisible">
      <template slot-scope="scope">
        <p style="text-align: center;font-size:24px;margin-bottom: 45px;">确认执行当前操作？</p>
        <div slot="footer" class="dialog-footer" style="text-align: center;">
          <el-button @click="cancleStatus()">取 消</el-button>
          <el-button type="primary" @click="handleUpdateStatus()">确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible">
      <el-form  ref="dataForm" :model="temp" label-position="left" label-width="95px" style='width: 400px; margin-left:50px;'>
        <el-form-item label="名称" prop="title">
          <el-input v-model="temp.Title"></el-input>
        </el-form-item>
        <el-form-item label="主机IP" prop="endpoint">
          <el-input v-model="temp.Endpoint"></el-input>
        </el-form-item>
        <el-form-item label="counter" prop="counter">
          <el-input v-model="temp.Counter"></el-input>
        </el-form-item>
        <el-form-item label="counter描述" prop="counter_cn">
          <el-input v-model="temp.Counter_cn"></el-input>
        </el-form-item>
        <!--<el-form-item label="最小值" prop="min">-->
          <!--<el-input v-model="temp.Min"></el-input>-->
        <!--</el-form-item>-->
        <!--<el-form-item label="最大值" prop="max">-->
          <!--<el-input v-model="temp.Max"></el-input>-->
        <!--</el-form-item>-->
        <el-form-item label="临界值" prop="avg">
          <el-input v-model="temp.Avg"></el-input>
        </el-form-item>
        <el-form-item label="Y轴单位" prop="axislabel">
          <el-input v-model="temp.Axislabel"></el-input>
        </el-form-item>
        <el-form-item label="被除数" prop="dividend">
          <el-input v-model="temp.Dividend"></el-input>
        </el-form-item>
        <el-form-item label="开启状态" prop="status">
          <el-select class="filter-item" v-model="temp.Status" placeholder="请选择">
            <el-option  label="开启" :value="1"></el-option>
            <el-option  label="关闭" :value="0"></el-option>
          </el-select>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">关闭</el-button>
        <el-button v-if="dialogStatus=='create'" type="primary" @click="createData">确认</el-button>
        <el-button v-else type="primary" @click="updateData">更新</el-button>
      </div>
    </el-dialog>

    <el-dialog :visible.sync="deldialogFormVisible">
      <template slot-scope="scope">
        <p style="text-align: center;font-size:24px;">确认要删除吗？</p>
        <div slot="footer" class="dialog-footer" style="text-align: center;">
          <el-button @click="deldialogFormVisible = false">取 消</el-button>
          <el-button type="primary" @click="handleDelete()">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>
<style>
  .el-dialog {
    width: 32%
  }
</style>
<script>
  import { getList, UpdateAutoStatus, updateGraph, createGraph, deleteGraph } from '@/api/configure'
  export default {
    data() {
      return {
        tableKey: 0,
        listLoading: true,
        deldialogFormVisible: false,
        updatedialogFormVisible: false,
        listQuery: {
          page: 1,
          limit: 10,
          sort: '-Id',
          query: ''
        },
        QueryItems: {
          Query: ''
        },
        autoStatusForm: {
          Id: '',
          Status: ''
        },
        list: null,
        autoStatus: '',
        total: null,
        dialogStatus: '',
        dialogFormVisible: false,
        textMap: {
          update: '修改',
          create: '创建'
        },
        temp: {
          title: '',
          counter_cn: '',
          min: '',
          max: '',
          avg: '',
          endpoint: '',
          counter: '',
          status: 0,
          axislabel: '',
          dividend: 1
        }
      }
    },
    created() {
      this.getList()
    },
    methods: {
      getList() {
        this.listLoading = true
        console.log('listQuery===>', this.listQuery)
        getList(this.listQuery).then(response => {
          this.list = response.data.list
          this.total = response.data.total
          this.listLoading = false
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
      resetTemp() {
        this.temp = {
          title: '',
          counter_cn: '',
          min: '',
          max: '',
          avg: '',
          endpoint: '',
          counter: '',
          status: 0,
          dividend: 1
        }
      },
      handleCreate() {
        this.resetTemp()
        this.dialogStatus = 'create'
        this.dialogFormVisible = true
        this.$nextTick(() => {
          this.$refs['dataForm'].clearValidate()
        })
      },
      handleUpdate(row) {
        this.resetTemp()
        this.temp = Object.assign({}, row)
        console.log('this.temp', this.temp)
        this.dialogStatus = 'update'
        this.dialogFormVisible = true
      },
      updateData(row) {
        this.$refs['dataForm'].validate((valid) => {
          if (valid) {
            const tempData = Object.assign({}, this.temp)
            tempData.endTime = +new Date(tempData.timestamp) // change Thu Nov 30 2017 16:41:05 GMT+0800 (CST) to 1512031311464
            console.log(tempData)
            updateGraph(tempData).then(() => {
              this.dialogFormVisible = false
              this.$notify({
                title: '成功',
                message: '更新成功',
                type: 'success',
                duration: 2000
              })
              this.getList()
            })
          }
        })
      },
      createData() {
        console.log(this.temp)
        if (this.temp.endpoint == null && this.temp.counter == null) {
          this.$notify({
            title: '提交错误',
            message: '请填写完成字段',
            type: 'warning'
          })
        } else {
          createGraph(this.temp).then(response => {
            if (response.data === 'OK') {
              this.dialogFormVisible = false
              this.total = this.total + 1
              this.getList()
              this.$notify({
                title: '成功',
                message: '添加成功',
                type: 'success',
                duration: 2000
              })
            } else {
              this.$notify({
                title: '失败',
                message: '添加失败',
                type: 'warning',
                duration: 2000
              })
            }
          }).catch(err => {
            console.log(1, err)
            this.$notify({
              title: '失败',
              message: '添加失败',
              type: 'warning',
              duration: 2000
            })
          })
        }
      },
      updateStatus(row, status, title) {
        this.autoStatus = status ? '开启' + '>' + title : '关闭' + '>' + title
        this.autoStatusForm = row
        this.autoStatusForm.Status = status ? 0 : 1
        this.updatedialogFormVisible = true
      },
      handleUpdateStatus() {
        UpdateAutoStatus(this.autoStatusForm).then(response => {
          if (response.data === 'OK') {
            this.updatedialogFormVisible = false
            this.$notify({
              title: '成功',
              message: '更新成功',
              type: 'success',
              duration: 2000
            })
          } else {
            this.$notify({
              title: '失败',
              message: '更新失败',
              type: 'warning',
              duration: 2000
            })
          }
        }).catch(err => {
          console.log(err)
          this.$notify({
            title: '失败',
            message: '更新失败',
            type: 'warning',
            duration: 2000
          })
        })
        this.deldialogFormVisible = false
      },
      cancle() {
        this.deldialogFormVisible = false
        this.getList()
      },
      delData(row) {
        this.deldialogFormVisible = true
        this.delFormData = row
      },
      handleDelete() {
        deleteGraph(this.delFormData.Id).then(response => {
          if (response.data === 'OK') {
            this.$notify({
              title: '成功',
              message: '删除成功',
              type: 'success',
              duration: 2000
            })
            const index = this.list.indexOf(this.delFormData)
            this.total = this.total - 1
            this.list.splice(index, 1)
            this.deldialogFormVisible = false
          } else {
            this.$notify({
              title: '失败',
              message: '删除失败',
              type: 'warning',
              duration: 2000
            })
          }
        }).catch(err => {
          this.fetchSuccess = false
          console.log(err)
          this.$notify({
            title: '失败',
            message: '删除失败',
            type: 'warning',
            duration: 2000
          })
        })
        this.deldialogFormVisible = false
      },
      handleFilter() {
        this.listQuery.query = JSON.stringify(this.QueryItems)
        this.getList()
      },
      cancleStatus() {
        this.updatedialogFormVisible = false
        this.getList()
      }
    }
  }

</script>