<script setup>
import Upload from './components/Upload.vue'
import FileTable from "./components/FileTable.vue";
import HeadQRCode from "./components/HeadQRCode.vue";

</script>

<template>
<!--  <img id="logo" alt="Wails logo" src="./assets/images/logo-universal.png"/>-->
<!--  <HelloWorld/>-->
  <div class="head-qr-code">
    <head-q-r-code/>
  </div>
  <div style="margin: 10px"></div>
  <div class="upload">
    <Upload/>
  </div>
  <div style="margin: 10px"></div>
  <FileTable/>
</template>

<style>
#logo {
  display: block;
  width: 50%;
  height: 50%;
  margin: auto;
  padding: 10% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}
.upload{
  width: 70%;
  margin: auto;
}
</style>
