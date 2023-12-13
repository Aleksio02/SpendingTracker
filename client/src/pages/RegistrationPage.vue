<template>
  <div class='signup_form'>

    <form id="wp_signup_form" @submit.prevent="register">
    <p id="message">Регистрация</p>
      <form-block>
        <username-text-field/>
      </form-block>
      <form-block>
        <email-text-field/>
      </form-block>

      <div class="form_block">
        <label for="password">Пароль</label>
        <div class="input_wrapper">
          <input type="password" name="password" id="password" placeholder="Введите пароль" value="" required>
        </div>
        <div id="password_err" class="msg"></div>
      </div>

      <div class="form_block">
        <label for="password_confirmation">Повторите пароль</label>
        <div class="input_wrapper">
          <input type="password" name="password_confirmation" id="password_confirmation" placeholder="Введите пароль"
            value="" required>
        </div>
        <div id="password_confirmation_err" class="msg"></div>
      </div>

      <div class="form_block terms">
        <div class="input_wrapper">
          <input name="terms" id="terms" type="checkbox" value="0" required>
          <label for="terms">Я согласен(-на) с условиями предоставления услуг</label>
        </div>
      </div>
      <input type="submit" id="submitbtn" class="submit_btn" name="submit" value="Зарегистрироваться" />

      <div class="no_account">Уже есть аккаунт? - <router-link to="/login">Войти</router-link></div>
    </form>

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
      confirmPasswordErrorMessage: '',
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

<style scoped>

</style>


