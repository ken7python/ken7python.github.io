<script setup lang="ts">
    import Header from '../components/Header.vue';
    import { ref, onMounted } from 'vue';

    const url :string = "https://api.github.com/users/ken7python/repos?sort=updated";

    const topRepos = ref<any[]>([]);
    onMounted(async () => {
        const res = await fetch(url);
        const data = await res.json();

        topRepos.value = data.slice(0,3).map(item => {
            const desc :string = item.description || null;
            let data = {
                name: item.name,
                url: item.html_url,
                description: desc,
                stars: item.stargazers_count,
                language: item.language || "未設定",
                updated: timeAgo(item.updated_at),
            }
            console.log(data);
            return data;
        })
    })
    function timeAgo(date: string | Date) {
        const now: number = new Date().getTime();
        const target: number = new Date(date).getTime();
        const diff = Math.floor((now - target) / (1000 * 60 * 60 * 24));
        return diff === 0 ? "今日" : `${diff}日前`;
    }
</script>

<template>
    <div class="page-container">
        <Header />
        <div class="main-content">
            <h1 class="page-title">🚀 直近のGitHubリポジトリ</h1>
            <div class="repos-container">
                <a v-for="repo in topRepos" :key="repo.url" class="repo-item" :href="repo.url" target="_blank">
                    <div class="repo-header">
                        <h2 class="repo-name">
                            <p>{{ repo.name }}</p>
                        </h2>
                        <div class="star-badge">⭐ {{ repo.stars }}</div>
                    </div>
                    <p class="repo-description">{{ repo.description || "説明がありません" }}</p>
                    <div class="repo-info">
                        <span class="language-tag">{{ repo.language }}</span>
                        <span class="update-time">{{ repo.updated }}に更新</span>
                    </div>
                </a>
            </div>
        </div>
    </div>
</template>

<style scoped>
/* ページ全体の設定 */
.page-container {
    background-color: #f0f4f8;
    min-height: 100vh;
    font-family: Arial, sans-serif;
}

.main-content {
    max-width: 800px;
    margin: 0 auto;
    padding: 20px;
}

/* タイトルのスタイル */
.page-title {
    text-align: center;
    color: #2d3748;
    font-size: 2.5em;
    margin-bottom: 30px;
    text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.1);
}

/* リポジトリ一覧のコンテナ */
.repos-container {
    display: flex;
    flex-direction: column;
    gap: 20px;
}

/* 各リポジトリのカード */
.repo-item {
    background-color: white;
    border-radius: 10px;
    padding: 20px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    border: 2px solid transparent;
    transition: all 0.3s ease;

    text-decoration: none;
}

.repo-item:hover {
    transform: translateY(-5px);
    box-shadow: 0 8px 15px rgba(0, 0, 0, 0.2);
    border-color: #4299e1;
}

/* リポジトリのヘッダー部分 */
.repo-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 15px;
}

.repo-name {
    margin: 0;
    font-size: 1.5em;
}

.repo-name a {
    color: #2b6cb0;
    text-decoration: none;
    font-weight: bold;
}

.repo-name a:hover {
    color: #2c5282;
    text-decoration: underline;
}

/* スターのバッジ */
.star-badge {
    color: #c53030;
    padding: 5px 10px;
    border-radius: 15px;
    border: 1px solid black;
    font-size: 0.9em;
    font-weight: bold;
}

/* 説明文 */
.repo-description {
    color: #4a5568;
    line-height: 1.6;
    margin-bottom: 15px;
    font-size: 1.1em;
}

/* リポジトリの情報部分 */
.repo-info {
    display: flex;
    justify-content: space-between;
    align-items: center;
}

/* 言語タグ */
.language-tag {
    background-color: #bee3f8;
    color: #2a69ac;
    padding: 5px 10px;
    border-radius: 5px;
    font-size: 0.9em;
    font-weight: bold;
}

/* 更新時間 */
.update-time {
    color: #718096;
    font-size: 0.9em;
}

/* スマホ対応 */
@media (max-width: 600px) {
    .main-content {
        padding: 15px;
    }
    
    .page-title {
        font-size: 2em;
    }
    
    .repo-header {
        flex-direction: column;
        align-items: flex-start;
        gap: 10px;
    }
    
    .repo-info {
        flex-direction: column;
        align-items: flex-start;
        gap: 10px;
    }
    
    .repo-name {
        font-size: 1.3em;
    }
}
</style>