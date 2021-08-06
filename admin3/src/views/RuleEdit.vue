<template>
  <div class="box-card">
    <div class="box-member">
      <el-descriptions :column="1" size="mini" border>
        <el-descriptions-item>
          <template #label>
            <span>规则功能描述：</span>
            <el-tooltip content="描述此规则的功能或者作用" placement="top">
              <i class="el-icon-info"></i>
            </el-tooltip>
          </template>
          <el-input size="mini" v-model="formData.description"></el-input>
        </el-descriptions-item>
        <el-descriptions-item>
          <template #label>
            <span>告警名称(alert)：</span>
            <el-tooltip
              content="警报规则的名称。alert: <string>"
              placement="top"
            >
              <i class="el-icon-question"></i>
            </el-tooltip>
          </template>
          <el-input size="mini" v-model="formData.alert"></el-input>
        </el-descriptions-item>
        <el-descriptions-item>
          <template #label>
            <span>表达式(expr)：</span>
            <el-tooltip
              content="使用PromQL表达式完成的警报触发条件，用于计算是否有满足触发条件。expr: <string>"
              placement="top"
            >
              <i class="el-icon-question"></i>
            </el-tooltip>
          </template>
          <el-input
            type="textarea"
            size="mini"
            wrap="off"
            :rows="4"
            maxlength="5000"
            show-word-limit
            v-model="formData.expr"
          ></el-input>
        </el-descriptions-item>
        <el-descriptions-item>
          <template #label>
            <span>等待时间(for)：</span>
            <el-tooltip
              content="带有for子句的警报触发以后首先会先转换成 Pending 状态，然后在转换为 Firing 状态。这里需要俩个周期才能触发警报条件，如果没有设置 for 子句，会直接 Inactive 状态转换成 Firing状态，然后触发警报，发送给 Receiver 设置的通知人。for: <string>，如：for: 5m，(m代表分钟)"
              placement="top"
              popper-class="filed-explain"
            >
              <i class="el-icon-question"></i>
            </el-tooltip>
          </template>
          <el-input size="mini" v-model="formData.for"></el-input>
        </el-descriptions-item>
        <el-descriptions-item>
          <template #label>
            <span>标签(labels)：</span>
            <el-tooltip
              content="自定义标签，允许自行定义标签附加在警报上。<lable_name>: <label_value>"
              placement="top"
            >
              <i class="el-icon-question"></i>
            </el-tooltip>
          </template>
          <div>
            <div v-for="data in formData.labels" :key="data.id">
              <el-card
                class="box-card-two"
                :body-style="{ padding: '0px 0px 0px 0px' }"
                shadow="never"
              >
                <el-select
                  v-model="data.key"
                  :key="data.id"
                  size="mini"
                  filterable
                  allow-create
                  default-first-option
                  placeholder="请选择"
                  style="width: 20%"
                >
                  <el-option :label="data.key" :value="data.key"></el-option>
                  <el-option
                    v-for="item in defaultLabels"
                    :key="item.id"
                    :label="item.label"
                    :value="item.label"
                  >
                  </el-option>
                </el-select>
                <el-input
                  size="mini"
                  style="width: 77%"
                  wrap="off"
                  v-model="data.value"
                  maxlength="100"
                  show-word-limit
                  :key="data.id"
                ></el-input>
                <el-popconfirm
                  title="确定删除吗？"
                  @confirm="delItem(data, 'labels')"
                >
                  <template #reference>
                    <el-button
                      type="text"
                      class="el-icon-delete-solid icon-action"
                      size="mini"
                    ></el-button>
                  </template>
                </el-popconfirm>
              </el-card>
            </div>
            <div style="text-align: right; padding-right: 14px">
              <el-button
                type="success"
                plain
                size="mini"
                icon="el-icon-circle-plus"
                @click="addLables('labels')"
                >增加</el-button
              >
            </div>
          </div>
        </el-descriptions-item>
        <el-descriptions-item>
          <template #label>
            <span>注释(annotations)：</span>
            <el-tooltip
              content="用来设置有关警报的一组描述信息，其中包括自定义的标签，以及expr计算后的值。<lable_name>: <tmpl_string>"
              placement="top"
              popper-class="filed-explain"
            >
              <i class="el-icon-question"></i>
            </el-tooltip>
          </template>
          <div>
            <div v-for="data in formData.annotations" :key="data.id">
              <el-card
                class="box-card-two"
                :body-style="{ padding: '0px 0px 0px 0px' }"
                shadow="never"
              >
                <div class="annotations-move">
                  <div class="annotations-left">
                    <el-select
                      v-model="data.key"
                      :key="data.id"
                      allow-create
                      size="mini"
                      filterable
                      default-first-option
                      placeholder="请选择"
                      style="width: 100%"
                    >
                      <el-option
                        :label="data.key"
                        :value="data.key"
                      ></el-option>
                      <el-option
                        v-for="item in defaultLabels"
                        :key="item.id"
                        :label="item.label"
                        :value="item.label"
                      >
                      </el-option>
                    </el-select>
                  </div>
                  <div class="annotations-right">
                    <el-input
                      type="textarea"
                      :rows="4"
                      wrap="off"
                      size="mini"
                      show-word-limit
                      style="width: 100%"
                      v-model="data.value"
                      :key="data.id"
                    ></el-input>
                  </div>
                  <div>
                    <el-popconfirm
                      title="确定删除吗？"
                      @confirm="delItem(data, 'annotations')"
                    >
                      <template #reference>
                        <el-button
                          type="text"
                          class="el-icon-delete-solid icon-action"
                          size="mini"
                        ></el-button>
                      </template>
                    </el-popconfirm>
                  </div>
                </div>
              </el-card>
            </div>
            <div style="text-align: right; padding-right: 14px">
              <el-button
                type="success"
                size="mini"
                plain
                @click="addLables('annotations')"
                icon="el-icon-circle-plus"
                >增加</el-button
              >
            </div>
          </div>
        </el-descriptions-item>
        <el-descriptions-item>
          <template #label>
            <span>是否启用</span>
            <el-tooltip content="是否启用" placement="top">
              <i class="el-icon-question"></i>
            </el-tooltip>
          </template>
          <el-switch
            v-model="formData.enabled"
            active-color="#13ce66"
            inactive-color="#ff4949"
          ></el-switch>
        </el-descriptions-item>
        <el-descriptions-item>
          <div class="rule-edit-area">
            <el-button
              style="margin-right: 10px"
              icon="el-icon-edit"
              type="primary"
              plain
              size="mini"
              @click="viewData"
              >预览</el-button
            >
            <el-button
              style="margin-right: 10px"
              icon="el-icon-edit"
              type="warning"
              size="mini"
              @click="importData"
              >导入</el-button
            >
            <el-button
              style="margin-right: 10px"
              icon="el-icon-upload"
              type="primary"
              size="mini"
              @click="onSubmit"
              >提交</el-button
            >
          </div>
        </el-descriptions-item>
      </el-descriptions>
    </div>
    <div class="dialog-area">
      <el-dialog
        title="导入规则数据-yaml格式"
        v-model="dialogFormVisible"
        custom-class="dialog-custom-class"
      >
        <el-descriptions class="margin-top" :column="1" size="mini" border>
          <el-descriptions-item>
            <template #label>
              <!-- <i class="el-icon-user"></i> -->
              说明
            </template>
            导入的数据，必须是符合yaml语法，格式可以参考以下样式
          </el-descriptions-item>
          <el-descriptions-item>
            <template #label>样例</template>
            <pre v-highlight="code"><code></code></pre>
          </el-descriptions-item>
          <el-descriptions-item>
            <template #label>输入内容</template>
            <el-input
              type="textarea"
              placeholder="请输入内容"
              v-model="textarea"
              rows="8"
              multiple
              wrap="off"
            >
            </el-input>
          </el-descriptions-item>
          <el-descriptions-item>
            <template #label>结果</template>
            <el-scrollbar height="80px">
              <div v-if="showError">
                <pre v-highlight="error"><code></code></pre>
              </div>
            </el-scrollbar>
          </el-descriptions-item>
          <el-descriptions-item>
            <div class="dialog-action">
              <el-button size="mini" type="info" @click="closeImport"
                >关闭</el-button
              >
              <el-button size="mini" type="primary" @click="parseYaml"
                >导入</el-button
              >
            </div>
          </el-descriptions-item>
        </el-descriptions>
      </el-dialog>
    </div>
    <div class="viewer-area">
      <el-dialog
        title="预览规则-yaml格式"
        v-model="viewDialogVisible"
        custom-class="dialog-custom-class"
      >
        <el-descriptions class="margin-top" :column="1" size="mini" border>
          <el-descriptions-item>
            <template #label>YAML格式数据</template>
            <div style="width: 630px">
              <el-scrollbar height="200px" class="flex-content">
                <pre v-highlight="viweCode"><code></code></pre>
              </el-scrollbar>
            </div>
          </el-descriptions-item>
          <el-descriptions-item>
            <div class="dialog-action">
              <el-button size="mini" type="primary" @click="closeViewDialog"
                >关闭</el-button
              >
            </div>
          </el-descriptions-item>
        </el-descriptions>
      </el-dialog>
    </div>
  </div>
</template>

<script>
import {
  getRuleDetail,
  getDefaultEnableLables,
  putNodeInfo,
  postNodeInfo,
  deleteNodeLable
} from '@/api/monitor.js'

import * as yaml from 'js-yaml'

export default ({
  props: {
    query: {
      type: Object,
      default: null
    }
  },
  components: {
  },
  data () {
    return {
      formData: {
        alert: '',
        expr: '',
        for: '',
        labels: [],
        annotations: [],
        enabled: false,
        description: ''
      },
      defaultLabels: [],
      submitType: '',
      formDisabled: true,
      switchValue: false,
      dialogFormVisible: false,
      code: `description: 功能描述
alert: PrometheusJobMissing
expr: absent(up{job="prometheus"})
for: 0m
labels:
  severity: warning
annotations:
  description: A Prometheus job has disappeared VALUE = {{ $value }} LABELS = {{ $labels }}
  summary: Prometheus job missing (instance {{ $labels.instance }})
`,
      textarea: '',
      error: '',
      showError: false,
      viewDialogVisible: false,
      viweCode: '',
      nodeInfo: null
    }
  },
  methods: {
    doGetRuleDetail (info) {
      if (!info) {
        return false
      }
      if (info.label === '[must rename me]') {
        return false
      }
      this.nodeInfo = info
      getDefaultEnableLables().then(
        r => {
          this.defaultLabels = r.data
          getRuleDetail(info).then(
            r => {
              this.formData = r.data
              this.submitType = 'put'
            }
          ).catch(
            e => { console.log(e) }
          )
        }
      ).catch(
        e => { console.log(e) }
      )
    },
    flush () {
      this.doGetRuleDetail(this.nodeInfo)
    },
    onSubmit () {
      const nodeData = this.formData
      if (this.submitType === 'put') {
        putNodeInfo(nodeData).then(
          r => {
            this.formData = r.data
            this.$emit('fatherMethod', nodeData, 'put');
            this.$notify({
              title: '成功',
              message: '更新成功！',
              type: 'success'
            });
          }
        ).catch(
          e => { console.log(e) }
        )
      } else if (this.submitType === 'post') {
        postNodeInfo(nodeData).then(
          r => {
            this.formData = r.data
            this.$emit('fatherMethod', nodeData, 'post');
            this.$notify({
              title: '成功',
              message: '创建成功！',
              type: 'success'
            });
          }
        ).catch(
          e => { console.log(e) }
        )
      } else {
        this.$notify({
          title: '失败',
          message: '不支持的操作！',
          type: 'error'
        });
      }
    },
    addLables (witch) {
      getDefaultEnableLables().then(
        r => {
          this.defaultLabels = r.data
          let newID = 0
          if (witch === 'annotations') {
            if (this.formData.annotations.length !== 0) {
              newID = this.formData.annotations[this.formData.annotations.length - 1].id + 1
            }
            this.formData.annotations.push({ id: newID, key: '', value: '', is_new: true })
          } else if (witch === 'labels') {
            if (this.formData.annotations.length !== 0) {
              newID = this.formData.labels[this.formData.labels.length - 1].id + 1
            }
            this.formData.labels.push({ id: newID, key: '', value: '', is_new: true })
          }
          this.formData = { ...this.formData }
        }
      ).catch(
        e => { console.log(e) }
      )
    },
    delItem (data, lType) {
      const labelInfo = data.valueOf()
      if (labelInfo.is_new) {
        if (lType === 'labels') {
          const index = this.formData.labels.findIndex(d => d.id === data.id);
          this.formData.labels.splice(index, 1)
        } else if (lType === 'annotations') {
          const index = this.formData.annotations.findIndex(d => d.id === data.id);
          this.formData.annotations.splice(index, 1);
        }
        this.formData = { ...this.formData }
      } else {
        deleteNodeLable(labelInfo, { type: lType }).then(
          r => {
            this.$notify({
              title: '成功',
              message: '删除成功！',
              type: 'success'
            });
            if (lType === 'labels') {
              const index = this.formData.labels.findIndex(d => d.id === data.id);
              this.formData.labels.splice(index, 1)
            } else if (lType === 'annotations') {
              const index = this.formData.annotations.findIndex(d => d.id === data.id);
              this.formData.annotations.splice(index, 1);
            }
            this.formData = { ...this.formData }
          }
        ).catch(
          e => { console.log(e) }
        )
      }
    },
    setFormDisable () {
      this.formDisabled = true
    },
    setFormEnable () {
      this.formDisabled = false
    },
    resetForm () {
      this.formData = {}
      // this.$refs['formRef'].resetFields();
    },
    closeImport () {
      this.textarea = ''
      this.error = ''
      this.showError = false
      this.dialogFormVisible = false
    },
    closeViewDialog () {
      this.viweCode = ''
      this.viewDialogVisible = false
    },
    viewData () {
      let textarea = []
      textarea.push(`alert: ${this.formData.alert}`)
      textarea.push(`expr: ${this.formData.expr}`)
      textarea.push(`for: ${this.formData.for}`)
      textarea.push(`labels:`)
      this.formData.labels.forEach(item => {
        const value = item.value.replace(/\n/g, '\\n')
        textarea.push(`  ${item.key}: ${value}`)
      })
      textarea.push(`annotations:`)
      this.formData.annotations.forEach(item => {
        const value = item.value.replace(/\n/g, '\\n')
        textarea.push(`  ${item.key}: ${value}`)
      })
      this.viweCode = textarea.join('\n')
      this.viewDialogVisible = true
    },
    importData () {
      this.showError = false
      this.error = ''
      let textarea = []
      textarea.push(`description: ${this.formData.description}`)
      textarea.push(`alert: ${this.formData.alert}`)
      textarea.push(`expr: ${this.formData.expr}`)
      textarea.push(`for: ${this.formData.for}`)
      textarea.push(`labels:`)
      this.formData.labels.forEach(item => {
        const value = item.value.replace(/\n/g, '\\n')
        textarea.push(`  ${item.key}: ${value}`)
      })
      textarea.push(`annotations:`)
      this.formData.annotations.forEach(item => {
        const value = item.value.replace(/\n/g, '\\n')
        textarea.push(`  ${item.key}: ${value}`)
      })
      this.textarea = textarea.join('\n')
      this.dialogFormVisible = true
    },
    parseYaml () {
      this.error = ''
      this.showError = false
      try {
        let yamlContext = yaml.load(this.textarea)
        const message = '数据格式正确，解析正常。分析到的字段有：'
        let fields = []
        if (yamlContext.description) {
          this.formData.description = yamlContext.description
          fields.push('description')
        }
        if (yamlContext.alert) {
          this.formData.alert = yamlContext.alert
          fields.push('alert')
        }
        if (yamlContext.expr) {
          this.formData.expr = yamlContext.expr
          fields.push('expr')
        }
        if (yamlContext.for) {
          this.formData.for = yamlContext.for
          fields.push('for')
        }
        // { id: newID, key: '', value: '', is_new: true }
        let genID = 0
        if (yamlContext.labels) {
          fields.push('labels')
          for (const key in yamlContext.labels) {
            let haveData = false
            this.formData.labels.map(item => {
              if (item.key === key) {
                item.value = yamlContext.labels[key].replace(/\\n/g, '\n')
                haveData = true
              }
              if (genID !== 0 && item.id < genID) {
                genID = item.id
              }
            })
            if (haveData === false) {
              this.formData.labels.push({
                id: genID,
                key: key,
                value: yamlContext.labels[key].replace(/\\n/g, '\n'),
                is_new: true
              })
              genID += 1
            }
          }
        }
        genID = 0
        if (yamlContext.annotations) {
          fields.push('annotations')
          for (const key in yamlContext.annotations) {
            let haveData = false
            this.formData.annotations.map(item => {
              if (item.key === key) {
                item.value = yamlContext.annotations[key].replace(/\\n/g, '\n')
                haveData = true
              }
              if (genID !== 0 && item.id < genID) {
                genID = item.id
              }
            })
            if (haveData === false) {
              this.formData.annotations.push({
                id: genID,
                key: key,
                value: yamlContext.annotations[key].replace(/\\n/g, '\n'),
                is_new: true
              })
              genID += 1
            }
          }
        }
        this.error = message + fields.join(",")
        this.showError = true
      } catch (e) {
        this.error = e.toString()
        this.showError = true
        // this.$nextTick(() => {//重新渲染分页
        //   this.showError = true
        // });
        console.log(e)
      }
    }
  }
})
</script>


<style scoped>
.annotations-move {
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
}
.annotations-left {
  width: 21%;
}
.annotations-right {
  width: 79%;
}
.rule-edit-area {
  width: 680px;
  align-content: right;
  text-align: right;
}
.dialog-action {
  align-content: right;
  text-align: right;
}
.dialog-area :deep() .el-dialog {
  margin: 20px auto !important;
}
.dialog-area :deep() .el-dialog__body {
  padding-top: 5px;
}
.box-member :deep() .annotations-labels {
  display: flex !important;
  flex-direction: row;
  /* flex-wrap: nowrap; */
  justify-content: space-between;
  /* width: 800px; */
}
.box-member {
  width: 100%;
}
.box-card {
  /* width: 98%; */
  margin-left: 5px;
  width: 840px;
  padding: 0px 0px;
}
.box-member :deep() .el-form-item__label {
  padding-bottom: 0px !important;
  /* width: 800px; */
  width: 100%;
}
.box-member :deep() .el-tabs__content {
  padding-right: 5px !important;
}
.box-member :deep() .el-descriptions__content {
  padding: 0px 0px 0px 0px;
}
</style>