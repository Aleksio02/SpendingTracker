import RegistrationForm from './components/RegistrationForm.vue';
import LoginForm from './components/LoginForm.vue';
import { createRouter, createWebHistory } from 'vue-router';

const routes = [
    { path: '/', redirect: '/login' },
    { path: '/register', component: RegistrationForm },
    { path: '/login', component: LoginForm }
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
