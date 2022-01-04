import { createApp } from 'vue'
import App from './App.vue'
import axios from "axios";

axios.defaults.headers.get['Access-Control-Allow-Origin'] = '*';

createApp(App).mount('#app')
