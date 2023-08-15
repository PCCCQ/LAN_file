<template>
  <el-upload
      drag
      :action="$IP+'/uploadFile'"
      multiple
  >
    <el-icon class="el-icon--upload"><upload-filled /></el-icon>
    <div class="el-upload__text">
      将文件拖到此处 或 <em>单击上传</em>
    </div>
<!--    <template #tip>-->
      <div class="el-upload__tip">
        如需批量上传，请使用浏览器打开，微信只支持单个文件上传
      </div>
<!--    </template>-->
  </el-upload>
</template>

<script setup lang="ts">
import { UploadFilled } from '@element-plus/icons-vue'
</script>


<style>



</style>