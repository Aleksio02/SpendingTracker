<script>
export default {
  name: 'birthdate-text-field'
}
</script>

<template>
    <label for="birthdate">Дата рождения</label>
    <div class="input_wrapper">
      <input type="date" name="birthdate" id="birthdate" placeholder="Введите дату рождения" value="" required>
    </div>
    <div id="birthdate_err" class="msg"></div>
</template>

<style scoped>

</style>