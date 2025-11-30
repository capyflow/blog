import { createRouter, createWebHistory } from 'vue-router'
import { useAuth } from '../store/auth'
import Home from '../pages/Home.vue'
import Article from '../pages/Article.vue'
import Login from '../pages/Login.vue'
import Workbench from '../pages/Workbench.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', component: Home },
    { path: '/article/:id', component: Article },
    { path: '/login', component: Login },
    { path: '/workbench', component: Workbench }
  ]
})

router.beforeEach((to) => {
  const auth = useAuth()
  if (to.path === '/login') return true
  if (to.path === '/workbench' && !auth.token) return '/login'
  return true
})

export default router
