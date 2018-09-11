/**
 * Created by Administrator on 2018-03-23.
 */
import PanelTitle from '@/components/PanelTitle'
import MDinput from '@/components/MDinput'
import Multiselect from 'vue-multiselect'// 使用的一个多选框组件，element-ui的select不能满足所有需求
import 'vue-multiselect/dist/vue-multiselect.min.css'// 多选框组件css
import Sticky from '@/components/Sticky' // 粘性header组件
import { validateURL } from '@/utils/validate'
import Vue from 'vue'
import { postAutoFill, putAutoFill, fetchEditAutoFill } from '@/api/autofill_access'

const defaultForm = {
  Title: '',
  Metric: '',
  Tag: '',
  Operator: '',
  Value: '',
  Author: '',
  Timeout: 40,
  BeginNotice: true,
  SuccNotice: true,
  FailNotice: true,
  NoticeUser: [],
  NoticeTeam: '',
  Isvalid: 0,
  Operation: [ {'Title':'', 'User':'root', 'Hosts':'', 'Command':'', 'Step':1, 'Expand': false} ],
  OperationDel: []
}

export default {
  name: 'autofill_access',
  components: { MDinput, Multiselect, Sticky, PanelTitle },
  data() {
    const validateValueRequire = (rule, value, callback) => {
      if (value) {
        var regPos = /^\d+(\.\d+)?$/; //非负浮点数
        var regNeg = /^(-(([0-9]+\.[0-9]*[1-9][0-9]*)|([0-9]*[1-9][0-9]*\.[0-9]+)|([0-9]*[1-9][0-9]*)))$/; //负浮点数
        if(regPos.test(value) || regNeg.test(value)){
          callback();
        }else{
          callback(new Error('请输入数字值'));
        }
      } else {
        return callback(new Error('该项不能为空'));
      }
    }
    return {
      fill_id: this.$route.params.id,
      optype: this.$route.params.id ? 'published' : 'draft',
      postForm: Object.assign({}, defaultForm),
      fetchSuccess: true,
      loading: false,
      teamLIstOptions: [],
      userLIstOptions: [],
      expands: [],
      steps: 1,
      rules: {
        Title: [{required: true, message: '自愈方案名称不能为空', trigger: 'blur'}],
        Metric: [{required: true, message: 'Metric不能为空', trigger: 'blur'}],
        Operator: [{required: true, message: '比较运算符必选', trigger: 'change'}],
        Value: [{ validator: validateValueRequire, trigger: 'blur', required: true }],
        Timeout: [{ validator: validateValueRequire, trigger: 'blur', required: true }]
      }
    }
  },
  computed: {
    operationLength() {
      return this.postForm.Operation.length
    }
  },
  created() {
    if (this.fill_id) {
      this.fetchData(this.fill_id)
    } else {
      this.postForm = Object.assign({}, defaultForm)
    }
    if (this.$store.state.user.userList.length === 0) {
      this.$store.dispatch('GetUserList');
    }
    if (this.$store.state.user.teamList.length === 0) {
      this.$store.dispatch('GetTeamList');
    }
  },
  methods: {
    fetchData(id) {
      fetchEditAutoFill(id).then(response => {
        this.postForm = response.data
        this.postForm.BeginNotice = this.postForm.BeginNotice ? true:false
        this.postForm.SuccNotice = this.postForm.SuccNotice ? true:false
        this.postForm.FailNotice = this.postForm.FailNotice ? true:false
        this.postForm.NoticeUser = this.postForm.NoticeUser ? this.postForm.NoticeUser.split(",") : []
        this.postForm.OperationDel = []
        this.steps = this.postForm.Operation.length
      }).catch(err => {
        this.fetchSuccess = false
        console.log(err)
      })
    },
    addOperation() {
      this.postForm.Operation.push({'Title':'', 'User':'root', 'Hosts':'', 'Command':'', 'Step':this.steps+1, 'Expand': false})
      this.steps++
    },
    delOperation(row) {
      this.$confirm('是否删除该步骤?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        if (this.operationLength <=1) {
          this.$message({
            message: '至少需要一个步骤',
            type: 'error'
          })
          return
        }
        this.postForm.Operation.splice(this.postForm.Operation.indexOf(row), 1)
        if (row.Id != '') {
          this.postForm.OperationDel.push(row.Id)
        }
      }).catch(() => {
      })
    },
    handleRowExpansion(row,event) {
      this.expands = []
      this.postForm.Operation.forEach((item) => {
        if (item.Step != row.Step) {
          item.Expand = false
        }
      })
      if (event) {
        row.Expand = true
      } else {
        row.Expand = row.Expand?false:true
      }
      if (row.Expand) {
        this.expands.push(row.Step)
      }
    },
    handleRowMove (row,event) {
      if (event == 'up') {
        var index = this.postForm.Operation.indexOf(row)
        if (index == 0) {
          return
        }
        var tempOption = this.postForm.Operation[index - 1]
        Vue.set(this.postForm.Operation, index - 1, this.postForm.Operation[index])
        Vue.set(this.postForm.Operation, index, tempOption)
      } else {
        var index = this.postForm.Operation.indexOf(row)
        if (index == (this.operationLength -1)) {
          return
        }
        var tempOption = this.postForm.Operation[index + 1]
        Vue.set(this.postForm.Operation, index + 1, this.postForm.Operation[index])
        Vue.set(this.postForm.Operation, index, tempOption)
      }
    },
    submitForm() {
      console.log(this.postForm)
      this.$refs.postForm.validate(valid => {
        if (valid) {
          let msg = this.checkOperation()
          if (msg != '') {
            this.$message({
              message: msg,
              type: 'error'
            })
            return
          }
          this.loading = true
          this.postForm.BeginNotice = this.postForm.BeginNotice ? 1:0
          this.postForm.SuccNotice = this.postForm.SuccNotice ? 1:0
          this.postForm.FailNotice = this.postForm.FailNotice ? 1:0
          this.postForm.NoticeUser = this.postForm.NoticeUser.toString()
          this.postForm.Timeout = parseInt(this.postForm.Timeout)
          if (this.fill_id) {
            putAutoFill(this.postForm).then(response => {
              if (response.data === 'OK') {
                this.$notify({
                  title: '成功',
                  message: '修改方案成功',
                  type: 'success',
                  duration: 2000
                })
                this.$router.push({path: '/autofill/access/list'})
              } else {
                this.loading = false
                this.$notify({
                  title: '失败',
                  message: response.data,
                  type: 'error',
                  duration: 2000
                })
              }
            })
          } else {
            this.postForm.Author = this.$store.state.user.userInfo.name
            postAutoFill(this.postForm).then(response => {
              if (response.data === 'OK') {
                this.$notify({
                  title: '成功',
                  message: '新增方案成功',
                  type: 'success',
                  duration: 2000
                })
                this.$router.push({path: '/autofill/access/list'})
              } else {
                this.loading = false
                this.$notify({
                  title: '失败',
                  message: response.data,
                  type: 'error',
                  duration: 2000
                })
              }
            })
          }
        } else {
          return false
        }
      })
    },
    checkOperation() {
      for (var i=0; i<this.postForm.Operation.length; i++) {
        let per = this.postForm.Operation[i]
        if (per.Title == '') {
          return '步骤'+(i+1)+'--步骤名不能为空'
        }
        if (per.User == '') {
          return '步骤'+(i+1)+'--执行账户不能为空'
        }
        if (per.Hosts == '') {
          return '步骤'+(i+1)+'--机器列表不能为空'
        }
      }
      if (this.postForm.NoticeTeam == '' && this.postForm.NoticeUser.length == 0) {
        return '至少需要一个接受组或接收人'
      }
      return ''
    },
    getRemoteUserList(query) {
      let list = [];
      this.$store.state.user.userList.forEach((row) => {
        list.push(row.name)
      })
      if (query !== '') {
        setTimeout(() => {
          this.userLIstOptions = list.filter(item => {
            return item.toLowerCase().indexOf(query.toLowerCase()) > -1;
          })
        }, 200)
      } else {
        this.userLIstOptions = []
      }
    },
    getRemoteTeamList(query) {
      let list = [];
      this.$store.state.user.teamList.forEach((row) => {
        list.push(row.name)
      })
      if (query !== '') {
        setTimeout(() => {
          this.teamLIstOptions = list.filter(item => {
            return item.toLowerCase().indexOf(query.toLowerCase()) > -1;
          })
        }, 200)
      } else {
        this.teamLIstOptions = []
      }
    },
    back() {
      this.$router.go(-1)
    }
  }
}