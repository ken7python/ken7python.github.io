<template>
  <div>
    <Header />
		<div id="container">
			<h1>KenCodeブログ</h1>
      <p>※このブログはRSSを利用してnoteと連携しています</p>

      <section>
        <h2>最新のnote投稿</h2>

        <ul>
          <li v-for="article in articles" :key="article.link">
            <a :href="article.link" target="_blank">{{ article.title }}</a>
            <small> ({{ new Date(article.pubDate).toLocaleDateString() }})</small>
          </li>  
        </ul>
      </section>
    </div>
  </div>
</template>

<script setup lang="ts">
	import Header from '../components/Header.vue';
  import { ref,onMounted } from "vue";
  interface Article {
    title: string
    link: string
    pubDate: string
  }

  const articles = ref<Article[]>([])
  
  async function fetchAndParse() {
    const rssURL = 'https://note.com/ken7python/rss'
    const apiURL = `https://api.allorigins.win/get?url=${rssURL}`
    try {
      const res = await fetch(apiURL)
      if (!res.ok) throw new Error('Network response was not ok')
      const data = await res.json();
      const xml = data.contents;

      const doc = new DOMParser().parseFromString(xml, 'application/xml')

      const items = Array.from(doc.querySelectorAll('item'));

      console.log(items);

      articles.value = items.map(item => ({
        title: item.querySelector('title')?.textContent ?? '(タイトルなし)',
        link: item.querySelector('link')?.textContent ?? '#',
        pubDate: item.querySelector('pubDate')?.textContent ?? '',
      })) as Article[];

    } catch (err) {
      console.error(err);
    }
  }

  onMounted(fetchAndParse)

</script>
