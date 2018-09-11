<template>
  <div class="createPost-container">
    <el-form class="form-container" :model="postForm" :rules="rules" ref="postForm">
      <sticky :className="'sub-navbar '+optype">
        <template v-if="fetchSuccess">
          <el-dropdown trigger="click">
            <el-button plain>{{postForm.Isvalid?'已启用':'未启用'}}
              <i class="el-icon-caret-bottom el-icon--right"></i>
            </el-button>
            <el-dropdown-menu class="no-padding" slot="dropdown">
              <el-dropdown-item>
                <el-radio-group style="padding: 10px;" v-model="postForm.Isvalid">
                  <el-radio :label="1">启用</el-radio>
                  <el-radio :label="0">停用</el-radio>
                </el-radio-group>
              </el-dropdown-item>
            </el-dropdown-menu>
          </el-dropdown>
          <el-button v-loading="loading" style="margin-left: 10px;" type="success" @click="submitForm()">保存
          </el-button>
          <el-button v-loading="loading" type="warning" @click="back()">取消</el-button>
        </template>
        <template v-else>
          <el-tag>发送异常错误,刷新页面,或者联系程序员</el-tag>
        </template>
      </sticky>
      <div class="createPost-main-container">
        <el-form-item  prop="Title">
          <MDinput name="name" v-model="postForm.Title" required :maxlength="100">
            自愈方案名称
          </MDinput>
        </el-form-item>
        <div class="panel" style="margin-top: 20px;">
          <panel-title title="自愈场景"></panel-title>
          <el-row style="margin-top: 20px;">
            <el-col :span="21">
              <div class="postInfo-container">
                <el-row>
                  <el-col :span="8">
                    <el-tooltip class="item" effect="dark" content="Falcon Metric" placement="top">
                      <el-form-item label-width="80px" label="Metric" prop="Metric">
                        <el-input placeholder="Falcon Metric" v-model="postForm.Metric">
                        </el-input>
                      </el-form-item>
                    </el-tooltip>
                  </el-col>

                  <el-col :span="8">
                    <el-tooltip class="item" effect="dark" content="Falcon Tags" placement="top">
                      <el-form-item label-width="80px" label="Tags">
                        <el-input placeholder="Falcon Tags" v-model="postForm.Tag">
                        </el-input>
                      </el-form-item>
                    </el-tooltip>
                  </el-col>

                  <el-col :span="3">
                    <el-tooltip class="item" effect="dark" content="比较运算符" placement="top">
                      <el-form-item label-width="80px" label=" " prop="Operator">
                        <el-select v-model="postForm.Operator"  size="300">
                          <el-option label=">" value=">"></el-option>
                          <el-option label="=" value="="></el-option>
                          <el-option label="<" value="<"></el-option>
                        </el-select>
                      </el-form-item>
                    </el-tooltip>
                  </el-col>

                  <el-col :span="5">
                    <el-tooltip class="item" effect="dark" content="Judge Value" placement="top">
                      <el-form-item label-width="80px" label="Value" prop="Value">
                        <el-input placeholder="Judge Value" v-model="postForm.Value">
                        </el-input>
                      </el-form-item>
                    </el-tooltip>
                  </el-col>

                </el-row>
              </div>
            </el-col>
          </el-row>
        </div>

        <div class="panel" style="margin-top: 20px;">
          <panel-title title="自愈处理"></panel-title>
          <el-row style="margin: 20px;">
            <el-table
              :data="postForm.Operation"
              ref="operation_table"
              row-key="Step"
              :expand-row-keys="expands"
              style="width: 100%;" >
              <el-table-column type="expand">
                <template slot-scope="scope">
                  <el-form label-position="left" inline class="demo-table-expand">
                    <el-form-item label="机器列表:"  required>
                      <el-tooltip class="item" effect="dark" content="要发布的机器列表，一行一个，22端口, 当前告警机器{CURRENT}"
                                  placement="top">
                        <el-input type="textarea" :rows="4" style="width:400px;" v-model="scope.row.Hosts" placeholder="192.168.0.1
192.168.0.2">
                        </el-input>
                      </el-tooltip>
                    </el-form-item>
                    <el-form-item label="脚本命令:" >
                      <el-tooltip class="item" effect="dark" content="一行一条,相关参数{HOSTS}"
                                  placement="top">
                        <el-input type="textarea" :rows="4" style="width:400px;" v-model="scope.row.Command">
                        </el-input>
                      </el-tooltip>
                    </el-form-item>
                  </el-form>
                </template>
              </el-table-column>
              <el-table-column label="执行顺序"  align="center"  width="100px">
                <template slot-scope="scope">
                  <div>{{scope.$index+1}}</div>
                </template>
              </el-table-column>
              <el-table-column  label="步骤名" label-class-name="reqlabel" align="center" >
                <template slot-scope="scope">
                  <el-input @focus="handleRowExpansion(scope.row,true)"  placeholder="" v-model="scope.row.Title">
                  </el-input>
                </template>
              </el-table-column>
              <el-table-column  align="center" label="执行账户" label-class-name="reqlabel">
                <template slot-scope="scope">
                  <el-input placeholder="Falcon Metric" v-model="scope.row.User">
                  </el-input>
                </template>
              </el-table-column>
              <el-table-column  align="center" label="所有操作" >
                <template slot-scope="scope">
                  <el-button  type="text" icon="el-icon-edit-outline" @click="handleRowExpansion(scope.row)" style="color: #409EFF; font-size:18px;" ></el-button>
                  <el-button  type="text" icon="el-icon-sort-up" @click="handleRowMove(scope.row,'up')" style="color: #409EFF; font-size:18px;"></el-button>
                  <el-button  type="text" icon="el-icon-sort-down" @click="handleRowMove(scope.row,'down')" style="color: #409EFF; font-size:18px;"></el-button>
                  <el-button  type="text" icon="el-icon-close" @click="delOperation(scope.row)" style="color: red; font-size:18px;"  ></el-button>
                </template>
              </el-table-column>
            </el-table>
            <h5 style="margin-bottom:5px;font-size: 15px;"><el-button type="text" class="fa fa-plus" @click="addOperation" style="color:#409EFF;font-size: 16px;font-weight: bold ">    添加步骤</el-button></h5>
          </el-row>
        </div>

        <div class="panel" style="margin-top: 20px;">
          <panel-title title="自愈通知"></panel-title>
          <el-row style="margin-top: 20px;">
              <div class="postInfo-container">
                <el-row>
                  <el-col :span="21">
                    <el-form-item label-width="100px" label="超时" prop="Timeout">
                      <el-col :span="3">
                        <el-input  v-model="postForm.Timeout" ><template slot="append" style="padding: 0 5px;">分</template></el-input>
                      </el-col>
                      <el-col :span="4">
                        <span style="padding-left: 10px;">以上按失败处理</span>
                      </el-col>
                    </el-form-item>
                  </el-col>
                  <el-col :span="21">
                    <el-form-item label-width="100px" label="通知方式">
                      <el-checkbox v-model="postForm.BeginNotice">开始时</el-checkbox>
                      <el-checkbox v-model="postForm.SuccNotice">成功时</el-checkbox>
                      <el-checkbox v-model="postForm.FailNotice">失败时</el-checkbox>
                    </el-form-item>
                  </el-col>
                  <el-col :span="21">
                    <el-form-item label-width="100px" label="通知人员">
                      <el-col :span="2">
                        接收组：
                      </el-col>
                      <el-col :span="8">
                        <multiselect v-model="postForm.NoticeTeam" :options="teamLIstOptions" @search-change="getRemoteTeamList" placeholder="搜索用户组" selectLabel="选择"
                                     deselectLabel="删除" :internalSearch="false">
                          <span slot='noResult'>无结果</span>
                        </multiselect>
                      </el-col>
                    </el-form-item>
                  </el-col>
                  <el-col :span="21">
                    <el-form-item label-width="100px" label="">
                      <el-col :span="2">
                        接收人：
                      </el-col>
                      <el-col :span="8">
                        <multiselect v-model="postForm.NoticeUser" :options="userLIstOptions" @search-change="getRemoteUserList" :multiple="true" placeholder="搜索用户" selectLabel="选择"
                                     deselectLabel="删除" :internalSearch="false" >
                          <span slot='noResult'>无结果</span>
                        </multiselect>
                      </el-col>
                    </el-form-item>
                  </el-col>
                  <el-col :span="21">
                    <el-form-item label-width="100px" label="">
                      <span style="color:#ffcc6b">注：</span><span style="color:#ccc">在UIC中管理接收组</span><a href="http://fe.juanpi.org/team/all" target="_blank">【快捷入口】</a>
                    </el-form-item>
                  </el-col>
                </el-row>
              </div>
          </el-row>
        </div>
      </div>
    </el-form>

  </div>
</template>

<script src="../js/add.js"></script>

<style rel="stylesheet/scss" lang="scss" scoped>
  @import "src/styles/mixin.scss";
  .title-prompt{
    position: absolute;
    right: 0px;
    font-size: 12px;
    top:10px;
    color:#ff4949;
  }
  .createPost-container {
    position: relative;
    .createPost-main-container {
      padding: 40px 45px 20px 50px;
      .postInfo-container {
        position: relative;
        @include clearfix;
        margin-bottom: 10px;
        .postInfo-container-item {
          float: left;
        }
      }
      .editor-container {
        min-height: 500px;
        margin: 0 0 30px;
        .editor-upload-btn-container {
          text-align: right;
          margin-right: 10px;
          .editor-upload-btn {
            display: inline-block;
          }
        }
      }
    }
    .word-counter {
      width: 40px;
      position: absolute;
      right: -10px;
      top: 0px;
    }
  }
  .demo-table-expand {
    font-size: 0;
  }
  .demo-table-expand label {
    width: 90px;
    color: #99a9bf;
  }
  .demo-table-expand .el-form-item {
    margin-right: 0;
    margin-bottom: 0;
    width: 50%;
  }
</style>

