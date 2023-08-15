<template>
  <el-table :data="tableData" class="table">
    <el-table-column prop="Name" label="文件名" >
      <template #default="scope">
        <el-link :href="IP+'/uploads/'+scope.row.Name"  target="_blank" type="primary"> {{scope.row.Name}}</el-link>

      </template>



    </el-table-column>
    <el-table-column prop="Date" label="日期" />
    <el-table-column prop="Size" label="大小(MB)" width="110" />

  </el-table>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      tableData:[],
      intervalId: null,
      IP:""
    };
  },
  created() {
    this.IP=this.$IP
    this.fetchData()
    this.forGet()

  },
  components:{

  },

  methods: {
    forGet(){
      setInterval( this.fetchData , 2000);
    },


    fetchData() {
      axios.get(this.$IP+'/getFileList')
          .then(response => {
            // 请求成功，处理响应数据
            console.log(response.data);
           this.tableData= response.data.FileList
          })
          .catch(error => {
            // 请求发生错误，处理错误信息
            console.error(error);
          });
    },
  },



}
</script>


<style>
.table{
  width: auto;
}


</style>