<template>
  <header>
    <div id="AlwaysMenu">
      <button id="MenuToggle" class="menu-button" @click="toggleMenu">...</button>
      <div id="OriginalMenu">
        <div id="linkHome">
          <RouterLink class="link" to="/">ホーム</RouterLink>
          <div class="underLine" v-if="routeName === 'Home'"></div>
        </div>
        <div id="linkBlog">
          <RouterLink class="link" to="/blog">ブログ</RouterLink>
          <div class="underLine" v-if="routeName === 'Blog'"></div>
        </div>
        <div id="linkGame">
          <RouterLink class="link" to="/game">ゲーム</RouterLink>
          <div class="underLine" v-if="routeName === 'Game'"></div>
        </div>
			</div>
		</div>
		<div id="SNSMenu" :class="{ 'open': menuOpen }">
			<a class="link" id="linkGitHub" target="_blank" rel="noopener noreferrer" href="https://github.com/ken7python">GitHub</a>
			<a class="link" id="linkX" target="_blank" rel="noopener noreferrer" href="https://x.com/ken7python">X</a>
			<a class="link" id="linkInstagram" target="_blank" rel="noopener noreferrer" href="https://www.instagram.com/ken7python/">Instagram</a>
			<a class="link" id="linkNote" target="_blank" rel="noopener noreferrer" href="https://note.com/ken7python">note</a>
			<a class="link" id="linkYouTube" target="_blank" rel="noopener noreferrer" href="https://www.youtube.com/@KenCodeHanabi">Youtube</a>
		</div>
	</header>
</template>

<script setup lang="ts">
	import { ref } from 'vue';
  import { useRoute } from 'vue-router';

  const menuOpen = ref(false);

  function toggleMenu() {
    menuOpen.value = !menuOpen.value;
  }

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
	}
	.menu-button {
    display: none;
    background: none;
    border: none;
    font-size: 24px;
    padding: 8px;
	}

	#AlwaysMenu {
			display: flex;
			justify-content: space-between;
			align-items: center;
			padding: 8px 16px;
      padding-right: 8vw;
	}

  #OriginalMenu {
    display: flex;
    flex-direction: row;
    /* gap: 16px; */
  }

	@media (max-width: 1024px) { /* タブレット・スマホ向けデザイン */
		#SNSMenu.open {
      display: flex;
    }
	}

	/* リンクたち */
	@media (max-width: 670px) { /* スマホ向けデザイン */ /* スマホでは「…」ボタンを表示 */
			.menu-button {
					display: block;      /* 「…」ボタンを表示 */
			}
			#SNSMenu {
					display: none;    /* 初期状態ではリンクを隠す */
					flex-direction: column;
					padding: 0 16px;
			}
			#SNSMenu.open {
					display: flex;       /* JS で「open」を付けると表示 */
			}
	}

	@media (min-width: 671px) { /* タブレット・PC向けデザイン */
			header {
					display: flex;
					flex-direction: row-reverse;
					justify-content: space-between;
					align-items: center;
					padding: 8px 16px;
			}
	}


	.link{
			text-decoration: none;
			text-align: center;

			padding: 8px 16px;
			margin: 8px 0;
			border-radius: 4px;
			color: #1f2329;
			font-weight: bold;
			font-size: 14px;
			border: none;
	}

  .underLine {
    width: 100%;
    height: 2px;
    background-color: #fd8c73;
    margin-top: 8px;
  }

	#linkGitHub {
			background-color: black;
			color: white;
	}
	#linkX {
			background-color: #0078D4;
			color: white;
	}
	#linkInstagram {
			background-color: #E1306C;
			color: white;
	}
	#linkNote {
			background-color: white;
	}
	#linkYouTube {
			background-color: #FF0000;
			color: white
	}
	#linkBlog {
			background-color: #f6f8fa;
	}
	#linkGames {
			background-color: #f6f8fa;
	}

	#linkGitHub a,#linkX a,#linkInstagram a,#linkNote a,#linkYouTube a,#linkBlog a,#linkGames a {
			color: inherit; /* 親要素の色を継承 */
			text-decoration: none; /* 下線を消す */
	}
</style>