<script>
export default {
    name: 'page-content'
}
</script>

<template>
    <div class="page-content">
        <content-block-title>Подробная статистика</content-block-title>
        <div class="user_about_content">
            <content-block-text>
                <p>За месяц вы потратили:</p>
            </content-block-text>
        </div>
    </div>
</template>

<style scoped>
.page-content {
    font-family: 'Montserrat', sans-serif;
    position: relative;
    margin-left: 360px;
    margin-top: -50px;

    width: 800px;
    height: fit-content;
    min-height: 1500px;
    padding: 10px 20px 50px;

    border-radius: 30px;
    color: #FFFFFF;
    background: rgba(0, 0, 0, 0.5);
}
</style>