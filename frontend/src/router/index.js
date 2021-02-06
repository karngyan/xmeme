import { createRouter, createWebHistory } from 'vue-router'
import PageHome from "@/pages/PageHome";

const routes = [
  {
    path: '/',
    name: 'Home',
    component: PageHome
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
