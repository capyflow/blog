<template>
  <div class="layout">
    <header v-if="!isArticle" class="topbar">
      <div class="brand">Aaron的博客</div>
      <nav class="nav">
        <router-link v-if="hasToken" to="/workbench" class="btn">工作台</router-link>
      </nav>
    </header>
    <main :class="[isArticle ? 'container-article' : 'container']">
      <router-view />
    </main>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useAuth } from './store/auth'
import { useRoute } from 'vue-router'
const auth = useAuth()
const hasToken = computed(()=>!!auth.token)
const route = useRoute()
const isArticle = computed(()=> route.path.startsWith('/article/'))
</script>

<style scoped>
.layout { min-height: 100vh; display: flex; flex-direction: column; }
.topbar { display:flex; align-items:center; justify-content:space-between; padding:16px 24px; }
.brand { font-weight:700; }
.nav .btn { padding:6px 12px; color:#e6e6e6; text-decoration:none; }
.container { padding:24px; flex:1 }
.container-article { padding:0; }
.home-bg { border-radius:12px; padding:16px }
.bg-dynamic { }
</style>
