import { ref, type Ref } from 'vue';
import type { Folder } from '@/models/Folder';
import { folderService } from '@/services/folderService';

export function useFolderExplorer() {
  const allFolders: Ref<Folder[]> = ref([]);
  const selectedFolderSubfolders: Ref<Folder[]> = ref([]);
  const selectedFolder: Ref<Folder | null> = ref(null);
  const selectedFile: Ref<File | null> = ref(null);

  async function loadAllFolders() {
    allFolders.value = await folderService.getAllFolders();
  }

  async function selectFolder(folder: Folder) {
    selectedFolder.value = folder;
    
    // Explicitly hit getSubfolders API for both root folders and subfolders
    try {
      selectedFolderSubfolders.value = await folderService.getSubfolders(folder.id);
    } catch (error) {
      console.error('Error fetching subfolders:', error);
      selectedFolderSubfolders.value = [];
    }
  }

  async function selectFile(file: File) {
    selectedFile.value = file;
  }

  return {
    allFolders,
    selectedFolderSubfolders,
    selectedFolder,
    loadAllFolders,
    selectFolder,
    selectFile
  };
}