<template>
  <div class="box-board">
    <el-card class="box-card-left">
      <template #header>
        <div class="card-header">
          <span>机房及线路</span>
          <span>
            <span style="margin-right: 10px"><el-checkbox v-model="search_ip" label="同时搜索线路地址池" size="small" /></span>
            <span>
            <el-input
              size="small"
              placeholder="请输入内容"
              @keyup.enter="onSearch()"
              v-model="searchContent"
              style="width: 300px"
            >
              <template #append>
                <el-button
                  size="small"
                  @click="onSearch()"
                  icon="el-icon-search"
                ></el-button>
              </template>
            </el-input>
          </span>
          </span>
        </div>
      </template>
      <div height="100%" class="box-card-left-body">
        <div class="box-card-left-body-content">
          <el-scrollbar class="card-scrollbar">
            <el-tree
              :data="data"
              node-key="id"
              highlight-current
              :expand-on-click-node="false"
              @node-click="nodeClick"
            >
              <template #default="{ node, data }">
                <span class="custom-tree-node">
                  <span>
                    <span v-if="data.tree_type === 'idc'"><svg class="icon" width="13" height="13" viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg" data-v-042ca774=""><path fill="currentColor" d="M512 128L128 447.936V896h255.936V640H640v256h255.936V447.936z"></path></svg></span>
                    <span v-if="data.tree_type === 'line'"><svg class="icon" width="13" height="13" viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg" data-v-042ca774=""><path fill="currentColor" d="M679.872 348.8l-301.76 188.608a127.808 127.808 0 015.12 52.16l279.936 104.96a128 128 0 11-22.464 59.904l-279.872-104.96a128 128 0 11-16.64-166.272l301.696-188.608a128 128 0 1133.92 54.272z"></path></svg></span>
                    {{ node.label }}
                  </span>
                  <span v-if="data.tree_type === 'idc'">
                    <el-tag size="small" @click="append(data)"> <svg class="icon" width="13" height="13" viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg" data-v-042ca774=""><path fill="currentColor" d="M832 384H576V128H192v768h640V384zm-26.496-64L640 154.496V320h165.504zM160 64h480l256 256v608a32 32 0 01-32 32H160a32 32 0 01-32-32V96a32 32 0 0132-32zm320 512V448h64v128h128v64H544v128h-64V640H352v-64h128z"></path></svg> </el-tag>
                    <el-tag size="small" @click="edit(node, data)"><svg t="1639990532110" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="12171" xmlns:xlink="http://www.w3.org/1999/xlink" width="13" height="13" data-v-042ca774=""><path d="M199.04 672.64l193.984 112 224-387.968-193.92-112-224 388.032z m-23.872 60.16l32.896 148.288 144.896-45.696-177.792-102.592zM455.04 229.248l193.92 112 56.704-98.112-193.984-112-56.64 98.112zM104.32 708.8l384-665.024 304.768 175.936-383.936 665.088h0.064l-248.448 78.336-56.448-254.336z m384 254.272v-64h448v64h-448z" p-id="12172"></path></svg> </el-tag>
                    <el-tag size="small" type="danger" @click="removeIDC(node, data)"> <svg class="icon" width="13" height="13" viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg" data-v-042ca774=""><path fill="currentColor" d="M352 192V95.936a32 32 0 0132-32h256a32 32 0 0132 32V192h256a32 32 0 110 64H96a32 32 0 010-64h256zm64 0h192v-64H416v64zM192 960a32 32 0 01-32-32V256h704v672a32 32 0 01-32 32H192zm224-192a32 32 0 0032-32V416a32 32 0 00-64 0v320a32 32 0 0032 32zm192 0a32 32 0 0032-32V416a32 32 0 00-64 0v320a32 32 0 0032 32z"></path></svg> </el-tag >
                  </span>
                  <span v-if="data.tree_type === 'line'">
                    <el-tag size="small" @click="editLine(node, data)"><svg t="1639990532110" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="12171" xmlns:xlink="http://www.w3.org/1999/xlink" width="13" height="13" data-v-042ca774=""><path d="M199.04 672.64l193.984 112 224-387.968-193.92-112-224 388.032z m-23.872 60.16l32.896 148.288 144.896-45.696-177.792-102.592zM455.04 229.248l193.92 112 56.704-98.112-193.984-112-56.64 98.112zM104.32 708.8l384-665.024 304.768 175.936-383.936 665.088h0.064l-248.448 78.336-56.448-254.336z m384 254.272v-64h448v64h-448z" p-id="12172"></path></svg></el-tag>
                    <el-tag size="small" type="danger" @click="removeLine(node, data)"> <svg class="icon" width="13" height="13" viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg" data-v-042ca774=""><path fill="currentColor" d="M352 192V95.936a32 32 0 0132-32h256a32 32 0 0132 32V192h256a32 32 0 110 64H96a32 32 0 010-64h256zm64 0h192v-64H416v64zM192 960a32 32 0 01-32-32V256h704v672a32 32 0 01-32 32H192zm224-192a32 32 0 0032-32V416a32 32 0 00-64 0v320a32 32 0 0032 32zm192 0a32 32 0 0032-32V416a32 32 0 00-64 0v320a32 32 0 0032 32z"></path></svg> </el-tag >
                  </span>
                </span>
              </template>
            </el-tree>
          </el-scrollbar>
        </div>
        <div class="box-card-left-body-action">
          <el-button size="small" type="success" icon="el-icon-baseball" class="button" @click="idcAppend">增加机房</el-button>
          <el-button size="small" type="info" plain class="button" @click="flushtree">刷新列表</el-button>
          <el-button v-if="pushing_all===false" icon="el-icon-upload" type="warning" plain size="small" class="button" 
            @click="updateAllIPAddrs">更新所有IP</el-button>
          <el-button v-if="pushing_all===true" icon="el-icon-loading" type="warning" plain size="small" class="button" 
            @click="updateAllIPAddrs">更新所有IP</el-button>
          <el-button v-if="pushing_part===false" icon="el-icon-upload" type="warning" plain size="small" class="button" 
            @click="updatePartIPAddrs">只更新未设置IP</el-button>
          <el-button v-if="pushing_part===true" icon="el-icon-loading"  type="warning" plain size="small" class="button" 
            @click="updatePartIPAddrs">只更新未设置IP</el-button>
          <el-button v-if="pushing_create_label===false" icon="el-icon-upload" type="info" plain size="small" class="button" 
            @click="doCreateLabelForAllIPs">JOB组中生成标签</el-button>
          <el-button v-if="pushing_create_label===true" icon="el-icon-loading" type="info" plain size="small" class="button" 
            @click="doCreateLabelForAllIPs">JOB组中生成标签</el-button>
        </div>
      </div>
    </el-card>
    <el-card class="box-card-right">
      <template #header>
        <div class="card-header">
          <span>
            当前选择机房：<el-tag v-if="currentIDCTitle!==''" size="mini" type="warning">{{currentIDCTitle}}</el-tag>
            线路：<el-tag v-if="currentLineTitle!==''" size="mini" type="warning">{{currentLineTitle}}</el-tag>
          </span>
          <el-button size="small" type="success" class="button" icon="el-icon-upload2" @click="putIDCOrLineIPData"> 更 新 </el-button>
        </div>
      </template>
      <div>
        <el-form
          label-position="top"
          :model="linedetailinfo"
        >
          <el-form-item label="IP列表：（以英文分号(;)隔开；如:192.168.1.0/24;10.10.10.1;172.16.1.1~172.16.2.1）">
            <el-input :disabled="should_disabled" :autosize="{ minRows: 10, maxRows: 20 }" type="textarea" placeholder="" v-model="linedetailinfo.ipaddrs_net_line"></el-input>
          </el-form-item>
          <el-form-item label="备注信息：">
            <el-input :autosize="{ minRows: 10, maxRows: 20 }" type="textarea" placeholder="" v-model="linedetailinfo.remark_info"></el-input>
          </el-form-item>
        </el-form>
      </div>
    </el-card>
    <el-dialog v-model="dialogLineVisible" :title="idcName" width="400px">
      <el-form :model="Lineform" :rules="lineRules" ref="Lineform">
        <el-form-item label="线路名称" :label-width="formLabelWidth" prop="label">
          <el-input size="small" v-model="Lineform.label" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button size="small" @click="dialogLineVisible = false">取消</el-button>
          <el-button size="small" type="primary" @click="lineAppend('Lineform')"
            >确定</el-button
          >
        </span>
      </template>
    </el-dialog>
    <el-dialog v-model="dialogIDCVisible" :title="dialogIDCTitle" width="400px">
      <el-form :model="IDCform" :rules="IDCRules" ref="IDCform">
        <el-form-item label="机房名称" :label-width="formLabelWidth" prop="label">
          <el-input size="small" v-model="IDCform.label" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button size="small" @click="dialogIDCVisible = false">取消</el-button>
          <el-button size="small" type="primary" @click="idcAppendConfirm('IDCform')"
            >确定</el-button
          >
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script>
  let tree_id = 1000;

  // import { getIDC, getIDCs, postIDC, putIDC, deleteIDC } from '@/api/idc.js'
  import { getIDC, postIDC, putIDC, putIDCRemark, deleteIDC } from '@/api/idc.js'
  import { getIDCTree } from '@/api/idc.js'
  import { postLine, putLine, delLine } from '@/api/idc.js'
  import { updateAllIPAddrsNetInfo, updatePartIPAddrsNetInfo } from '@/api/idc.js'
  // import { getLine, getLines, postLine, putLine, delLine } from '@/api/idc.js'
  import { getLineIpAddrs, putLineIpAddrs, createLabelForAllIPs } from '@/api/idc.js'

  export default {
    data() {
      const data = [];
      function validateStr (rule, value, callback) {
        if (value === '' || typeof value === 'undefined' || value == null) {
          callback(new Error('请输入正确名称'))
        } else {
          const reg = /\s+/
          if (reg.test(value)) {
            callback(new Error('名称中不允许有空字符'))
          } else {
            callback()
          }
        }
      };
      return {
        data: JSON.parse(JSON.stringify(data)),
        linedetailinfo: {
          'ipaddrs_net_line': '',  // IP地址列表
          'remark_info': ''  // 备注信息
        },
        dialogLineVisible: false,
        dialogIDCVisible: false,
        dialogIDCTitle: '',
        Lineform: {
          id: 0,
          tree_id: 0,
          tree_type: 'line',
          label: '',
          children: [],
        },
        IDCform: {
          id: 0,
          tree_id: 0,
          tree_type: 'idc',
          label: '',
          children: [],
        },
        lineRules: {
          label: [
            { required: true, message: '请输入机房名称', validator: validateStr, trigger: ['blur'] }
          ],
        },
        IDCRules: {
          label: [
            { required: true, message: '请输入机房名称', validator: validateStr, trigger: ['blur'] }
          ],
        },
        currentDataPoint: null,
        formLabelWidth: '80px',
        idcName: '添加线路',
        idEdit: false,
        currentPoolObj: null,
        currentTitle: '',
        currentIDCTitle: '',
        currentLineTitle: '',
        pushing_all: false,
        pushing_part: false,
        pushing_create_label: false,
        should_disabled: false,
        searchContent: '',
        search_ip: false,
      }
    },

    mounted () {
      this.doGetTree()
    },

    methods: {
      doGetTree(show_notice){
        getIDCTree({content: this.searchContent, search_ip: this.search_ip}).then(r => {
          this.data = r.data.tree
          tree_id = r.data.id + 1
          if (show_notice) {
            this.$notify({
              title: '成功',
              message: show_notice,
              type: 'success'
            });
          }
        }).catch(e => console.log(e))
      },
      onSearch(){
         this.doGetTree()
      },
      append(data) {
        this.currentDataPoint = data
        // this.Lineform.id = data.id
        this.idcName = '添加线路：' + data.label
        this.Lineform = {
          id: data.id,    // 因为是增加新的，此时这个ID指向idc_id
          tree_id: 0,
          tree_type: 'line',
          label: '',
          children: [],
        },
        this.idEdit = false
        this.dialogLineVisible = true
      },
      lineAppend(formName){
        this.$refs[formName].validate((valid) => {
          if (valid) {
            if (this.idEdit === false) {
              postLine({label: this.Lineform.label, idc_id: this.Lineform.id}).then(r=>{
                this.doGetTree()
                this.dialogLineVisible = false
              }).catch(e=>{
                this.dialogLineVisible = false
                console.log(e)
              })
            } else {
              putLine({label: this.Lineform.label, id: this.Lineform.id}).then(r=>{
                this.doGetTree()
                this.dialogLineVisible = false
              }).catch(e=>{
                this.dialogLineVisible = false
                console.log(e)
              })
            }
          } else {
            return false
          }
        })
      },
      idcAppendConfirm(formName){
        this.$refs[formName].validate((valid) => {
          if (valid) {
            if (this.idEdit === false) {
              postIDC({'label': this.IDCform.label}).then(r=>{
                this.doGetTree()
                this.dialogIDCVisible = false
              }).catch(e=> {
                console.log(e)
                this.dialogIDCVisible = false
              })
            } else {
              putIDC({'label': this.IDCform.label, 'id': this.IDCform.id}).then(r=>{
                this.doGetTree()
                this.dialogIDCVisible = false
              }).catch(e=> {
                console.log(e)
                this.dialogIDCVisible = false
              })
            }
          } else {
            return false
          }
        })
      },
      idcAppend(data){
        this.currentDataPoint = data
        this.idEdit = false
        this.dialogIDCTitle = '增加机房: '
        this.dialogIDCVisible = true
      },
      flushtree(){
        this.doGetTree('刷新列表成功！')
      },
      removeIDC(node, data) {
        this.$confirm('是否确定删除？', '确认信息', {
          distinguishCancelAndClose: true,
          confirmButtonText: '确定',
          cancelButtonText: '放弃'
        }).then(_ => {
          deleteIDC({id: data.id}).then(r=>{
            this.$notify({
              title: '成功',
              message: '删除成功！',
              type: 'success'
            });
            this.doGetTree()
          }).catch(e=>{
            console.log(e)
          })
        }).catch(e => console.log(e))
      },
      removeLine(node, data){
        this.$confirm('是否确定删除？', '确认信息', {
          distinguishCancelAndClose: true,
          confirmButtonText: '确定',
          cancelButtonText: '放弃'
        }).then(_ => {
          delLine({id: data.id}).then(r=>{
          this.$notify({
            title: '成功',
            message: '删除成功！',
            type: 'success'
          });
            this.doGetTree()
          }).catch(e=>{
            console.log(e)
          })
        }).catch(e => console.log(e))
      },
      edit(node, data){
        // console.log(data)
        this.IDCform = {...data}
        this.idEdit = true
        this.dialogIDCTitle = '编辑机房属性: ' + data.label
        this.dialogIDCVisible = true
      },
      editLine(node, data){
        // console.log(data)
        this.Lineform = {...data}
        this.idEdit = true
        this.dialogLineVisible = true
      },
      renderContent(h, { node, data, store }) {
        return h("span", {
          class: "custom-tree-node"
        }, h("span", null, node.label), h("span", null, h("a", {
          onClick: () => this.append(data)
        }, "Append "), h("a", {
          onClick: () => this.remove(node, data)
        }, "Delete")));
      },
      nodeClick(node, tree, event){
        this.currentPoolObj = node
        if (node.tree_type === 'idc') {
          this.currentIDCTitle = tree.data.label
          this.currentLineTitle = ''
          this.should_disabled = true
          getIDC({id: node.id}).then(r=>{
            this.linedetailinfo = {
              'ipaddrs_net_line': '请把此机房的IP信息填写入相应的线路中',  // IP地址列表
              'remark_info': r.data.remark        // 备注信息
            }
          }).catch(e=>console.log(e))
        } else if (node.tree_type === 'line') {
          this.should_disabled = false
          this.currentIDCTitle = tree.parent.data.label
          this.currentLineTitle = node.label
          getLineIpAddrs({id: node.id}).then(r=>{
            this.linedetailinfo = {
              'ipaddrs_net_line': r.data.ipaddrs,  // IP地址列表
              'remark_info': r.data.remark        // 备注信息
            }
          }).catch(e=>{
            console.log(e)
          })
        }
      },
      putIDCOrLineIPData(){
        // if (!this.currentPoolObj) {
        //   return
        // }
        if (this.currentPoolObj.tree_type === 'idc') {
          putIDCRemark({
            "id": this.currentPoolObj.id,
            "remark": this.linedetailinfo.remark_info
          }).then(r=>{
              this.$notify({
                title: '成功',
                message: '更新成功！',
                type: 'success'
              });
          }).catch(e=>console.log(e))
        } else if (this.currentPoolObj.tree_type === 'line') {
          putLineIpAddrs({
            "id": this.currentPoolObj.idc_id,
            "line_id": this.currentPoolObj.id,
            "ipaddrs": this.linedetailinfo.ipaddrs_net_line,
            "remark": this.linedetailinfo.remark_info
          }).then(r=>{
              this.$notify({
                title: '成功',
                message: '更新成功！',
                type: 'success'
              });
          }).catch(e=>console.log(e))
        } 
      },
      updateAllIPAddrs(){
        this.pushing_all = true
        updateAllIPAddrsNetInfo().then(r=>{
          this.$notify({
            title: '成功',
            message: '更新成功！',
            type: 'success'
          });
          this.pushing_all = false
        }).catch(e=>{
          console.log(e)
          this.pushing_all = false
        })
      },
      updatePartIPAddrs(){
        this.pushing_part = true
        updatePartIPAddrsNetInfo().then(r=>{
          this.$notify({
            title: '成功',
            message: '更新成功！',
            type: 'success'
          });
          this.pushing_part = false
        }).catch(e=>{
          console.log(e)
          this.pushing_part = false
        })
      },
      doCreateLabelForAllIPs(){
        this.pushing_create_label = true
        createLabelForAllIPs().then(r=>{
          this.$notify({
            title: '成功',
            message: '生成成功！',
            type: 'success'
          });
          this.pushing_create_label = false
        }).catch(e=>{
          this.pushing_create_label = false
          console.log(e)
        })
      }
    }
  };
</script>

<style scoped>
.custom-tree-node {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: space-between;
  font-size: 14px;
  padding-right: 8px;
}
.box-board {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.text {
  font-size: 14px;
}
.item {
  margin-bottom: 18px;
}
.box-card-left {
  width: 55%;
}
.box-card-left-body{
  display: flex;
  flex-wrap: nowrap;
}
.box-card-left-body-content {
  width: 85%;
}
.box-card-left-body-action {
  display: flex;
  flex-wrap: wrap;
  flex-direction: column;
  width: 150px;
  max-width: 150px;
  /* border-left: 1px solid #F00 */
}
.box-card-left-body-action :deep() .el-button {
  margin-bottom: 20px;
  margin-left: 10px!important;
}
.box-card-right {
  width: 44%;
}

.card-scrollbar {
  height: 75vh;
}
</style>