<template>
    <link href="https://fonts.googleapis.com/css2?family=Kaushan+Script&family=Montserrat:wght@500&display=swap
    " rel="stylesheet">

    <div class="header" v-if="!isLoginPage && !isRegisterPage">
        <div class="header_inner">

            <div class="user_info">
                <div class="profile_photo" style="background: #555555 center no-repeat; background-size: cover">
                </div>

                <div class="user_top">
                    <div class="user_icons">
                        <img class="icon logout" src="./assets/icons/logout.png" alt="" v-on:click="logoutUser">
                        <img class="icon" src="./assets/icons/notifications.png" alt="">
                        <img class="icon" src="./assets/icons/rating.png" alt="">
                    </div>

                    <router-link to="/" class="user_name">{{ user.username ? user.username : "-" }}</router-link>
                </div>
            </div>

            <div class="nav" id="nav">
                <router-link to="/news" class="nav_item">Новости</router-link>
                <router-link to="/" class="nav_item">Статистика</router-link>
                <router-link to="/guide" class="nav_item">Инструкция</router-link>
                <router-link to="/about" class="nav_item">О нас</router-link>
                <a class="logo" href="index.html" target="_self"></a>
            </div>
        </div>

        <div class="top_buttons">
            <div class="top_buttons-left">
                <router-link to="/" class="button" id="account">Личный кабинет</router-link>
            </div>
            <div class="top_buttons-right">
                <a class="button" id="co-working" href="#">Сотрудничество</a>
            </div>
        </div>
    </div>

    <div class="intro">
        <div class="intro_fade"></div>
    </div>

    <div class="section">
        <div class="container">
            <router-view></router-view>
        </div>
    </div>
</template>

<script>
import { deleteSession } from './modules/auth';

export default {
    name: 'App',
    data() {
        return {
            user: {
                username: '',
            },
        }
    },
    mounted() {
        const userData = localStorage.getItem('user');
        if (userData) {
            this.user = JSON.parse(userData);
        }
    },
    computed: {
        isLoginPage() {
            return this.$route.path === '/login'; // Проверка текущего маршрута на страницу входа (login)
        },
        isRegisterPage() {
            return this.$route.path === '/register'; // Проверка текущего маршрута на страницу регистрации (register)
        }
    },
    methods: {
        logoutUser() {
            deleteSession();
        }
    }
}
</script>

<style scoped>
* {
    box-sizing: content-box;
}

.chart {
    margin-bottom: 30px;
}

.container {
    display: flex;
    justify-content: space-between;
    max-width: 1200px;
    padding: 0 20px;
    margin: 0 auto;
    font-family: 'Montserrat', sans-serif;
}

.header {
    width: auto;
    height: 70px;
    padding: 0 100px;

    font-family: 'Montserrat';
    font-size: 15px;
    text-transform: uppercase;
    line-height: 70px;

    border-bottom: 1px solid transparent;
    border-image: radial-gradient(#ffffff 60%, transparent);
    border-image-slice: 1;

    background-color: #000000;
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    z-index: 100;
}

.header_inner {
    width: 100%;
    height: inherit;
    display: flex;
    justify-content: space-between;
}

.user_info {
    display: flex;
}

.user_top {
    display: flex;
    flex-direction: column;
    justify-content: center;
    margin-left: 20px;
}

.user_icons {
    display: flex;
    justify-content: left;
    margin: 5px 0;
}

.user_icons .icon {
    margin-right: 20px;
}

.icon {
    -webkit-filter: sepia(1) saturate(0);
    filter: sepia(1) saturate(0);
    width: 15px;
    height: 15px;
}

.icon:hover {
    -webkit-filter: sepia(1) hue-rotate(70deg) saturate(100);
    filter: sepia(1) hue-rotate(70deg) saturate(100);
}

.logout_btn {
    line-height: 0;
}

.profile_photo {
    width: 100px;
    height: 100px;

    margin-top: 15px;

    background-color: #555555;

    border: 2px #FFFFFF solid;
    border-radius: 50%;
}

.user_name {
    margin-top: 7px;
    line-height: 20px;
    color: #ffffff;
    text-transform: none;
}

.user_name:hover {
    color: #6dff72;
}

.top_buttons {
    display: flex;
    justify-content: space-between;

    width: inherit;
    line-height: 25px;

    margin-top: 15px;
    padding: 0 0 0 110px;
}

.top_buttons * {
    line-height: inherit;
}

.button {
    margin: 0 15px;
    padding: 5px 30px;

    border: 1px #ffffff solid;
    border-radius: 50px;

    font-size: 13px;
    text-transform: uppercase;
    text-decoration: none;
    color: #dff4ff;

    transition: color, background .3s linear;
}

.button.profile {
    margin-right: 130px;
}

.button:hover {
    background-color: #ffffff;
    color: #000000;

    transition: color, background .3s linear;
}

.logo {
    height: 100%;
    width: 130px;

    background: url(assets/logo.png) center no-repeat;
    background-size: contain;

    filter: drop-shadow(0 0 0 #000000) brightness(150%);
    transition: filter .3s linear;
}

.logo:hover {
    filter: drop-shadow(0 0 30px #003971) brightness(200%);
    transition: filter .3s linear;
}

.nav {
    display: flex;
}

.nav_item {
    padding: 0 15px;

    /* font-weight: 700; */
    color: #FFFFFF;
    text-decoration: none;

    transition: background .3s linear;
}

.nav_item:hover {
    background-color: #00366c;
    transition: background .3s linear;
}

.nav_item.active {
    background-color: #00366c;
}

.header_fade {
    width: 100%;
    height: 300px;
    background-image: linear-gradient(#00111e 70px, transparent);

    position: fixed;
    top: 0;
    left: 0;
    z-index: 90;
}

/* Section */

.section {
    margin-top: 200px;
    padding: 50px 0 100px;
    line-height: 30px;
    text-align: justify;
}

.section_header {
    font-size: 25px;
    text-transform: uppercase;
    text-align: center;
}

.section_header:after {
    content: "";
    display: block;
    margin: 20px auto 30px;

    width: 400px;
    height: 1px;

    background-image: radial-gradient(#ffffff -30%, transparent);
}

.container p {
    max-width: 750px;
    margin: 15px auto;
}

.menu_inner {
    height: 50px;
    display: flex;
    justify-content: center;
    font-size: 15px;
}

.menu_item {
    padding: 10px 30px;

    color: #fff;
    text-transform: uppercase;
    text-decoration: none;

    transition: color .3s linear;

    border-bottom: 1px #ffffff solid;
}

.menu_item:hover {
    border-top: 1px #30ff58 solid;
    border-bottom: 1px #30ff58 solid;
    color: #30ff58;
    transition: color .3s linear;
}
</style>