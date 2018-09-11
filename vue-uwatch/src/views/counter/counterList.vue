<template>
  <div class="app-container calendar-list-container">
    <div class="filter-container">
      <el-input @keyup.enter.native="handleFilter" style="width: 200px;" class="filter-item" placeholder="标题" v-model="QueryItems.Query">
      </el-input>
      <el-button class="filter-item" type="primary" v-waves icon="el-icon-search" @click="handleFilter">搜索</el-button>
      <el-button class="filter-item" style="margin-left: 10px;" @click="handleCreate" type="primary" icon="el-icon-edit">添加</el-button>
      <el-button class="filter-item" type="primary" :loading="downloadLoading" v-waves icon="el-icon-download" @click="handleDownload">导出</el-button>
    </div>

    <el-table :key='tableKey' :data="list" v-loading="listLoading" element-loading-text="给我一点时间" border fit highlight-current-row
              style="width: 100%">
      <el-table-column align="center" label="序号" width="65">
        <template slot-scope="scope">
          <span>{{scope.row.Id}}</span>
        </template>
      </el-table-column>
      <el-table-column min-width="100px" label="Counter">
        <template slot-scope="scope">
          <span>{{scope.row.Counter}}</span>
        </template>
      </el-table-column>
      <el-table-column width="110px" align="center" label="Counter描述">
        <template slot-scope="scope">
          <span>{{scope.row.Counter_cn}}</span>
        </template>
      </el-table-column>
      <el-table-column width="110px" align="center" label="纵轴单位">
        <template slot-scope="scope">
          <span>{{scope.row.Axislabel}}</span>
        </template>
      </el-table-column>
      <el-table-column width="110px" align="center" label="平均线">
        <template slot-scope="scope">
          <span>{{scope.row.Avg}}</span>
        </template>
      </el-table-column>
      <el-table-column width="110px" align="center" label="换算基数">
        <template slot-scope="scope">
          <span>{{scope.row.Dividend}}</span>
        </template>
      </el-table-column>
      <el-table-column   align="center" label="是否开启">
        <template slot-scope="scope">
          <el-button type="success" size="mini" @click="updateStatus(scope.row,scope.row.Status,scope.row.Counter)" v-if="scope.row.Status">关闭</el-button>
          <el-button type="info" size="mini" @click="updateStatus(scope.row,scope.row.Status,scope.row.Counter)" v-else>开启</el-button>
        </template>
      </el-table-column>
      <el-table-column width="200px" align="center" label="创建日期">
        <template slot-scope="scope">
          <span>{{scope.row.Createtime.substring(0,10)}}</span>
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

    <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible">
      <el-form  ref="dataForm" :model="temp" label-position="left" label-width="90px" style='width: 400px; margin-left:50px;'>
        <el-form-item label="Counter" prop="Counter">
          <el-input v-model="temp.Counter"></el-input>
        </el-form-item>
        <el-form-item label="Counter描述" prop="Counter_cn">
          <el-input v-model="temp.Counter_cn"></el-input>
        </el-form-item>
        <el-form-item label="平均线" prop="Avg">
          <el-input v-model="temp.Avg"></el-input>
        </el-form-item>
        <el-form-item label="纵轴单位" prop="Axislabel">
          <el-input v-model="temp.Axislabel"></el-input>
        </el-form-item>
        <el-form-item label="换算基数" prop="Dividend">
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

    <el-dialog :title="this.autoStatus" :visible.sync="updatedialogFormVisible">
      <template slot-scope="scope">
        <p style="text-align: center;font-size:24px;margin-bottom: 45px;">确认执行当前操作？</p>
        <div slot="footer" class="dialog-footer" style="text-align: center;">
          <el-button @click="updatedialogFormVisible = false">取 消</el-button>
          <el-button type="primary" @click="handleUpdateStatus()">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
  import { getList, createCounter, updateCounter, deleteCounter } from '@/api/counter'
  import waves from '@/directive/waves' // 水波纹指令
  import { parseTime } from '@/utils'

  export default {
    name: 'complexTable',
    directives: {
      waves
    },
    data() {
      return {
        tableKey: 0,
        list: null,
        delFormData: '',
        total: null,
        listLoading: true,
        listQuery: {
          page: 1,
          limit: 10,
          sort: '-Id',
          query: ''
        },
        QueryItems: {
          Query: ''
        },
        calendarTypeOptions: '',
        temp: {
          Counter: '',
          Status: 1,
          Counter_cn: '',
          Axislabel: '',
          Dividend: '',
          Avg: '',
          Createtime: null,
          Updatetime: null
        },
        dialogFormVisible: false,
        deldialogFormVisible: false,
        dialogStatus: '',
        textMap: {
          update: '修改',
          create: '创建'
        },
        autoStatusForm: {
          Id: '',
          Status: ''
        },
        downloadLoading: false,
        updatedialogFormVisible: false,
        autoStatus: ''
      }
    },
    filters: {
      statusFilter(status) {
        const statusMap = {
          published: 'success',
          draft: 'info',
          deleted: 'danger'
        }
        return statusMap[status]
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
      handleFilter() {
        this.listQuery.query = JSON.stringify(this.QueryItems)
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
      resetTemp() {
        this.temp = {
          Counter: '',
          Status: 1,
          Counter_cn: '',
          Createtime: null,
          Updatetime: null
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
        this.temp = Object.assign({}, row)
        console.log(this.temp)
        this.dialogStatus = 'update'
        this.dialogFormVisible = true
      },

      createData() {
        console.log(this.temp)
        if (this.temp.Counter_cn == null || this.temp.Counter == null) {
          this.$notify({
            title: '提交错误',
            message: '请填写完成字段',
            type: 'warning'
          })
        } else {
          createCounter(this.temp).then(response => {
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
      updateData(row) {
        console.log('===>', this.temp)
        this.$refs['dataForm'].validate((valid) => {
          if (valid) {
            const tempData = Object.assign({}, this.temp)
            console.log(tempData)
            updateCounter(tempData).then(() => {
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
      updateStatus(row, status, title) {
        this.autoStatus = status ? '关闭' + '>' + title : '开启' + '>' + title
        this.autoStatusForm = row
        this.autoStatusForm.Status = status ? 0 : 1
        this.updatedialogFormVisible = true
      },
      handleUpdateStatus() {
        updateCounter(this.autoStatusForm).then(response => {
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
        deleteCounter(this.delFormData.Id).then(response => {
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
      handleDownload() {
        this.downloadLoading = true
        import('@/vendor/Export2Excel').then(excel => {
          const tHeader = ['timestamp', 'title', 'type', 'importance', 'status']
          const filterVal = ['timestamp', 'title', 'type', 'importance', 'status']
          const data = this.formatJson(filterVal, this.list)
          excel.export_json_to_excel(tHeader, data, 'table-list')
          this.downloadLoading = false
        })
      },
      formatJson(filterVal, jsonData) {
        return jsonData.map(v => filterVal.map(j => {
          if (j === 'timestamp') {
            return parseTime(v[j])
          } else {
            return v[j]
          }
        }))
      }
    }
  }
</script>
