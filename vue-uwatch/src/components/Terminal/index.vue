<template>
  <div>
    <el-row v-show="isShowStep">
      <el-col :span="24">
        <el-steps :active="active">
          <el-step
            :title="'步骤'+String(index+1)"
            :description="record.title"
            v-for="(record,index) in records"
            :key="index">
          </el-step>
        </el-steps>
      </el-col>
    </el-row>
    <div v-show="showText.length" style="margin: 5px 5px 0px;
                      padding: 3px;
                      border: 1px dashed rgb(0, 160, 198);
                      background-color: rgb(0,0,0);">
      <code style="background-color: rgb(0, 0, 0);color:#00ff00">
        <br>
        <span v-for="n in showText" :style="{'color': n.color}">  <pre  style="background-color: rgb(0, 0, 0);padding: 0px;font-size:16px;border: 0px;margin:0px " :style="{'color': n.color}" v-html="n.text"></pre> <br></span>
        <br>
      </code>
    </div>
  </div>

</template>
<script type="text/javascript">
  import { getRecord } from '@/api/uwTask'
  export default{
    props: ['taskId', 'isJson'],
    data() {
      return {
        records: [],
        showText: [],
        taskid: this.taskId * 1,
        active: 0,
        time: (Date.parse(new Date()) / 1000) - 10
      }
    },
    computed: {
      test1: function() {
        this.get_data()
      },
      isShowStep: function() {
        if (this.taskId * 1 > 0) {
          return true
        } else {
          return false
        }
      }
    },
    created() {
      this.get_data()
      var self = this
      this.intervalid1 = setInterval(function() {
        self.get_data()
      }, 2000)
    },
    beforeDestroy() {
      this.time = (Date.parse(new Date()) / 1000)
      clearInterval(this.intervalid1)
    },
    methods: {
      get_data() {
        var params = {
          taskId: this.taskId,
          time: this.time
        }
        getRecord(params).then(response => {
          this.records = response.data
          var data = this.records
          this.active = 0
          this.showText = []
          for (var i = 0; i < data.length; i++) {
            var color = '#00ff00'
            if (data[i].status === '0') {
              color = 'red'
            }
            this.showText.push({ text: data[i].command, 'color': color })
            try {
              var text = JSON.parse(data[i].memo)
            } catch (e) {
              console.log(e)
              continue
            }
            if (typeof text === 'string') {
              this.showText.push({ text: data[i].memo, 'color': color })
            } else if (Object.prototype.toString.call(text) === '[object Array]') {
              for (var j = 0; j < text.length; j++) {
                try {
                  this.showText.push({ text: 'IP:' + text[j].Host + '    执行命令:' + data[i].command, 'color': color })
                  if (text[j].ErrorInfo) {
                    this.showText.push({ text: '错误结果:\n' + text[j].Result, 'color': color })
                  } else {
                    this.showText.push({ text: '执行结果:\n' + text[j].Result, 'color': color })
                  }
                  if (text[j].ErrorInfo) {
                    this.showText.push({ text: '错误:' + text[j].ErrorInfo, 'color': color })
                  }
                  this.showText.push({ text: '=============', 'color': color })
                } catch (e) {
                  console.log(e)
                }
              }
            } else {
              this.showText.push({ text: '执行结果:\n' + text.Result, 'color': color })
              if (text.ErrorInfo) {
                this.showText.push({ text: '错误:' + text.ErrorInfo, 'color': color })
              }
              this.showText.push({ text: '=============', 'color': color })
            }
            this.active = this.active + 1
            if (data[i].is_final === '1') {
              setTimeout(clearInterval(this.intervalid1), 5000)
              this.showText.push({ text: '执行完成', 'color': color })
            }
          }
        })
      }
    }
  }
</script>
