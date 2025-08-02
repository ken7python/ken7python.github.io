<template>
  <Header />
  <div class="blog-container">
    <div class="content">
      <h1 class="title">KenCodeブログ</h1>
      <p class="description">※このブログはRSSを利用してnoteと連携しています</p>

      <section class="latest-posts">
        <h2 class="section-title">最新のnote投稿</h2>
        <ul class="post-list">
          <li v-for="article in articles" :key="article.link" class="post-item">
            <a :href="article.link" target="_blank" class="post-link">{{ article.title }}</a>
            <small class="post-date">({{ new Date(article.pubDate).toLocaleDateString() }})</small>
          </li>
        </ul>
      </section>
    </div>
  </div>
</template>

<script setup lang="ts">
import Header from '../components/Header.vue';
import { ref, onMounted } from "vue";

interface Article {
  title: string;
  link: string;
  pubDate: string;
}

const articles = ref<Article[]>([]);

async function fetchAndParse(): Promise<void> {
  const rssURL = 'https://note.com/ken7python/rss';
  const apiURL = `https://api.allorigins.win/get?url=${rssURL}`;
  try {
    const res = await fetch(apiURL);
    if (!res.ok) throw new Error('Network response was not ok');
    const data = await res.json();
    const xml = data.contents;

    const doc = new DOMParser().parseFromString(xml, 'application/xml');
    const items = Array.from(doc.querySelectorAll('item'));

    articles.value = items.map(item => ({
      title: item.querySelector('title')?.textContent ?? '(タイトルなし)',
      link: item.querySelector('link')?.textContent ?? '#',
      pubDate: item.querySelector('pubDate')?.textContent ?? '',
    })) as Article[];
  } catch (err) {
    console.error(err);
  }
}

onMounted(fetchAndParse);
</script>

<style scoped>
.blog-container {
  padding: 20px;
  background-color: #f9f9f9;
  font-family: Arial, sans-serif;
}

.content {
  max-width: 800px;
  margin: 0 auto;
}

.title {
  text-align: center;
  color: #333;
  font-size: 2rem;
  margin-bottom: 10px;
}

.description {
  text-align: center;
  color: #666;
  margin-bottom: 20px;
}

.latest-posts {
  background: #fff;
  padding: 15px;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.section-title {
  font-size: 1.5rem;
  color: #444;
  margin-bottom: 10px;
}

.post-list {
  list-style: none;
  padding: 0;
}

.post-item {
  margin-bottom: 10px;
}

.post-link {
  color: #1a73e8;
  text-decoration: none;
}

.post-link:hover {
  text-decoration: underline;
}

.post-date {
  color: #888;
  font-size: 0.9rem;
}
</style>
