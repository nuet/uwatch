<template>
  <div class="panel">
    <panel-title :title="$route.meta.title"></panel-title>
    <div class="panel-body"
         v-loading="load_data"
         element-loading-text="拼命加载中">
      <terminal :taskId="-1"></terminal>
    </div>
  </div>
</template>
<script type="text/javascript">
  import PanelTitle from '@/components/PanelTitle'
  import Terminal from '@/components/Terminal'
  import { detectionAutoFill } from '@/api/autofill_access'
  export default{
    data() {
      return {
        route_id: this.$route.params.id,
        load_data: false
      }
    },
    created() {
      if (this.route_id) {
        this.detection_data()
      } else {
        this.$message({
          message: '自愈方案不存在',
          type: 'warning'
        })
        setTimeout(() => {
          this.$router.push({
            name: 'autofill-access-list'
          })
        }, 500)
      }
    },
    methods: {
      detection_data() {
        detectionAutoFill(this.route_id).then(response => {
          if (response.data === 'OK') {
            this.$message({
              message: '检测成功',
              type: 'success'
            })
          }
        }).catch(err => {
          console.log(err)
        })
      }
    },
    components: { PanelTitle, Terminal }
  }
</script>
