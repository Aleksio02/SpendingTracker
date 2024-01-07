<template>
    <div class='signup_form'>
        <form id="login" name="form" action="" method="post">

            <page-title>Вход на сайт</page-title>
            <form-block>
                <username-text-field />
            </form-block>
            <form-block>
                <password-text-field>Пароль</password-text-field>
            </form-block>
            <input id="submit" class="submit_btn" type="submit" name="submit" value="Войти">

        </form>
        <div class="no_account">Нет аккаунта? - <router-link to="/register">Зарегистрироваться</router-link></div>
        <div class="lost_password"><a href="">Забыли пароль?</a></div>
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
  
  
  