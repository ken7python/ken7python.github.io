import { createRouter, createWebHistory} from 'vue-router'
import Home from '../views/Home.vue'
import Blog from '../views/Blog.vue'
import Game from '../views/Game.vue'

const routes = [
    { path: '/', component: Home, name: 'Home' },
    { path: '/blog', component: Blog, name: 'Blog' },
    { path: '/game', component: Game, name: 'Game' }
]

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes,
})

export default router