<template>
  <div class="folder-tree">
    <ul>
      <RecursiveFolderItem 
        v-for="folder in folders" 
        :key="folder.id"
        :folder="folder"
        :depth="0"
        :isSelected="selectedFolder?.id === folder.id"
        @select-folder="onSelectFolder"
      />
    </ul>
  </div>
</template>

<script setup lang="ts">
import type { Folder } from '@/models/Folder';
import RecursiveFolderItem from './RecursiveFolderItem.vue';
import { ref } from 'vue';

const props = defineProps<{
  folders: Folder[]
}>()

const emit = defineEmits<{
  (e: 'select-folder', folder: Folder): void
}>()

const selectedFolder = ref<Folder | null>(null)

function onSelectFolder(folder: Folder) {
  selectedFolder.value = folder;
  emit('select-folder', folder);
}
</script>

<style scoped>
.folder-tree {
  height: 100%;
  overflow-y: auto;
  border-right: 1px solid #e0e0e0;
}
</style>