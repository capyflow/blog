<template>
  <div class="list">
    <div v-for="(item,i) in items" :key="i" class="list-row">
      <span class="date">{{ item.date }}</span>
      <span class="sep">—</span>
      <router-link :to="`/article/${item.id}`" class="title link">{{ item.title }}</router-link>
    </div>
  </div>
</template>

<script setup lang="ts">
const props = defineProps<{ items: Array<{id:string;date:string;title:string}> }>()
</script>

<style scoped>
.list{display:block}
.list-row{padding:8px 6px}
.link{color:inherit; text-decoration:none}
.link:hover{color:#4da3ff; text-decoration:underline}
.date{color:#666; margin-right:8px}
.sep{color:#999; margin-right:8px}
.title{color:#fff}
</style>
