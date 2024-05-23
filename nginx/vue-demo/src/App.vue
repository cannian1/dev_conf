<script setup>
import {ref} from 'vue'
import axios from 'axios'

let data = ref(0)
let serverPort = ref()

const url = "http://10.1.30.123:8080" // 服务端地址
const getRandomNum = () => {
  axios.get(url + "/api/v1/ccc/j").then(function (response) {
        console.log(response.data)
        data.value = response.data.message
        serverPort.value = response.data.whoami
      }
  )
// console.log(data)
}
</script>

<template>
  <header>

  </header>

  <h2>你好啊</h2>
  <button @click="getRandomNum">获取随机数</button>
  <br>
  <p>接收来自 {{ serverPort }} 的响应</p>
  <p>接收到后端返回的随机数：{{ data }}</p>
</template>

