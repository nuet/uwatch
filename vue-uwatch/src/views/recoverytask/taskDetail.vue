<template>
  <div class="panel">
    <div class="panel-body"
         element-loading-text="拼命加载中">
      <el-form label-width="100px">
        <el-row>
          <el-col :span="12">
            <el-form-item label="告警信息:">
              {{task.Source}}
            </el-form-item>
            <el-form-item label="自愈项目:">
              {{task.Item}}
            </el-form-item>
            <el-form-item label="自愈时间:">
              {{task.Duration}}
            </el-form-item>
          </el-col>

          <el-col :span="12">
            <el-form-item label="任务状态:">
              <span v-if="task.Status == 0">未完成</span>
              <span v-else>已完成</span>
            </el-form-item>
            <el-form-item label="自愈结果:">
              {{task.Result}}
            </el-form-item>
            <el-form-item label="产生时间:">
              {{task.CeateTime}}
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>


      <terminal :taskId="task_id"></terminal>

    </div>
  </div>
</template>

<script>
  import { getTask } from '@/api/uwTask'
  import Terminal from '@/components/Terminal'
  export default {
    name: 'taskssDetail',
    data() {
      return {
        task: [],
        task_id: this.$route.params.id
      }
    },
    methods: {
      getList() {
        getTask(this.task_id).then(response => {
          this.task = response.data.task
        })
      }
    },
    components: { Terminal }
  }
</script>
