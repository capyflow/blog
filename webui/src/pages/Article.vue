<template>
  <div>
    <div class="hero hero-bg full" :style="{ backgroundImage: `url(${bg})` }">
      <div class="overlay"></div>
      <button class="back" @click="$router.back()">←</button>
      <div class="title page-title">{{ title }}</div>
      <div class="muted date">{{ date }}</div>
    </div>
    <div class="content md" v-html="mdHtml"></div>
  </div>
</template>

<script setup lang="ts">
import { useRoute } from 'vue-router'
import * as markedNs from 'marked'
const route = useRoute()
const title = 'title'
const date = '2025-11-30'
const bg = 'https://images.unsplash.com/photo-1484591974057-265bb767ef71?q=80&w=1200&auto=format&fit=crop'
const md = `# Markdown 示例\n\n这里是正文内容，支持**加粗**、列表等。\n\n- 条目一\n- 条目二\n- 条目三 \n ![图片描述](https://images.unsplash.com/photo-1484591974057-265bb767ef71?q=80&w=1200&auto=format&fit=crop)`
const mdHtml = markedNs.parse(md)

</script>

<style scoped>
.hero-bg{background-size:cover;background-position:center;position:relative}
.full{width:100%; margin:0; border-radius:0;height:480px}
.overlay{position:absolute;inset:0;background:rgba(65, 65, 65, 0.35)}
.page-title{position:absolute;left:50%;top:50%;transform:translate(-50%,-50%);font-size:28px;font-weight:700;color:#ffffff}
.date{position:absolute;right:12px;bottom:12px;color:rgba(255, 255, 255, 0.77)}
.md :where(h1,h2,h3){margin:20px 0 12px}
.md p{line-height:1.8}
.md ul{padding-left:20px}
:deep(.md img),:deep(.md video){display:block;max-width:100%;margin:12px 0;border-radius:8px}
:deep(.md video){max-height:480px;width:100%}
.content {border-radius:10px}
</style>
