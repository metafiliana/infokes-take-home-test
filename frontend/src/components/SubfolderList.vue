<template>
  <div class="subfolder-list">
    <h3 v-if="selectedFolder">Subfolders of {{ selectedFolder.name }}</h3>
    <h3 v-else>Select a folder to view subfolders</h3>
    <ul v-if="selectedFolder && subfolders.length">
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
    <p v-else-if="selectedFolder">No subfolders</p>
    <!-- Files list for each subfolder -->
    <ul v-if="selectedFolder && selectedFolder.files && selectedFolder.files.length" class="files-list">
      <li 
        v-for="file in selectedFolder.files" 
        :key="file.id" 
        class="file-item"
      >
        {{ file.name }}
      </li>
    </ul>
    <p v-else>No files</p>
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
  height: 100%;
  margin: 0;
  box-sizing: border-box;
}
.subfolder-item {
  cursor: pointer;
}
.subfolder-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.file-count {
  color: #888;
  font-size: 0.8em;
}
.files-list {
  margin-top: 5px;
  padding-left: 15px;
  font-size: 0.9em;
  color: #666;
}
.file-item {
  padding: 2px 0;
  border-bottom: 1px solid #f5f5f5;
}
.subfolder-item.selected {
  background-color: #e0e0e0;
  font-weight: bold;
}
</style>