import {createApp} from 'vue'
import App from './App.vue'
// 定义全局常量
const IP = "http://127.0.0.1:9999"
// const IP = ""
// 创建Vue应用
const app = createApp(App);

// 在应用程序上下文中设置全局常量
app.config.globalProperties.$IP = IP

// 挂载到DOM元素上
app.mount('#app');
