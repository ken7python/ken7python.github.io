<template>
  <header>
    <div id="menuBar">
      <!-- <div id="AlwaysMenu"> -->
        <!-- <button id="MenuToggle" class="menu-button" @click="toggleMenu">...</button> -->
        <!-- <div id="OriginalMenu"> -->
          
        <!-- </div> -->
      <!-- </div> -->

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

      <!-- <div id="SNSMenu" :class="{ 'open': menuOpen }"> -->
      <div class="link-sns">
        <a class="link" id="linkGitHub" target="_blank" rel="noopener noreferrer" href="https://github.com/ken7python">
          <img class="logo" :src="githubLogo" alt="GitHub Logo" id="icon-github">
          <!-- GitHub -->
        </a>

        <a class="link" id="linkX" target="_blank" rel="noopener noreferrer" href="https://x.com/ken7python">
          <img class="logo" :src="xLogo" alt="X Logo" id="icon-x">
          <!-- X -->
        </a>
        <a class="link" id="linkInstagram" target="_blank" rel="noopener noreferrer" href="https://www.instagram.com/ken7python/">
          <img class="logo" :src="InstagramLogo" alt="Instagram Logo" id="icon-instagram">
          <!-- Instagram -->
        </a>
        <a class="link" id="linkNote" target="_blank" rel="noopener noreferrer" href="https://note.com/ken7python">
          <img class="logo" :src="noteLogo" alt="Note Logo" id="icon-note">
          <!-- Note -->
        </a>
        <a class="link" id="linkYouTube" target="_blank" rel="noopener noreferrer" href="https://www.youtube.com/@KenCodeHanabi">
          <img class="logo" :src="youtubeLogo" alt="YouTube Logo" id="icon-youtube">
          <!-- YouTube -->
        </a>
      </div>
    </div>
	</header>
</template>

<script setup lang="ts">
	import { ref } from 'vue';
  import { useRoute } from 'vue-router';

  const githubLogo = './github-logo.png'
  const xLogo = './x-logo.png'
  const InstagramLogo = './Instagram-logo.png'
  const noteLogo = './note-logo.png'
  const youtubeLogo = './youtube-logo.png'

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
  #menuBar {
    padding-top: 50px;

    display: flex;
    flex-direction: row;
    /* justify-content: flex-end; */
    

    width: 98vw;
  }
  .logo {
    /* width: 24px; */
    height: 24px;
    vertical-align: middle; /* アイコンとテキストを揃える */
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
			/* padding: 8px 16px; */
      /* padding-right: 8vw; */
	}

  #OriginalMenu {
    display: flex;
    flex-direction: row;
    /* margin-top: auto; */
    /* gap: 16px; */
    /* position: relative;*/
  }

	@media (max-width: 1024px) { /* PC向けデザイン */
		#SNSMenu.open {
      display: flex;
    }
	}

	/* リンクたち */
	@media (max-width: 770px) { /* スマホ向けデザイン */ /* スマホでは「…」ボタンを表示 */
			.menu-button {
					display: block;      /* 「…」ボタンを表示 */
			}
			#SNSMenu {
					display: none;    /* 初期状態ではリンクを隠す */
					flex-direction: column;
					/* padding: 0 16px; */
			}
			#SNSMenu.open {
					display: flex;       /* JS で「open」を付けると表示 */
			}
      #AlwaysMenu .link {
        font-size: 10px;
        /* padding: 4px 8px; */
        margin: 4px 0;
      }
	    #SNSMenu .link {
		    font-size: 12px;
		    /* padding: 8px 16px; */
		    margin: 8px 0;
	    }
      ion-icon {
        position: relative;
        font-size: 12px;
        top: 2px; /* アイコンを少し上に */
        margin-right: 8px; /* アイコンとテキストの間にスペースを追加 */
      }

      #menuBar {
        flex-direction: column;
        align-items: center;
        padding: 10px;
      }

      .menu-item {
        margin: 0;
        text-align: center;
      }

      .link-sns {
        display: flex;
        flex-wrap: wrap;
        justify-content: center;
        gap: 10px;
        margin-top: 20px;
      }

      .link {
        font-size: 12px;
      }

      .logo {
        /* width: 20px; */
        height: 20px;
      }
	}

	@media (min-width: 771px) { /* タブレット・PC向けデザイン */
			header {
					/* display: flex;
					flex-direction: column;
					justify-content: flex-end; */
					/* align-items: center; */
					/* padding: 8px 16px; */
			}

      .link {
        font-size: 14px;
        padding: 8px 16px;
        margin: 8px 0;
      }
      ion-icon {
        position: relative;
        font-size: 16px;
        top: 2px; /* アイコンを少し上に */
        margin-right: 8px; /* アイコンとテキストの間にスペースを追加 */
      }

      .menu-item {
        /* margin: 0 15px; */
      }

      .link-sns {
        display: flex;
        justify-content: flex-end;
        /* gap: 15px; */
      }

      .link {
        font-size: 14px;
      }

      .logo {
        /* width: 24px; */
        height: 24px;
      }
	}

	.link{
			text-decoration: none;
			text-align: center;
			border-radius: 4px;
			color: #1f2329;
			/* font-weight: bold; */
			border: none;
	}

  .underLine {
    width: 100%;
    height: 2px;
    background-color: #fd8c73;
    margin-top: 8px;
  }

  /* HomeとBlogなどのブロック間の間隔を詰め、ヘッダーの下に寄せる */
#menuBar {
  flex-direction: row;
  justify-content: start;
  padding: 0;
}

.menu-item {
  margin: 0;
}

.header {
  padding-bottom: 0;
}

/* .menu-itemの下の余白を削除して下に寄せる */
.menu-item {
  margin-bottom: 0;
}

/* .menu-item内のaタグを下に寄せる */
.menu-item .link {
  display: flex;
  align-items: flex-end;
}
</style>