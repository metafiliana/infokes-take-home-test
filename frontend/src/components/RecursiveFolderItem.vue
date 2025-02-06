<template>
  <li 
    :class="['folder-item', { 'selected': isSelected }]"
    :style="{ paddingLeft: `${depth * 20}px` }"
    @click.stop="selectFolder"
  >
    {{ folder.name }}
  </li>
</template>

<script setup lang="ts">
import type { Folder } from '@/models/Folder';
import { computed } from 'vue';

const props = defineProps<{
  folder: Folder;
  depth: number;
  isSelected: boolean;
}>()

const emit = defineEmits<{
  (e: 'select-folder', folder: Folder): void
}>()

function selectFolder() {
  emit('select-folder', props.folder);
}
</script>

<style scoped>
.folder-item {
  cursor: pointer;
  padding: 5px 10px;
  user-select: none;
  transition: background-color 0.2s;
}
.folder-item:hover {
  background-color: #f0f0f0;
}
.selected {
  background-color: #e0e0e0;
}
</style>