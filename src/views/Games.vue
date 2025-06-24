<script setup lang="ts">
	import Header from '../components/Header.vue';
  import { ref } from 'vue';

  class Game {
    name: string;
    url: string;

    constructor(name: string, url: string) {
      this.name = name;
      this.url = url;
    }
  }

  let games = ref<Game[]>([]);
  games.value = [
    new Game('キッキングスナイパー', 'https://unityroom.com/games/kencode_kickgame'),
    new Game('Tank-Rush', 'https://unityroom.com/games/kencode_tankrush'),
    new Game('スカイバルーンアドベンチャー', 'https://unityroom.com/games/kencode_freefall/'),
    new Game('デジタル絵合わせ', 'https://ken7python.github.io/Eawase'),
    new Game('弾幕シューティング', 'http://ken7python.github.io/biba/'),
    new Game('オセロ', 'http://ken7python.github.io/Osero/'),
    new Game('勝山恐竜めぐり', 'http://ken7python.github.io/DinosaurTour/'),
    new Game('回り将棋', 'http://ken7python.github.io/MawariShogi/'),
    new Game('県庁所在地クイズ', 'http://ken7python.github.io/Todouhuken/'),
    new Game('1-9パズル', 'http://ken7python.github.io/pazuru/'),
    new Game('迷路3', 'https://unityroom.com/games/kencode_maze3'),
    new Game('Bring', 'https://unityroom.com/games/kencode_bring')
  ];

</script>

<template>
  <div>
    <Header />
    <div id="container">
      <h1>ゲーム</h1>
      <div class="game-grid">
        <a :href="game.url" target="_blank" rel="noopener noreferrer" class="game-card" v-for="game in games" :key="game.name">
          <p class="gameName">
            {{ game.name }}
          </p>
        </a>
      </div>
    </div>
  </div>
</template>

<style scope>

  .game-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
    gap: 20px;
  }

  .game-card {
    background-color: #f6f8fa;
    border: 1px solid #dfe5ea;
    padding: 10px;
    text-align: center;
    border-radius: 8px;
    transition: transform 0.2s;

    text-decoration: none;
  }

  .gameName {
    font-size: 1.2em;
    color: #1f2329;
    font-weight: bold;
  }

  .game-card:hover {
    transform: scale(1.05);
  }
</style>
