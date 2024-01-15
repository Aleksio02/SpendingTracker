<template>
    <div class='signup_form'>
        <form id="login" name="form" action="<?php bloginfo('template_directory')?>/registration/signin.php" method="post">

            <p id="message">Вход на сайт</p>
            <div class="form_block">
                <div class="input_wrapper">
                    <input id="username" type="text" placeholder="Логин" name="username">
                </div>
            </div>

            <div class="form_block">
                <div class="input_wrapper">
                    <input id="password" type="password" placeholder="Пароль" name="password">
                </div>
            </div>
            
            <input id="submit" class="submit_btn" type="submit" name="submit" value="Войти">

        </form>
        <div class="no_account">Нет аккаунта? - <router-link to="/register">Зарегистрироваться</router-link></div>
        <div class="lost_password"><a href="<?= home_url(); ?>/lost-password/">Забыли пароль?</a></div>
    </div>
</template>
  
<script>
export default {
    data() {
        return {
            login: '',
            email: '',
            password: '',
            confirmPassword: '',
            agreeTerms: true
        };
    },
    methods: {
        register() {
            // Verify password match
            if (this.password !== this.confirmPassword) {
                console.log('Passwords do not match');
                return;
            }

            // Verify terms agreement
            if (!this.agreeTerms) {
                console.log('Please agree to the terms of service');
                return;
            }

            // Registration logic
            console.log('Registration Details:', {
                login: this.login,
                email: this.email,
                password: this.password
            });

            // Clear form fields after successful registration
            this.login = '';
            this.email = '';
            this.password = '';
            this.confirmPassword = '';
            this.agreeTerms = true;
        },
        switchPage(page) {
            this.$emit('switch-page', page);
        }
    }
};
</script>
  
<style scoped></style>
  
  
  