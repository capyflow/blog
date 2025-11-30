import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router'
import './styles/sketch.css'

const app = createApp(App)
app.use(createPinia())
app.use(router)
app.mount('#app')
document.body.classList.add('bg-dynamic')
