import RegistrationPage from './pages/RegistrationPage.vue';
import LoginPage from './pages/LoginPage.vue';
import HomePage from './pages/HomePage.vue';
import ProfileSettingsPage from './pages/ProfileSettingsPage.vue';
import { createRouter, createWebHistory } from 'vue-router';

const routes = [
  { path: '/register', component: RegistrationPage },
  { path: '/login', component: LoginPage },
  { path: '/profile-settings', component: ProfileSettingsPage },
  { path: '/', component: HomePage },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;