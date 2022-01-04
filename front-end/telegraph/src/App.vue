<template>
  <img alt="Vue logo" src="./assets/logo.png">
  <HelloWorld msg="Welcome to Your Vue.js App"/>
  <p v-bind="info"></p>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'
import axios from "axios";

export default {
  name: 'App',
  data() {
    return{
      info: ""
    }
  },
  components: {
    HelloWorld
  },
  mounted() {
    axios
      .get('http://localhost:5000/api/', {
        headers: {
          'Access-Control-Allow-Origin': '*',
        }
      })
      .then(response => (
          this.info = response
      ))
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
