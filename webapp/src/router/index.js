import Vue from 'vue'
import VueRouter from 'vue-router'
import Root from '../views/Root.vue'

Vue.use(VueRouter)

  const routes = [
  {
    path: '/',
    name: 'Root',
    component: Root
  },

  {
    path: '/home',
    name: 'Home',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/Home.vue')
  },

  {
    path: '/home/editor/:id',
    name: 'Editor',
    component: () => import('../views/Editor.vue')
  },

  {
    path: '/article/:id',
    name: 'Article',
    component: () => import('../views/Article.vue')
  },
  
  {
    path: '/confirm',
    name: 'Confirm',
    component: () => import( '../views/Confirm.vue')
  },

  {
    path: '/forgot',
    name: 'Forgot',
    component: () => import( '../views/Forgot.vue')
  },

  {
    path: '/:uuid',
    name: 'Section',
    component: () => import('../views/Section.vue')
  },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
