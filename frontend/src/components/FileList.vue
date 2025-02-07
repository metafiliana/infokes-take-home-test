<template>
    <div class="file-list">
      <h3>Files of {{ selectedFolder?.name }}</h3>
      <ul v-if="files.length">
        <li 
          v-for="file in files" 
          :key="file.id" 
          class="file-item"
          :class="{ 'selected': selectedFile?.id === file.id }"
          @click="onSelectFile(file)"
        >
          {{ file.name }}
        </li>
      </ul>
      <p v-else>No files</p>
    </div>
  </template>
  
  <script setup lang="ts">
  import type { File } from '@/models/File';
  import { ref } from 'vue';
  
  defineProps<{
    files: File[];
    selectedFolder: File | null;
  }>()
  
  const emit = defineEmits<{
    (e: 'select-file', file: File): void
  }>()
  
  const selectedFile = ref<File | null>(null)
  
  function onSelectFile(file: File) {
    selectedFile.value = file;
    emit('select-file', file);
  }
  </script>
  
  <style scoped>
  .file-list {
    padding: 10px;
    height: 100%;
    overflow-y: auto;
  }
  .file-item {
    padding: 5px;
    border-bottom: 1px solid #f0f0f0;
    cursor: pointer;
    transition: background-color 0.3s ease;
  }
  .file-item:hover {
    background-color: #f5f5f5;
  }
  .file-item.selected {
    background-color: #e0e0e0;
    font-weight: bold;
  }
  </style>