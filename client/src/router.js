import RegistrationPage from './pages/RegistrationPage.vue';
import LoginPage from './pages/LoginPage.vue';
import { createRouter, createWebHistory } from 'vue-router';

const routes = [
    { path: '/', redirect: '/login' },
    { path: '/register', component: RegistrationPage },
    { path: '/login', component: LoginPage }
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
