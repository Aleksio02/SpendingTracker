<template>
    <side-bar></side-bar>
    <page-content></page-content>
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
  
  
  