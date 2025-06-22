<script setup lang="ts">
	import Header from '../components/Header.vue';
  const icon = '/Icon.png'; // publicフォルダーの画像を参照
</script>

<template>
  <div>
		<Header />
    <meta charset="utf-8">
    <title>KenCode Portfolio</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <div id="main">
      <div id="container">
        <div id="idLeft">
          <div id="idDivIcon">
            <div id="idIcon">
              <img :src="icon" alt="Icon">
            </div>
            <div id="idName">
              <h2 style="margin-top: 0;margin-bottom: 0;">KenCode</h2>
              <h2 style="margin-top: 0;margin-bottom: 0;font-weight: 300;">ken7python</h2>
            </div>
          </div>
          <hr>
          <div id="idSkill">
            <h3>資格</h3>
            <ul>
              <li>全商ビジネス文書実務検定２級</li>
              <li>全商情報処理検定２級</li>
              <li>ITパスポート</li>
              <li>基本情報技術者</li>
            </ul>
          </div>
          <hr>
          <div id="idAchievements">
            <h3>受賞歴</h3>
            <ul>
              <li>第６回越前がにロボコン　優勝</li>
              <li>PCNこどもプロコン2021<br>北陸地方　一次審査　最優秀賞・優秀賞</li>
              <li>PCNこどもプロコン2023　協賛賞</li>
              <li>ふくいソフトウェアコンペティション2023<br>IT産業連合会会長賞</li>
              <li>ふくいソフトウェアコンペティション2024<br>大賞・知事賞</li>
            </ul>
          </div>
          <hr>
        </div>
        <div id="idRight">
          <div id="idREADME">
            <h2>KenCode/README.<span style="color: gray">md</span></h2>
            <br>
            <p>KenCodeのポートフォリオにようこそ</p>
            <p>プログラミングと未来設計が大好き！</p>
            <p>未来をコードで変える高校生エンジニア</p>
          </div>
          <div id="idGrass">

          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
	@media (min-width: 1025px) { /* PC向けデザイン */
    #container {
        display: flex;
        margin-top: 44px;
        margin-left: 7vw;
        margin-right: 4vw;
    }
    #idLeft {
        float:left;
        width: 23vw;
    }
    #idRight {
        float:right;
        width: 67vw;
    }
    #idIcon {
        margin-right: 32px;
        margin-bottom: 16px;
    }
    #idIcon img {
        display: block;
        /* margin: 0 auto; */
        width: 100%;
        height: auto;
    }
    li {
        font-size: 14px;
    }
	}

	@media (max-width: 1024px) { /* タブレット・スマホ向けデザイン */
			#container {
					display: block;
					margin-top: 44px;
					margin-left: 2vw;
					margin-right: 2vw;
			}
			#idLeft {
					width: 100%;
					float:none;
			}
			#idRight {
					width: 100%;
					float:none;
			}
			#idDivIcon {
					display: flex;
					margin-top: 50px;
					margin-bottom: 50px;
					align-items: center;
			}
			#idName {
					margin-left: 24px;
			}
			#idIcon img {
					width: 16vw;
					height: auto;
			}
	}

	/* アイコンのスタイル */

	#idIcon img{
			border-radius: 50%;
			border: 3px solid #e1e2e3;
			object-fit: cover;
	}

	/* README */

	#idREADME {
			border: 1px solid #D9D9D9;
			height:45vh;
			padding: 16px;
	}

	/* 草 */

	#idGrass {
			border: 1px solid #D9D9D9;
			height:34vh;
	}

	/* hrの色 */

	hr {
			color: #dfe5ea;
	}

	/* フォントの色 */

	p,h2,h3,h4,li,a{
			color: #1f2329;
	}
	</style>