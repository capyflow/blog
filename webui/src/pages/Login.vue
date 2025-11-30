<template>
  <div style="max-width:360px;margin:40px auto" class="card">
    <h3 style="margin:0 0 12px">控制台登录</h3>
    <form @submit.prevent="onSubmit" class="grid" style="grid-template-columns:1fr;gap:12px">
      <label>
        <div class="muted">Email</div>
        <input v-model="email" type="email" required />
      </label>
      <label>
        <div class="muted">Password</div>
        <input v-model="password" type="password" required />
      </label>
      <button type="submit">登录</button>
      <div v-if="err" class="muted">{{ err }}</div>
    </form>
  </div>
</template>

<script setup lang="ts">
import { loginByPwd } from '../api/console'
import { useAuth } from '../store/auth'
import { useRouter } from 'vue-router'
import { ref } from 'vue'

const email = ref('')
const password = ref('')
const err = ref('')
const auth = useAuth()
const router = useRouter()

async function onSubmit(){
  err.value = ''
  try{
    const { data } = await loginByPwd(email.value, password.value)
    const token = (data?.body?.token) || ''
    if(!token) throw new Error('无效的登录响应')
    auth.setToken(token)
    router.push('/')
  }catch(e:any){
    err.value = e?.message || '登录失败'
  }
}
</script>
