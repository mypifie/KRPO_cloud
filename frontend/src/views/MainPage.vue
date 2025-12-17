<template>
  Главная страница
  <section>
    <FileIcon
      v-for="(file, index) of files"
      :key="index"
      :file-name="file"
    ></FileIcon>
  </section>
</template>

<script setup lang="ts">
import FileIcon from '@/shared/components/ui/file/FileIcon.vue'

const files = ['фаsssssssssssssssssssssssssйл_1', "файл_2"]
</script>

<style scoped></style>
