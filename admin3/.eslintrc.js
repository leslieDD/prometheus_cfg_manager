module.exports = {
  extends: [
    'plugin:vue/vue3-recommended',
    'prettier',
    'prettier/vue',
    'eslint:recommended',
    '@vue/typescript/recommended',
    'airbnb-base'
  ],
  rules: {
    'vue/no-unused-vars': 'error',
    'vue/no-multiple-template-root': 'off',
    'vue/no-v-for-template-key': 'on',
  },
};
