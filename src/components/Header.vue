<template>
  <header>
    <div id="menuBar">
      <div id="linkHome" class="menu-item">
        <RouterLink class="link" to="/" :style="routeName === 'Home' ? { fontWeight: 'bold' } : {}"><ion-icon name="home-outline"></ion-icon>Home</RouterLink>
        <div class="underLine" v-if="routeName === 'Home'"></div>
      </div>
      <div id="linkBlog" class="menu-item">
        <RouterLink class="link" to="/blog" :style="routeName === 'Blog' ? { fontWeight: 'bold' } : {}"><ion-icon name="book-outline"></ion-icon>Blog</RouterLink>
        <div class="underLine" v-if="routeName === 'Blog'"></div>
      </div>
      <div id="linkGame" class="menu-item">
        <RouterLink class="link" to="/game" :style="routeName === 'Game' ? { fontWeight: 'bold' } : {}"><ion-icon name="game-controller-outline"></ion-icon>Game</RouterLink>
        <div class="underLine" v-if="routeName === 'Game'"></div>
      </div>
      <div id="linkRepos" class="menu-item">
        <RouterLink class="link" to="/repos" :style="routeName === 'Repos' ? { fontWeight: 'bold' } : {}"><ion-icon name="folder-outline"></ion-icon>Repository</RouterLink>
        <div class="underLine" v-if="routeName === 'Repos'"></div>
      </div>
    </div>
	</header>
</template>

<script setup lang="ts">
	import { ref } from 'vue';
  import { useRoute } from 'vue-router';

  const route = useRoute();
  const routeName = ref(route.name);
</script>

<style scoped>
	header {
    width: 100svw;
    background-color: #ffffff;
    border-bottom: 1px solid #dfe5ea;
    z-index: 1000;

    background-color: #f6f8fa;
    padding-left: 2vw;
	}
  #menuBar {
    padding-top: 50px;

    display: flex;
    flex-direction: row;
    /* justify-content: flex-end; */
    flex-direction: row;
    justify-content: start;

    width: 98vw;
  }
  ion-icon {
    position: relative;
    font-size: 16px;
    top: 2px; /* アイコンを少し上に */
    margin-right: 8px; /* アイコンとテキストの間にスペースを追加 */
  }

	.link{
      font-size: 14px;
			text-decoration: none;
			text-align: center;
			border-radius: 4px;
			color: #1f2329;
			/* font-weight: bold; */
			border: none;

      padding: 8px 16px;
      margin: 8px 0;
	}

  .underLine {
    width: 100%;
    height: 2px;
    background-color: #fd8c73;
    margin-top: 8px;
  }

  .menu-item {
    margin: 0;
    margin-bottom: 0;
  }

  .header {
    padding-bottom: 0;
  }

  .menu-item .link {
    display: flex;
    align-items: flex-end;
  }
</style>