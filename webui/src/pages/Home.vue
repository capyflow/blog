<template>
  <div class="home-page">
  <div class="grid home-grid">
    <section class="section left">
      <h3>Blog Posts</h3>
      <CardList :items="lists.article" />
    </section>
    <section class="section right-top">
      <h3>Projects</h3>
      <CardList :items="lists.project" />
    </section>
    <section class="section right-bottom">
      <h3>Miscellaneous Talk</h3>
      <CardList :items="lists.entertain" />
    </section>
  </div>
  </div>
</template>

<script setup lang="ts">
import CardList from '../components/CardList.vue'

const lists = {
  article: mock('article', 12),
  project: mock('project', 8),
  entertain: mock('entertain', 8),
}

function mock(type:string, n:number){
  return Array.from({length:n}).map((_,i)=>({
    id: `${type}-${i+1}`,
    date: '2025-11-30',
    title: `${type} title ${i+1}`,
  }))
}
</script>
<style scoped>
.home-page{padding:8px}
.home-grid{grid-template-columns:2fr 1fr; grid-template-rows:auto auto; gap:16px}
.left{grid-column:1 / 2; grid-row:1 / 3}
.right-top{grid-column:2 / 3; grid-row:1 / 2}
.right-bottom{grid-column:2 / 3; grid-row:2 / 3}
</style>
