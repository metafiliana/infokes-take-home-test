<template>
  <div class="explorer-container">
    <div class="left-panel">
      <FolderTree 
        :folders="allFolders" 
        @select-folder="selectFolder"
      />
    </div>
    <div class="right-panel">
      <SubfolderList 
        :subfolders="selectedFolderSubfolders"
        :selected-folder="selectedFolder"
        @select-folder="selectSubfolder"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue';
import FolderTree from '@/components/FolderTree.vue';
import SubfolderList from '@/components/SubfolderList.vue';
import { useFolderExplorer } from '@/composables/useFolderExplorer';
import type { Folder } from '@/models/Folder';

const { 
  allFolders, 
  selectedFolderSubfolders, 
  selectedFolder,
  loadAllFolders, 
  selectFolder 
} = useFolderExplorer();

const selectSubfolder = (subfolder: Folder) => {
  selectFolder(subfolder);
};

onMounted(loadAllFolders);
</script>

<style scoped>
.explorer-container {
  display: flex;
  height: 100vh;
}
.left-panel, .right-panel {
  width: 50%;
}
</style>