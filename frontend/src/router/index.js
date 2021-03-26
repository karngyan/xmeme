import { createRouter, createWebHistory } from 'vue-router'
import PageHome from "@/pages/PageHome";
import PageConfigure from "@/pages/PageConfigure";

const routes = [
  {
    path: '/',
    name: 'Home',
    component: PageHome
  },
  {
    path: '/configure',
    name: 'Configure',
    component: PageConfigure
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
