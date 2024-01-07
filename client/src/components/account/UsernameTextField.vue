<script>
export default {
  name: 'username-text-field'
}
</script>

<template>
    <label for="username">Логин</label>
    <div class="input_wrapper">
      <input id="username" type="text" placeholder="Введите логин" name="username">
    </div>
</template>

<style scoped>
label {
  line-height: 25px;
}
</style>