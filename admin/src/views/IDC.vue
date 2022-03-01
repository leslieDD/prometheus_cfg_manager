<template>
  <div class="box-board">
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span>机房及线路</span>
          <el-button size="small" class="button" @click="idcAppend"> 增加机房 </el-button>
        </div>
      </template>
      <div>
        <el-tree
          :data="data"
          node-key="id"
          highlight-current
          default-expand-all
          :expand-on-click-node="false"
          @node-click="nodeClick"
        >
          <template #default="{ node, data }">
            <span class="custom-tree-node">
              <span>{{ node.label }}</span>
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
      </div>
    </el-card>
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span>IP地址</span>
          <el-button size="small" class="button" @click="putLineIPData"> 更新 </el-button>
        </div>
      </template>
      <div>
        <el-input
          v-model="ipaddrs_net_line"
          :autosize="{ minRows: 30 }"
          type="textarea"
          placeholder=""
        />
      </div>
    </el-card>
    <el-dialog v-model="dialogLineVisible" :title="idcName" width="400px">
      <el-form :model="Lineform" :rules="lineRules" ref="LineformRef">
        <el-form-item label="线路名称" :label-width="formLabelWidth">
          <el-input size="small" v-model="Lineform.label" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button size="small" @click="dialogLineVisible = false">取消</el-button>
          <el-button size="small" type="primary" @click="lineAppend('LineformRef')"
            >确定</el-button
          >
        </span>
      </template>
    </el-dialog>
    <el-dialog v-model="dialogIDCVisible" title="添加新机房" width="400px">
      <el-form :model="IDCform" :rules="IDCRules" ref="IDCformRef">
        <el-form-item label="机房名称" :label-width="formLabelWidth">
          <el-input size="small" v-model="IDCform.label" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button size="small" @click="dialogIDCVisible = false">取消</el-button>
          <el-button size="small" type="primary" @click="idcAppendConfirm('IDCformRef')"
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
  import { postIDC, putIDC, deleteIDC } from '@/api/idc.js'
  import { getIDCTree } from '@/api/idc.js'
  import { postLine, putLine, delLine } from '@/api/idc.js'
  // import { getLine, getLines, postLine, putLine, delLine } from '@/api/idc.js'
  import { getLineIpAddrs, putLineIpAddrs } from '@/api/idc.js'

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
        ipaddrs_net_line: '',
        dialogLineVisible: false,
        dialogIDCVisible: false,
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
      }
    },

    mounted () {
      this.doGetTree()
    },

    methods: {
      doGetTree(){
        getIDCTree().then(r => {
          this.data = r.data.tree
          tree_id = r.data.id + 1
        }).catch(e => console.log(e))
      },
      append(data) {
        this.currentDataPoint = data
        this.idcName = '添加线路：' + data.label
        this.idEdit = false
        this.dialogLineVisible = true
      },
      lineAppend(formName){
        this.$refs[formName].validate((valid) => {
          if (valid) {
            if (this.idEdit === false) {
              postLine({label: this.Lineform.label, id: this.Lineform.id}).then(r=>{
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
        this.dialogIDCVisible = true
      },
      removeIDC(node, data) {
        this.$confirm('是否确定删除？', '确认信息', {
          distinguishCancelAndClose: true,
          confirmButtonText: '确定',
          cancelButtonText: '放弃'
        }).then(_ => {
          deleteIDC({id: data.id}).then(r=>{
            this.doGetTree()
          }).catch(e=>{
            console.log(e)
          })
          const parent = node.parent;
          const children = parent.data.children || parent.data;
          const index = children.findIndex(d => d.tree_id === data.tree_id);
          children.splice(index, 1);
        }).catch(e => console.log(e))
      },
      removeLine(node, data){
        this.$confirm('是否确定删除？', '确认信息', {
          distinguishCancelAndClose: true,
          confirmButtonText: '确定',
          cancelButtonText: '放弃'
        }).then(_ => {
          delLine({id: data.id}).then(r=>{
            this.doGetTree()
          }).catch(e=>{
            console.log(e)
          })
          const parent = node.parent;
          const children = parent.data.children || parent.data;
          const index = children.findIndex(d => d.tree_id === data.tree_id);
          children.splice(index, 1);
        }).catch(e => console.log(e))
      },
      edit(node, data){
        // console.log(data)
        this.IDCform = {...data}
        this.idEdit = true
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
        // console.log(node, tree, event)
        if (node.tree_type !== 'line') {
          return
        }
        this.currentPoolObj = node
        getLineIpAddrs({id: node.id}).then(r=>{
          this.ipaddrs_net_line = r.data.ipaddrs
        }).catch(e=>{
          console.log(e)
        })
      },
      putLineIPData(){
        if (!this.currentPoolObj) {
          return
        }
        putLineIpAddrs({
          "id": this.currentPoolObj.idc_id,
          "line_id": this.currentPoolObj.id,
          "ipaddrs": this.ipaddrs_net_line
        }).then(r=>{
          
        }).catch(e=>console.log(e))
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

.box-card {
  width: 49%;
}
</style>