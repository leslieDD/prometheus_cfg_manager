<template>
  <el-card
    class="box-card"
    :body-style="{ padding: '10px 0px 0px 0px' }"
    shadow="never"
  >
    <div class="box-member">
      <el-form
        label-position="right"
        ref="formRef"
        :model="formData"
        label-width="170px"
        size="small"
        :disabled="formDisabled"
      >
        <el-form-item>
          <template #label>
            <span>告警名称(alert)：</span>
            <el-tooltip
              content="警报规则的名称。alert: <string>"
              placement="top"
            >
              <i class="el-icon-question"></i>
            </el-tooltip>
          </template>
          <el-input v-model="formData.alert"></el-input>
        </el-form-item>
        <el-form-item>
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
            :rows="6"
            maxlength="5000"
            show-word-limit
            v-model="formData.expr"
          ></el-input>
        </el-form-item>
        <el-form-item>
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
          <el-input v-model="formData.for"></el-input>
        </el-form-item>
        <el-form-item>
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
                  style="width: 77%"
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
                @click="addLables(formData, 'labels')"
                >增加</el-button
              >
            </div>
          </div>
        </el-form-item>
        <el-form-item>
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
                      :rows="2"
                      maxlength="500"
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
                @click="addLables(formData, 'annotations')"
                icon="el-icon-circle-plus"
                >增加</el-button
              >
            </div>
          </div>
        </el-form-item>
        <el-form-item>
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
        </el-form-item>
        <el-form-item align="right">
          <el-button
            style="margin-right: 10px"
            icon="el-icon-upload"
            type="primary"
            size="medium"
            @click="onSubmit"
            >提交数据</el-button
          >
        </el-form-item>
      </el-form>
    </div>
  </el-card>
</template>

<script>
// import KeyValue from '@/components/KeyValue.vue'
import { getRuleDetail, getDefaultLabels, putNodeInfo, postNodeInfo, deleteNodeLable } from '@/api/monitor.js'

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
        enabled: false
      },
      defaultLabels: [],
      submitType: '',
      formDisabled: true,
      switchValue: false
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
      getDefaultLabels().then(
        r => {
          this.defaultLabels = r.data
          //   if (info.is_new) {
          //     this.formData = {
          //       "id": 0,
          //       "alert": info.label,
          //       "expr": "",
          //       "for": "",
          //       "sub_group_id": info.parent,
          //       "labels": [],
          //       "annotations": []
          //     }
          //     this.submitType = 'post'
          //   } else {
          //     getRuleDetail(info).then(
          //       r => {
          //         this.formData = r.data
          //         this.submitType = 'put'
          //       }
          //     ).catch(
          //       e => { console.log(e) }
          //     )
          //   }
          getRuleDetail(info).then(
            r => {
              this.formData = r.data
              this.submitType = 'put'
            }
          ).catch(
            e => { console.log(e) }
          )
          //   this.doGetRuleDetail(info)
        }
      ).catch(
        e => { console.log(e) }
      )
    },
    onSubmit () {
      const nodeData = this.formData.valueOf()
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
            // this.doGetRuleDetail(this.storeInfo)
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
            // this.doGetRuleDetail(this.storeInfo)
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
    addLables (label, witch) {
      getDefaultLabels().then(
        r => {
          this.defaultLabels = r.data
          let newID = 0
          if (witch === 'annotations') {
            if (this.formData.annotations.length !== 0) {
              newID = this.formData.annotations[(this.formData.annotations.length - 1)].id + 1
            }
            this.formData.annotations.push({ id: newID, key: '', value: '', is_new: true })
          } else if (witch === 'labels') {
            if (this.formData.annotations.length !== 0) {
              newID = this.formData.labels[(this.formData.labels.length - 1)].id + 1
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
      this.$refs['formRef'].resetFields();
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
</style>