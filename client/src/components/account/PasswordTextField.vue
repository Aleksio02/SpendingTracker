<script>
export default {
  name: 'password-text-field'
}
</script>

<template>
    <label for="username"><slot></slot></label>
    <div class="input_wrapper">
      <input id="password" type="password" placeholder="Введите пароль" name="password">
    </div>
</template>

<style scoped>

</style>