<template>
  <div class="subfolder-list">
    <h3>Subfolders of {{ selectedFolder?.name }}</h3>
    <ul v-if="subfolders.length">
      <li 
        v-for="subfolder in subfolders" 
        :key="subfolder.id" 
        class="subfolder-item"
        :class="{ 'selected': selectedSubfolder?.id === subfolder.id }"
        @click="onSelectSubfolder(subfolder)"
      >
        {{ subfolder.name }}
      </li>
    </ul>
    <p v-else>No subfolders</p>
  </div>
</template>

<script setup lang="ts">
import type { Folder } from '@/models/Folder';
import { ref } from 'vue';

defineProps<{
  subfolders: Folder[];
  selectedFolder: Folder | null;
}>()

const emit = defineEmits<{
  (e: 'select-folder', folder: Folder): void
}>()

const selectedSubfolder = ref<Folder | null>(null)

function onSelectSubfolder(folder: Folder) {
  selectedSubfolder.value = folder;
  emit('select-folder', folder);
}
</script>

<style scoped>
.subfolder-list {
  padding: 10px;
  height: 100%;
  overflow-y: auto;
}
.subfolder-item {
  padding: 5px;
  border-bottom: 1px solid #f0f0f0;
  cursor: pointer;
  transition: background-color 0.3s ease;
}
.subfolder-item:hover {
  background-color: #f5f5f5;
}
.subfolder-item.selected {
  background-color: #e0e0e0;
  font-weight: bold;
}
</style>