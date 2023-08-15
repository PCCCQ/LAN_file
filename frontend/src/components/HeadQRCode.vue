<template>
  <el-card>
<!--    <div v-for="o in 4" :key="o" class="text item">{{ 'List item ' + o }}</div>-->
<!--    <el-form :inline="true">-->
    <div>
    <el-select v-model="value" placeholder="选择局域网IP">
      <el-option
          v-for="item in options"
          :key="item.ip"
          :label="item.ip"
          :value="item.ip"
      />
    </el-select>
    </div>
<!--      <el-button style="margin-left:30px" @click="generateQRCode">显示二维码</el-button>-->
    <div v-if="value">
      <qrcode  :value="'http://'+value+':9999/index'"  />
<!--     </el-form>-->
      <div>
      {{'http://'+value+':9999/index'}}
    </div>
    </div>
  </el-card>
</template>
<script>
import axios from "axios";
import VueQrcode from "vue-qrcode";
export default {
  data() {
    return {
      options:[],
      value:""
    };
  },
  created() {

    this.fetchData();

  },
components:{
  // 注册 vue-qrcode 组件
  qrcode: VueQrcode,
},

  methods: {
    fetchData() {
      axios.get(this.$IP+'/getIP')
          .then(response => {
            // 请求成功，处理响应数据
            console.log(response.data);
            response.data.IPList.forEach((item)=>{
             if(item!==""){
               this.options.push({ip:item})
             }
            })
            console.log(this.options);
          })
          .catch(error => {
            // 请求发生错误，处理错误信息
            console.error(error);
          });
    },
  },



}
</script>
<style scoped>
</style>



