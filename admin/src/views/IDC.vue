<template>
  <div class="box-board">
    <el-card class="box-card-left" shadow="never">
      <template #header>
        <div class="card-header">
          <span class="title-left">机房及线路<span class="download-btn"><el-button
            icon="el-icon-download"
            size="small"
            type="text"
            @click="doOutportXls()"
          >点击导出</el-button
        ></span></span>
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
              node-key="tree_id"
              :highlight-current="true"
              :expand-on-click-node="true"
              :default-expanded-keys="defaultExpandedKeys"
              @node-click="nodeClick"
              @node-expand="nodeExpand"
              @node-collapse="nodeCollapse"
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
                    <el-tag size="small" @click="edit(node, data)"> <svg t="1639990532110" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="12171" xmlns:xlink="http://www.w3.org/1999/xlink" width="13" height="13" data-v-042ca774=""><path d="M199.04 672.64l193.984 112 224-387.968-193.92-112-224 388.032z m-23.872 60.16l32.896 148.288 144.896-45.696-177.792-102.592zM455.04 229.248l193.92 112 56.704-98.112-193.984-112-56.64 98.112zM104.32 708.8l384-665.024 304.768 175.936-383.936 665.088h0.064l-248.448 78.336-56.448-254.336z m384 254.272v-64h448v64h-448z" p-id="12172"></path></svg> </el-tag>
                    <el-tooltip :content="expand_idc_title[data.id]" placement="right">
                     <span>
                      <el-tag v-if="data.expand === true" size="small" @click="do_expand_idc(data, false)"> <svg class="icon" width="13" height="13" viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg" data-v-78e17ca8=""><path fill="currentColor" d="M329.956 257.138a254.862 254.862 0 0 0 0 509.724h364.088a254.862 254.862 0 0 0 0-509.724H329.956zm0-72.818h364.088a327.68 327.68 0 1 1 0 655.36H329.956a327.68 327.68 0 1 1 0-655.36z"></path><path fill="currentColor" d="M694.044 621.227a109.227 109.227 0 1 0 0-218.454 109.227 109.227 0 0 0 0 218.454zm0 72.817a182.044 182.044 0 1 1 0-364.088 182.044 182.044 0 0 1 0 364.088z"></path></svg></el-tag>
                      <el-tag v-if="data.expand === false" type="info" size="small" @click="do_expand_idc(data, true)"> <svg class="icon" width="13" height="13" viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg" data-v-78e17ca8=""><path fill="currentColor" d="M329.956 257.138a254.862 254.862 0 0 0 0 509.724h364.088a254.862 254.862 0 0 0 0-509.724H329.956zm0-72.818h364.088a327.68 327.68 0 1 1 0 655.36H329.956a327.68 327.68 0 1 1 0-655.36z"></path><path fill="currentColor" d="M329.956 621.227a109.227 109.227 0 1 0 0-218.454 109.227 109.227 0 0 0 0 218.454zm0 72.817a182.044 182.044 0 1 1 0-364.088 182.044 182.044 0 0 1 0 364.088z"></path></svg></el-tag>
                     </span>
                    </el-tooltip>
                    <el-tooltip :content="view_idc_title[data.id]" placement="right">
                     <span>
                      <el-tag v-if="data.view === true" size="small" @click="do_expand_idc_view(data, false)"> <svg class="icon" width="13" height="13" viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg" data-v-78e17ca8=""><path fill="currentColor" d="M512 160c320 0 512 352 512 352S832 864 512 864 0 512 0 512s192-352 512-352zm0 64c-225.28 0-384.128 208.064-436.8 288 52.608 79.872 211.456 288 436.8 288 225.28 0 384.128-208.064 436.8-288-52.608-79.872-211.456-288-436.8-288zm0 64a224 224 0 1 1 0 448 224 224 0 0 1 0-448zm0 64a160.192 160.192 0 0 0-160 160c0 88.192 71.744 160 160 160s160-71.808 160-160-71.744-160-160-160z"></path></svg></el-tag>
                      <el-tag v-if="data.view === false" type="info" size="small" @click="do_expand_idc_view(data, true)"> <svg class="icon" width="13" height="13" viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg" data-v-78e17ca8=""><path d="M876.8 156.8c0-9.6-3.2-16-9.6-22.4-6.4-6.4-12.8-9.6-22.4-9.6-9.6 0-16 3.2-22.4 9.6L736 220.8c-64-32-137.6-51.2-224-60.8-160 16-288 73.6-377.6 176C44.8 438.4 0 496 0 512s48 73.6 134.4 176c22.4 25.6 44.8 48 73.6 67.2l-86.4 89.6c-6.4 6.4-9.6 12.8-9.6 22.4 0 9.6 3.2 16 9.6 22.4 6.4 6.4 12.8 9.6 22.4 9.6 9.6 0 16-3.2 22.4-9.6l704-710.4c3.2-6.4 6.4-12.8 6.4-22.4Zm-646.4 528c-76.8-70.4-128-128-153.6-172.8 28.8-48 80-105.6 153.6-172.8C304 272 400 230.4 512 224c64 3.2 124.8 19.2 176 44.8l-54.4 54.4C598.4 300.8 560 288 512 288c-64 0-115.2 22.4-160 64s-64 96-64 160c0 48 12.8 89.6 35.2 124.8L256 707.2c-9.6-6.4-19.2-16-25.6-22.4Zm140.8-96c-12.8-22.4-19.2-48-19.2-76.8 0-44.8 16-83.2 48-112 32-28.8 67.2-48 112-48 28.8 0 54.4 6.4 73.6 19.2L371.2 588.8ZM889.599 336c-12.8-16-28.8-28.8-41.6-41.6l-48 48c73.6 67.2 124.8 124.8 150.4 169.6-28.8 48-80 105.6-153.6 172.8-73.6 67.2-172.8 108.8-284.8 115.2-51.2-3.2-99.2-12.8-140.8-28.8l-48 48c57.6 22.4 118.4 38.4 188.8 44.8 160-16 288-73.6 377.6-176C979.199 585.6 1024 528 1024 512s-48.001-73.6-134.401-176Z" fill="currentColor"></path><path d="M511.998 672c-12.8 0-25.6-3.2-38.4-6.4l-51.2 51.2c28.8 12.8 57.6 19.2 89.6 19.2 64 0 115.2-22.4 160-64 41.6-41.6 64-96 64-160 0-32-6.4-64-19.2-89.6l-51.2 51.2c3.2 12.8 6.4 25.6 6.4 38.4 0 44.8-16 83.2-48 112-32 28.8-67.2 48-112 48Z" fill="currentColor"></path></svg></el-tag>
                     </span>
                    </el-tooltip>
                    <el-tag size="small" type="danger" @click="removeIDC(node, data)"> <svg class="icon" width="13" height="13" viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg" data-v-042ca774=""><path fill="currentColor" d="M352 192V95.936a32 32 0 0132-32h256a32 32 0 0132 32V192h256a32 32 0 110 64H96a32 32 0 010-64h256zm64 0h192v-64H416v64zM192 960a32 32 0 01-32-32V256h704v672a32 32 0 01-32 32H192zm224-192a32 32 0 0032-32V416a32 32 0 00-64 0v320a32 32 0 0032 32zm192 0a32 32 0 0032-32V416a32 32 0 00-64 0v320a32 32 0 0032 32z"></path></svg> </el-tag >
                  </span>
                  <span v-if="data.tree_type === 'line'">
                    <el-tag size="small" @click="editLine(node, data)"><svg t="1639990532110" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="12171" xmlns:xlink="http://www.w3.org/1999/xlink" width="13" height="13" data-v-042ca774=""><path d="M199.04 672.64l193.984 112 224-387.968-193.92-112-224 388.032z m-23.872 60.16l32.896 148.288 144.896-45.696-177.792-102.592zM455.04 229.248l193.92 112 56.704-98.112-193.984-112-56.64 98.112zM104.32 708.8l384-665.024 304.768 175.936-383.936 665.088h0.064l-248.448 78.336-56.448-254.336z m384 254.272v-64h448v64h-448z" p-id="12172"></path></svg></el-tag>
                    <el-tooltip :content="expand_line_title[data.id]" placement="right">
                     <span>
                      <el-tag v-if="data.expand === true" size="small" @click="do_expand_line(data, false)"> <svg class="icon" width="13" height="13" viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg" data-v-78e17ca8=""><path fill="currentColor" d="M329.956 257.138a254.862 254.862 0 0 0 0 509.724h364.088a254.862 254.862 0 0 0 0-509.724H329.956zm0-72.818h364.088a327.68 327.68 0 1 1 0 655.36H329.956a327.68 327.68 0 1 1 0-655.36z"></path><path fill="currentColor" d="M694.044 621.227a109.227 109.227 0 1 0 0-218.454 109.227 109.227 0 0 0 0 218.454zm0 72.817a182.044 182.044 0 1 1 0-364.088 182.044 182.044 0 0 1 0 364.088z"></path></svg></el-tag>
                      <el-tag v-if="data.expand === false" type="info" size="small" @click="do_expand_line(data, true)"> <svg class="icon" width="13" height="13" viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg" data-v-78e17ca8=""><path fill="currentColor" d="M329.956 257.138a254.862 254.862 0 0 0 0 509.724h364.088a254.862 254.862 0 0 0 0-509.724H329.956zm0-72.818h364.088a327.68 327.68 0 1 1 0 655.36H329.956a327.68 327.68 0 1 1 0-655.36z"></path><path fill="currentColor" d="M329.956 621.227a109.227 109.227 0 1 0 0-218.454 109.227 109.227 0 0 0 0 218.454zm0 72.817a182.044 182.044 0 1 1 0-364.088 182.044 182.044 0 0 1 0 364.088z"></path></svg></el-tag>
                     </span>
                    </el-tooltip>
                    <el-tooltip :content="view_line_title[data.id]" placement="right">
                     <span>
                      <el-tag v-if="data.view === true" size="small" @click="do_expand_line_view(data, false)"> <svg class="icon" width="13" height="13" viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg" data-v-78e17ca8=""><path fill="currentColor" d="M512 160c320 0 512 352 512 352S832 864 512 864 0 512 0 512s192-352 512-352zm0 64c-225.28 0-384.128 208.064-436.8 288 52.608 79.872 211.456 288 436.8 288 225.28 0 384.128-208.064 436.8-288-52.608-79.872-211.456-288-436.8-288zm0 64a224 224 0 1 1 0 448 224 224 0 0 1 0-448zm0 64a160.192 160.192 0 0 0-160 160c0 88.192 71.744 160 160 160s160-71.808 160-160-71.744-160-160-160z"></path></svg></el-tag>
                      <el-tag v-if="data.view === false" type="info" size="small" @click="do_expand_line_view(data, true)"> <svg class="icon" width="13" height="13" viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg" data-v-78e17ca8=""><path d="M876.8 156.8c0-9.6-3.2-16-9.6-22.4-6.4-6.4-12.8-9.6-22.4-9.6-9.6 0-16 3.2-22.4 9.6L736 220.8c-64-32-137.6-51.2-224-60.8-160 16-288 73.6-377.6 176C44.8 438.4 0 496 0 512s48 73.6 134.4 176c22.4 25.6 44.8 48 73.6 67.2l-86.4 89.6c-6.4 6.4-9.6 12.8-9.6 22.4 0 9.6 3.2 16 9.6 22.4 6.4 6.4 12.8 9.6 22.4 9.6 9.6 0 16-3.2 22.4-9.6l704-710.4c3.2-6.4 6.4-12.8 6.4-22.4Zm-646.4 528c-76.8-70.4-128-128-153.6-172.8 28.8-48 80-105.6 153.6-172.8C304 272 400 230.4 512 224c64 3.2 124.8 19.2 176 44.8l-54.4 54.4C598.4 300.8 560 288 512 288c-64 0-115.2 22.4-160 64s-64 96-64 160c0 48 12.8 89.6 35.2 124.8L256 707.2c-9.6-6.4-19.2-16-25.6-22.4Zm140.8-96c-12.8-22.4-19.2-48-19.2-76.8 0-44.8 16-83.2 48-112 32-28.8 67.2-48 112-48 28.8 0 54.4 6.4 73.6 19.2L371.2 588.8ZM889.599 336c-12.8-16-28.8-28.8-41.6-41.6l-48 48c73.6 67.2 124.8 124.8 150.4 169.6-28.8 48-80 105.6-153.6 172.8-73.6 67.2-172.8 108.8-284.8 115.2-51.2-3.2-99.2-12.8-140.8-28.8l-48 48c57.6 22.4 118.4 38.4 188.8 44.8 160-16 288-73.6 377.6-176C979.199 585.6 1024 528 1024 512s-48.001-73.6-134.401-176Z" fill="currentColor"></path><path d="M511.998 672c-12.8 0-25.6-3.2-38.4-6.4l-51.2 51.2c28.8 12.8 57.6 19.2 89.6 19.2 64 0 115.2-22.4 160-64 41.6-41.6 64-96 64-160 0-32-6.4-64-19.2-89.6l-51.2 51.2c3.2 12.8 6.4 25.6 6.4 38.4 0 44.8-16 83.2-48 112-32 28.8-67.2 48-112 48Z" fill="currentColor"></path></svg></el-tag>
                     </span>
                    </el-tooltip>
                    <el-tag size="small" type="danger" @click="removeLine(node, data)"> <svg class="icon" width="13" height="13" viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg" data-v-042ca774=""><path fill="currentColor" d="M352 192V95.936a32 32 0 0132-32h256a32 32 0 0132 32V192h256a32 32 0 110 64H96a32 32 0 010-64h256zm64 0h192v-64H416v64zM192 960a32 32 0 01-32-32V256h704v672a32 32 0 01-32 32H192zm224-192a32 32 0 0032-32V416a32 32 0 00-64 0v320a32 32 0 0032 32zm192 0a32 32 0 0032-32V416a32 32 0 00-64 0v320a32 32 0 0032 32z"></path></svg> </el-tag >
                  </span>
                </span>
              </template>
            </el-tree>
          </el-scrollbar>
        </div>
        <div class="box-card-left-body-action">
          <el-button size="small" type="success" icon="el-icon-baseball" class="button" @click="idcAppend">增加机房</el-button>
          <el-button size="small" type="info" icon="el-icon-refresh" plain class="button" @click="flushtree">刷新列表</el-button>
          <el-button v-if="pushing_all===false" icon="el-icon-upload" type="warning" plain size="small" class="button" 
            @click="updateAllIPAddrs">为所有IP打标签</el-button>
          <el-button v-if="pushing_all===true" icon="el-icon-loading" type="warning" plain size="small" class="button" 
            @click="updateAllIPAddrs">为所有IP打标签</el-button>
          <el-button v-if="pushing_part===false" icon="el-icon-upload" type="warning" plain size="small" class="button" 
            @click="updatePartIPAddrs">为没有标签IP打标签</el-button>
          <el-button v-if="pushing_part===true" icon="el-icon-loading"  type="warning" plain size="small" class="button" 
            @click="updatePartIPAddrs">只更新未设置IP</el-button>
          <el-button v-if="pushing_create_label===false" icon="el-icon-upload" type="info" plain size="small" class="button" 
            @click="doCreateLabelForAllIPs">JOB组中生成标签</el-button>
          <el-button v-if="pushing_create_label===true" icon="el-icon-loading" type="info" plain size="small" class="button" 
            @click="doCreateLabelForAllIPs">JOB组中生成标签</el-button>
          <el-button v-if="pushing_create_ips===false" icon="el-icon-upload" type="danger" plain size="small" class="button" 
            @click="doCreateIPForAllJob">扩展所有组地址</el-button>
          <el-button v-if="pushing_create_ips===true" icon="el-icon-loading" type="danger" plain size="small" class="button" 
            @click="doCreateIPForAllJob">扩展所有组地址</el-button>
        </div>
      </div>
    </el-card>
    <el-card class="box-card-right" shadow="never">
      <template #header>
        <div class="card-header">
          <span>
            当前选择机房：<el-tag v-if="currentIDCTitle!==''" size="mini" type="danger">{{currentIDCTitle}}</el-tag>
            线路：<el-tag v-if="currentLineTitle!==''" size="mini" type="danger">{{currentLineTitle}}</el-tag>
          </span>
          <el-button size="small" type="success" class="button" icon="el-icon-upload2" @click="putIDCOrLineIPData"> 更 新 </el-button>
        </div>
      </template>
      <div class="card-body">
        <el-form
          label-position="top"
          :model="linedetailinfo"
          size="mini" 
        >
          <el-form-item label="IP列表：（以英文分号(;)隔开；如:192.168.1.0/24;10.10.10.1;172.16.1.1~172.16.2.1）">
            <template #label>
              <div class="display-title">
                <span>IP列表：（以英文分号(;)隔开；如:192.168.1.0/24;10.10.10.1;172.16.1.1~172.16.2.1）</span>
                <el-button size="mini" plain type="danger" @click="expandIPAddr">扩展当前组选择项IP</el-button>
              </div>
            </template>
            <el-input :disabled="should_disabled || onobject_disabled" resize="none" :rows="16" type="textarea" placeholder="" v-model="linedetailinfo.ipaddrs_net_line"></el-input>
          </el-form-item>
          <el-form-item label="备注信息：">
            <el-input :disabled="onobject_disabled" resize="none" :rows="15" type="textarea" placeholder="" v-model="linedetailinfo.remark_info"></el-input>
          </el-form-item>
        </el-form>
      </div>
    </el-card>
    <el-dialog v-model="dialogLineVisible" :title="idcName" width="450px">
      <el-form :model="Lineform" :rules="lineRules" ref="Lineform">
        <el-form-item label="线路名称" :label-width="formLabelWidth" prop="label">
          <el-input size="small" v-model="Lineform.label" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="是否扩展地址" :label-width="formLabelWidth" prop="expand">
          <el-switch v-model="Lineform.expand" class="ml-2" active-color="#13ce66" inactive-color="#ff4949"/>
        </el-form-item>
        <el-form-item label="地址是否可见" :label-width="formLabelWidth" prop="view">
          <el-switch v-model="Lineform.view" class="ml-2" active-color="#13ce66" inactive-color="#ff4949"/>
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
    <el-dialog v-model="dialogIDCVisible" :title="dialogIDCTitle" width="450px">
      <el-form :model="IDCform" :rules="IDCRules" ref="IDCform">
        <el-form-item label="机房名称" :label-width="formLabelWidth" prop="label">
          <el-input size="small" v-model="IDCform.label" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="是否扩展地址" :label-width="formLabelWidth" prop="expand">
          <el-switch v-model="IDCform.expand" class="ml-2" active-color="#13ce66" inactive-color="#ff4949"/>
        </el-form-item>
        <el-form-item label="地址是否可见" :label-width="formLabelWidth" prop="view">
          <el-switch v-model="IDCform.view" class="ml-2" active-color="#13ce66" inactive-color="#ff4949"/>
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
    <el-dialog v-model="delLineDialogVisable" title="确认信息" width="430px" :before-close="cancelRemoveLine">
      <div class="confirm_rm_line"><span>是否确定删除？</span></div>
      <div class="confirm_rm_line_cb"><el-checkbox v-model="delIpAddrsForLine">在“IP管理”同步删除此线路的所有地址</el-checkbox></div>
      <template #footer>
        <span class="dialog-footer">
          <el-button size="small" @click="cancelRemoveLine">放弃</el-button>
          <el-button v-if="delIpAddrsForLineDoing === false" size="small" type="primary" @click="confirmRemoveLine">确定</el-button>
          <el-button v-if="delIpAddrsForLineDoing === true" icon="el-icon-loading" size="small" type="primary">提交中</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script>
  let tree_id = 1000
  // import { getIDC, getIDCs, postIDC, putIDC, deleteIDC } from '@/api/idc.js'
  import { getIDC, postIDC, putIDC, putIDCRemark, deleteIDC, putIdcExpand, putLineExpand, putIdcView, putLineView } from '@/api/idc.js'
  import { getIDCTree, outportXls } from '@/api/idc.js'
  import { postLine, putLine, delLine } from '@/api/idc.js'
  import { updateAllIPAddrsNetInfo, updatePartIPAddrsNetInfo } from '@/api/idc.js'
  // import { getLine, getLines, postLine, putLine, delLine } from '@/api/idc.js'
  import { getLineIpAddrs, putLineIpAddrs, createLabelForAllIPs, createIPForJob } from '@/api/idc.js'

  import { Base64 } from 'js-base64';

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
          expand: false,
          view: true,
          children: [],
        },
        IDCform: {
          id: 0,
          tree_id: 0,
          tree_type: 'idc',
          label: '',
          expand: false,
          view: true,
          children: [],
        },
        lineRules: {
          label: [
            { required: true, validator: validateStr, trigger: ['blur'] }
          ],
        },
        IDCRules: {
          label: [
            { required: true, validator: validateStr, trigger: ['blur'] }
          ],
        },
        currentDataPoint: null,
        formLabelWidth: '100px',
        idcName: '添加线路',
        idEdit: false,
        currentPoolObj: null,
        currentTitle: '',
        currentIDCTitle: '',
        currentLineTitle: '',
        pushing_all: false,
        pushing_part: false,
        pushing_create_label: false,
        pushing_create_ips: false,
        should_disabled: false,
        searchContent: '',
        search_ip: false,
        onobject_disabled: false,
        defaultExpandedKeysStore: new Set(),
        defaultExpandedKeys: [],
        delLineDialogVisable: false,
        delLineCurrNode: null,
        delLineCurrData: null,
        delIpAddrsForLine: false,
        delIpAddrsForLineDoing: false,
        expand_idc_title: {},
        expand_line_title: {},
        view_idc_title: {},
        view_line_title: {}
      }
    },

    mounted () {
      this.doGetTree()
      this.onobject_disabled = true
    },

    methods: {
      doGetTree(show_notice){
        getIDCTree({content: this.searchContent, search_ip: this.search_ip}).then(r => {
          r.data.tree.forEach(each => {
            if (each.expand === true) {
              this.expand_idc_title[each.id] = '扩展'
            } else {
              this.expand_idc_title[each.id] = '不扩展，包括所有线路'
            }
            if (each.view === true) {
              this.view_idc_title[each.id] = '可见'
            } else {
              this.view_idc_title[each.id] = '不可见，包括所有线路'
            }
            each.children.forEach(child => {
              if(child.expand === true) {
                this.expand_line_title[child.id] = '扩展'
              } else {
                this.expand_line_title[child.id] = '不扩展'
              }
              if (child.view === true) {
                this.view_line_title[child.id] = '可见'
              } else {
                this.view_line_title[child.id] = '不可见'
              }
            })
          })
          this.data = r.data.tree
          this.defaultExpandedKeys = Array.from(this.defaultExpandedKeysStore)
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
          expand: false,
          view: true,
          children: [],
        },
        this.idEdit = false
        this.dialogLineVisible = true
      },
      lineAppend(formName){
        this.$refs[formName].validate((valid) => {
          if (valid) {
            if (this.idEdit === false) {
              postLine({
                label: this.Lineform.label, 
                idc_id: this.Lineform.id, 
                expand: this.Lineform.expand,
                view: this.Lineform.view,
              }).then(r=>{
                this.doGetTree()
                this.dialogLineVisible = false
              }).catch(e=>{
                this.dialogLineVisible = false
                console.log(e)
              })
            } else {
              putLine({
                label: this.Lineform.label, 
                id: this.Lineform.id, 
                expand: this.Lineform.expand,
                view: this.Lineform.view,
              }).then(r=>{
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
              postIDC({
                'label': this.IDCform.label, 
                expand: this.IDCform.expand,
                view: this.IDCform.view,
              }).then(r=>{
                this.doGetTree()
                this.dialogIDCVisible = false
              }).catch(e=> {
                console.log(e)
                this.dialogIDCVisible = false
              })
            } else {
              putIDC({
                'label': this.IDCform.label, 
                'id': this.IDCform.id, 
                expand: this.IDCform.expand,
                view: this.IDCform.view,
              }).then(r=>{
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
      do_expand_idc(data, v){
        putIdcExpand({id: data.id, switch: v}).then(
          r => {
            data.expand = v
            if (v===true) {
              this.expand_idc_title[data.id] = '扩展地址'
            } else {
              this.expand_idc_title[data.id] = '不扩展地址，包括所有线路'
            }
          }
        ).catch(e=>console.log(e))
      },
      do_expand_line(data, v){
        putLineExpand({id: data.id, switch: v}).then(r=>{
          data.expand = v
          if (v===true) {
            this.expand_line_title[data.id] = '扩展地址'
          } else {
            this.expand_line_title[data.id] = '不扩展地址'
          }
        }).catch(e=>console.log(e))
      },
      do_expand_idc_view(data, v) {
        putIdcView({id: data.id, switch: v}).then(
          r => {
            data.view = v
            if (v===true) {
              this.view_idc_title[data.id] = '可见'
            } else {
              this.view_idc_title[data.id] = '不可见，包括所有线路'
            }
          }
        ).catch(e=>console.log(e))
      },
      do_expand_line_view(data, v) {
        putLineView({id: data.id, switch: v}).then(r=>{
          data.view = v
          if (v===true) {
            this.view_line_title[data.id] = '可见'
          } else {
            this.view_line_title[data.id] = '不可见'
          }
        }).catch(e=>console.log(e))
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
        this.delLineDialogVisable = true
        // this.delLineCurrNode = node
        this.delLineCurrData = data
        // this.$confirm('是否确定删除？', '确认信息', {
        //   distinguishCancelAndClose: true,
        //   confirmButtonText: '确定',
        //   cancelButtonText: '放弃'
        // }).then(_ => {
        //   delLine({id: data.id}).then(r=>{
        //   this.$notify({
        //     title: '成功',
        //     message: '删除成功！',
        //     type: 'success'
        //   });
        //     this.doGetTree()
        //   }).catch(e=>{
        //     console.log(e)
        //   })
        // }).catch(e => console.log(e))
      },
      confirmRemoveLine() {
        let data = this.delLineCurrData
        this.delIpAddrsForLineDoing = true
        delLine({id: data.id, rm_addrs: this.delIpAddrsForLine}).then(r=>{
          this.$notify({
            title: '成功',
            message: '删除成功！',
            type: 'success'
          });
          this.delLineDialogVisable = false
          this.delIpAddrsForLineDoing = false
          this.doGetTree()
        }).catch(e=>{
          this.delLineDialogVisable = false
          this.delIpAddrsForLineDoing = false
          console.log(e)
        })
      },
      cancelRemoveLine(){
        this.delLineDialogVisable = false
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
        this.idcName = '编辑线路：' + data.label
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
        this.onobject_disabled = false
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
      nodeExpand(data,node,ele){
        this.defaultExpandedKeysStore.add(data.tree_id)
        // console.log('nodeExpand', data,node,ele)
      },
      nodeCollapse(data,node,ele){
        this.defaultExpandedKeysStore.delete(data.tree_id)
        // console.log('nodeCollapse', data,node,ele)
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
      },
      expandIPAddr(){
        if (!this.currentPoolObj) {
          return 
        }
        let expandData = {'idc': [], 'line': [], all: false}
        let explain = ''
        if (this.currentPoolObj.tree_type === 'idc') {
          expandData.idc.push(this.currentPoolObj.id)
          explain = '是否确定扩展IDC: "' + this.currentPoolObj.label + '"下所有线路地址？'
        } else if (this.currentPoolObj.tree_type === 'line') {
          expandData.line.push(this.currentPoolObj.id)
          explain = '是否确定扩展Line: "' + this.currentPoolObj.label + '"线路地址？'
        }
        this.$confirm(explain, '确认信息', {
          distinguishCancelAndClose: true,
          confirmButtonText: '确定',
          cancelButtonText: '放弃'
        }).then(_ => {
          this.doCreateIPForJob(expandData)
        }).catch(e => console.log(e))
      },
      doCreateIPForAllJob(){
        this.$confirm('是否确定扩展所有IDC下所有线路地址？', '确认信息', {
          distinguishCancelAndClose: true,
          confirmButtonText: '确定',
          cancelButtonText: '放弃'
        }).then(_ => {
          this.doCreateIPForJob({'idc': [], 'line': [], all: true})
        }).catch(e => console.log(e))
      },
      doCreateIPForJob(data){
        createIPForJob(data).then(r=>{
          this.$notify({
            title: '成功',
            message: '扩展成功！',
            type: 'success'
          });
        }).catch(e=>{
          console.log(e)
        })
      },
      doOutportXls(){
        outportXls().then(r=>{
          let content = Base64.toUint8Array(r.data.data)
          let filename = r.data.name  //导出的文件名
          let type = "application/vnd.ms-excel";                      //头部数据类型
          let file = new Blob([content], { type: type });
          if (window.navigator.msSaveOrOpenBlob)
            // IE10+
            window.navigator.msSaveOrOpenBlob(file, filename);
          else {
            // Others
            let a = document.createElement("a"),
            url = URL.createObjectURL(file);
            a.href = url;
            a.download = filename;
            document.body.appendChild(a);
            a.click();
            setTimeout(function() {
              document.body.removeChild(a);
              window.URL.revokeObjectURL(url);
            }, 0);
          }
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
  height: 770px;
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
.box-card-left-body{
  display: flex;
  flex-wrap: nowrap;
}
.box-card-left-body-content {
  border-left: 5px solid #ccc!important;border-color: #DDDDDD!important;
  width: 85%;
  height: 670px;;
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
.box-card-left {
  width: 45%;
}
.box-card-right {
  width: 55%;
}

.card-body {
  margin-top: -15px;
}
.card-scrollbar {
  max-height: 660px;
}

.box-board :deep() .el-card__body {
  padding-bottom: 0px;
}
.display-title {
  display: flex;
  justify-content: space-between;
}
.confirm_rm_line {
  margin-top: -20px;
  margin-bottom: 10px;
  margin-left: 30px;
}
.confirm_rm_line_cb {
  margin-bottom: -20px;
  margin-left: 30px;
}
.download-btn {
  margin-left: 10px;
}
</style>